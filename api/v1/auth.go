/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Nicknamezz00/mercury/api/auth"
	"github.com/Nicknamezz00/mercury/internal/errmsg"
	"github.com/Nicknamezz00/mercury/internal/types"
	storepb "github.com/Nicknamezz00/mercury/proto/gen/store"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func (s *APIV1Service) registerAuth(g *echo.Group) {
	g.POST("/auth/signin", s.SignIn)
	g.POST("/auth/signout", s.SignOut)
	g.POST("/auth/signup", s.SignUp)
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var usernameRegex = regexp.MustCompile("^[a-z]([a-z0-9-]{1,30}[a-z0-9])?$")

func (s *APIV1Service) SignIn(c echo.Context) error {
	ctx := c.Request().Context()
	disablePasswordLogin, err := s.Store.GetSystemSetting(ctx, &store.FindSystemSetting{
		Name: SystemSettingDisablePasswordLoginName.String(),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetSystemSettingsErr).SetInternal(err)
	}
	if disablePasswordLogin != nil {
		var disable bool
		if err = json.Unmarshal([]byte(disablePasswordLogin.Value), &disable); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUnmarshalSystemSettingErr).SetInternal(err)
		}
		if disable {
			return echo.NewHTTPError(http.StatusUnauthorized, errmsg.LoginDisabledErr)
		}
	}

	signin := &SignIn{}
	if err := json.NewDecoder(c.Request().Body).Decode(signin); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedSignInRequestErr).SetInternal(err)
	}
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		Username: &signin.Username,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.IncorrectCredentialsErr)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.IncorrectCredentialsErr)
	} else if user.RowStatus == types.Archived {
		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("user: %s has been archived", signin.Username))
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(signin.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.IncorrectCredentialsErr)
	}

	accessToken, err := auth.GenerateAccessToken(user.Username, user.ID, time.Now().Add(auth.AccessTokenDuration), []byte(s.Secret))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGenerateAccessTokenErr).SetInternal(err)
	}
	if err := s.UpsertAccessToken(ctx, user, accessToken); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUpsertAccessTokenErr).SetInternal(err)
	}
	// TODO: activity
	cookieExpire := time.Now().Add(auth.CookieExpiration)
	setTokenCookie(c, auth.AccessTokenCookieName, accessToken, cookieExpire)
	msg := convertUserFromStore(user)
	return c.JSON(http.StatusOK, msg)
}

func (s *APIV1Service) UpsertAccessToken(ctx context.Context, user *store.User, token string) error {
	userAccessTokens, err := s.Store.GetUserAccessTokens(ctx, user.ID)
	if err != nil {
		return errors.Wrap(err, errmsg.FailedToGetAccessTokensErr)
	}
	accessToken := storepb.AccessTokensUserSetting_AccessToken{
		AccessToken: token,
		Description: "Sign in",
	}
	userAccessTokens = append(userAccessTokens, &accessToken)
	_, err = s.Store.UpsertUserSettingV1(ctx, &storepb.UserSetting{
		UserId: user.ID,
		Key:    storepb.UserSettingKey_USER_SETTING_ACCESS_TOKENS,
		Value: &storepb.UserSetting_AccessTokens{
			AccessTokens: &storepb.AccessTokensUserSetting{
				AccessTokens: userAccessTokens,
			},
		},
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUpsertUserSettingErr).SetInternal(err)
	}
	return nil
}

func (s *APIV1Service) SignOut(c echo.Context) error {
	ctx := c.Request().Context()
	accessToken := findAccessToken(c)
	userID, _ := getUserIDFromAccessToken(accessToken, s.Secret)
	userAccessTokens, err := s.Store.GetUserAccessTokens(ctx, userID)
	if err == nil && len(userAccessTokens) != 0 {
		var accessTokens []*storepb.AccessTokensUserSetting_AccessToken
		for _, t := range userAccessTokens {
			if accessToken != t.AccessToken {
				accessTokens = append(accessTokens, t)
			}
		}
		if _, err := s.Store.UpsertUserSettingV1(ctx, &storepb.UserSetting{
			UserId: userID,
			Key:    storepb.UserSettingKey_USER_SETTING_ACCESS_TOKENS,
			Value: &storepb.UserSetting_AccessTokens{
				AccessTokens: &storepb.AccessTokensUserSetting{
					AccessTokens: accessTokens,
				},
			},
		}); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUpsertUserSettingErr).SetInternal(err)
		}
	}
	removeAccessTokenAndCookie(c)
	return c.JSON(http.StatusOK, true)
}

func (s *APIV1Service) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	signup := &SignUp{}
	if err := json.NewDecoder(c.Request().Body).Decode(signup); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedSignUpRequestErr).SetInternal(err)
	}
	hostUserType := types.RoleHost
	existedHostUsers, err := s.Store.ListUsers(ctx, &store.FindUser{
		Role: &hostUserType,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if !usernameRegex.MatchString(strings.ToLower(signup.Username)) {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidUsernameErr).SetInternal(err)
	}
	userCreate := &store.User{
		Username: signup.Username,
		Role:     types.RoleUser,
		Nickname: signup.Username,
	}
	if len(existedHostUsers) == 0 {
		// Change the default role to host if there is no host user.
		userCreate.Role = types.RoleHost
	} else {
		allowSignUp, err := s.Store.GetSystemSetting(ctx, &store.FindSystemSetting{
			Name: SystemSettingAllowSignUpName.String(),
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetSystemSettingsErr).SetInternal(err)
		}
		var allow bool
		if allowSignUp != nil {
			err = json.Unmarshal([]byte(allowSignUp.Value), &allow)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUnmarshalSystemSettingErr).SetInternal(err)
			}
		}
		if !allow {
			return echo.NewHTTPError(http.StatusUnauthorized, errmsg.SignUpDisabledErr).SetInternal(err)
		}
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGeneratePasswordHashErr).SetInternal(err)
	}
	userCreate.PasswordHash = string(hash)
	user, err := s.Store.CreateUser(ctx, userCreate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToCreateUserErr).SetInternal(err)
	}
	accessToken, err := auth.GenerateAccessToken(user.Username, user.ID, time.Now().Add(auth.AccessTokenDuration), []byte(s.Secret))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGenerateAccessTokenErr).SetInternal(err)
	}
	if err := s.UpsertAccessToken(ctx, user, accessToken); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUpsertAccessTokenErr).SetInternal(err)
	}
	cookieExpiration := time.Now().Add(auth.CookieExpiration)
	setTokenCookie(c, auth.AccessTokenCookieName, accessToken, cookieExpiration)
	msg := convertUserFromStore(user)
	return c.JSON(http.StatusOK, msg)
}

func setTokenCookie(c echo.Context, name string, token string, expiration time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteStrictMode
	c.SetCookie(cookie)
}

func removeAccessTokenAndCookie(c echo.Context) {
	exp := time.Now().Add(-1 * time.Hour)
	setTokenCookie(c, auth.AccessTokenCookieName, "", exp)
}
