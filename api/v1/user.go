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
	"encoding/json"
	"github.com/Nicknamezz00/mercury/common/util"
	"github.com/Nicknamezz00/mercury/internal/errmsg"
	"github.com/Nicknamezz00/mercury/internal/types"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type User struct {
	ID int32 `json:"id"`

	// Standard fields
	RowStatus types.RowStatus `json:"rowStatus"`
	CreatedTs int64           `json:"createdTs"`
	UpdatedTs int64           `json:"updatedTs"`

	// Domain specific fields
	Username        string         `json:"username"`
	Role            types.Role     `json:"role"`
	Email           string         `json:"email"`
	Nickname        string         `json:"nickname"`
	PasswordHash    string         `json:"-"`
	AvatarURL       string         `json:"avatarUrl"`
	UserSettingList []*UserSetting `json:"userSettingList"`
}

type CreateUserRequest struct {
	Username string     `json:"username"`
	Role     types.Role `json:"role"`
	Email    string     `json:"email"`
	Nickname string     `json:"nickname"`
	Password string     `json:"password"`
}

type UpdateUserRequest struct {
	RowStatus *types.RowStatus `json:"rowStatus"`
	Username  *string          `json:"username"`
	Email     *string          `json:"email"`
	Nickname  *string          `json:"nickname"`
	Password  *string          `json:"password"`
	AvatarURL *string          `json:"avatarURL"`
}

func (s *APIV1Service) registerUserRoutes(g *echo.Group) {
	g.GET("/user", s.GetUserList)
	g.POST("/user", s.CreateUser)
	g.GET("/user/me", s.GetCurrentUser)
	// TODO: move to /api/v2
	g.GET("/user/name/:username", s.GetUserByUsername)
	g.GET("/user/:id", s.GetUserByID)
	g.PATCH("/user/:id", s.UpdateUser)
	g.DELETE("/user/:id", s.DeleteUser)
}

func (s *APIV1Service) GetUserList(c echo.Context) error {
	ctx := c.Request().Context()
	list, err := s.Store.ListUsers(ctx, &store.FindUser{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	users := make([]*User, 0, len(list))
	for _, user := range list {
		u := convertUserFromStore(user)
		// desensitize
		u.Email = ""
		users = append(users, u)
	}
	return c.JSON(http.StatusOK, users)
}

func (s *APIV1Service) CreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr)
	}
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &userID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr).SetInternal(err)
	}
	if user.Role != types.RoleHost {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.FailedToCreateUserErr).SetInternal(err)
	}
	userCreate := &CreateUserRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(userCreate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedCreateUserRequestErr).SetInternal(err)
	}
	if err := userCreate.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidUserCreateErr).SetInternal(err)
	}
	if !usernameRegex.MatchString(strings.ToLower(userCreate.Username)) {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidUsernameErr).SetInternal(err)
	}
	if userCreate.Role == types.RoleHost {
		return echo.NewHTTPError(http.StatusForbidden, "cannot create host user")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGeneratePasswordHashErr).SetInternal(err)
	}
	newUser, err := s.Store.CreateUser(ctx, &store.User{
		Username:     userCreate.Username,
		Role:         types.Role(userCreate.Role),
		Email:        userCreate.Email,
		Nickname:     userCreate.Nickname,
		PasswordHash: string(hash),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToCreateUserErr).SetInternal(err)
	}
	u := convertUserFromStore(newUser)
	// TODO: activity
	return c.JSON(http.StatusOK, u)
}

func (s *APIV1Service) GetCurrentUser(c echo.Context) error {
	ctx := c.Request().Context()
	userID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr)
	}
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &userID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr).SetInternal(err)
	}
	list, err := s.Store.ListUserSettings(ctx, &store.FindUserSetting{
		UserID: &userID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserSettings).SetInternal(err)
	}
	var userSettings []*UserSetting
	for _, us := range list {
		userSettings = append(userSettings, convertUserSettingFromStore(us))
	}
	u := convertUserFromStore(user)
	u.UserSettingList = userSettings
	return c.JSON(http.StatusOK, u)
}

func (s *APIV1Service) GetUserByUsername(c echo.Context) error {
	ctx := c.Request().Context()
	username := c.Param("username")
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		Username: &username,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr).SetInternal(err)
	}
	u := convertUserFromStore(user)
	u.Email = ""
	return c.JSON(http.StatusOK, u)
}

func (s *APIV1Service) GetUserByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := util.ConvertStringToInt32(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedUserIDErr).SetInternal(err)
	}
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr).SetInternal(err)
	}
	u := convertUserFromStore(user)
	u.Email = ""
	return c.JSON(http.StatusOK, u)
}

func (s *APIV1Service) UpdateUser(c echo.Context) error {
	ctx := c.Request().Context()
	userID, err := util.ConvertStringToInt32(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedUserIDErr).SetInternal(err)
	}
	curUserID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr)
	}
	curUser, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &curUserID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if curUser == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr).SetInternal(err)
	} else if curUser.Role != types.RoleHost && curUserID != userID {
		return echo.NewHTTPError(http.StatusForbidden, errmsg.UnauthorizedErr).SetInternal(err)
	}
	req := &UpdateUserRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedUpdateUserRequestErr).SetInternal(err)
	}
	if err := req.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidUserUpdateErr).SetInternal(err)
	}
	timestamp := time.Now().Unix()
	userUpdate := &store.UpdateUser{
		ID:               userID,
		UpdatedTimestamp: &timestamp,
	}
	if req.RowStatus != nil {
		rowStatus := types.RowStatus(req.RowStatus.String())
		userUpdate.RowStatus = &rowStatus
	}
	if req.Username != nil {
		if !usernameRegex.MatchString(strings.ToLower(*req.Username)) {
			return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidUsernameErr).SetInternal(err)
		}
		userUpdate.Username = req.Username
	}
	if req.Email != nil {
		userUpdate.Email = req.Email
	}
	if req.Nickname != nil {
		userUpdate.Nickname = req.Nickname
	}
	if req.Password != nil {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGeneratePasswordHashErr).SetInternal(err)
		}
		passwordHashStr := string(passwordHash)
		userUpdate.PasswordHash = &passwordHashStr
	}
	if req.AvatarURL != nil {
		userUpdate.AvatarURL = req.AvatarURL
	}
	user, err := s.Store.UpdateUser(ctx, userUpdate)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUpdateUserErr).SetInternal(err)
	}
	list, err := s.Store.ListUserSettings(ctx, &store.FindUserSetting{
		UserID: &userID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserSettings).SetInternal(err)
	}
	var userSettingList []*UserSetting
	for _, us := range list {
		userSettingList = append(userSettingList, convertUserSettingFromStore(us))
	}
	u := convertUserFromStore(user)
	u.UserSettingList = userSettingList
	return c.JSON(http.StatusOK, u)
}

func (s *APIV1Service) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	curUserID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr)
	}
	curUser, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &curUserID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetUserErr).SetInternal(err)
	}
	if curUser == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UserNotFoundErr).SetInternal(err)
	} else if curUser.Role != types.RoleHost {
		return echo.NewHTTPError(http.StatusForbidden, errmsg.UnauthorizedErr).SetInternal(err)
	}
	userID, err := util.ConvertStringToInt32(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedUserIDErr).SetInternal(err)
	}
	if err := s.Store.DeleteUser(ctx, &store.DeleteUser{
		ID: userID,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToDeleteUserErr).SetInternal(err)
	}
	return c.JSON(http.StatusOK, true)
}

func convertUserSettingFromStore(userSetting *store.UserSetting) *UserSetting {
	return &UserSetting{
		UserID: userSetting.UserID,
		Key:    types.UserSettingKey(userSetting.Key),
		Value:  userSetting.Value,
	}
}

func convertUserFromStore(user *store.User) *User {
	return &User{
		ID:           user.ID,
		RowStatus:    user.RowStatus,
		CreatedTs:    user.CreatedTimestamp,
		UpdatedTs:    user.UpdatedTimestamp,
		Username:     user.Username,
		Role:         user.Role,
		Email:        user.Email,
		Nickname:     user.Nickname,
		PasswordHash: user.PasswordHash,
		AvatarURL:    user.AvatarURL,
	}
}

func (create CreateUserRequest) Validate() error {
	if len(create.Username) < 3 {
		return errors.New("username is too short, minimum length is 3")
	}
	if len(create.Username) > 32 {
		return errors.New("username is too long, maximum length is 32")
	}
	if len(create.Password) < 3 {
		return errors.New("password is too short, minimum length is 3")
	}
	if len(create.Password) > 512 {
		return errors.New("password is too long, maximum length is 512")
	}
	if len(create.Nickname) > 64 {
		return errors.New("nickname is too long, maximum length is 64")
	}
	if create.Email != "" {
		if len(create.Email) > 256 {
			return errors.New("email is too long, maximum length is 256")
		}
		if !util.ValidateEmail(create.Email) {
			return errors.New("invalid email format")
		}
	}
	return nil
}

func (update UpdateUserRequest) Validate() error {
	if update.Username != nil && len(*update.Username) < 3 {
		return errors.New("username is too short, minimum length is 3")
	}
	if update.Username != nil && len(*update.Username) > 32 {
		return errors.New("username is too long, maximum length is 32")
	}
	if update.Password != nil && len(*update.Password) < 3 {
		return errors.New("password is too short, minimum length is 3")
	}
	if update.Password != nil && len(*update.Password) > 512 {
		return errors.New("password is too long, maximum length is 512")
	}
	if update.Nickname != nil && len(*update.Nickname) > 64 {
		return errors.New("nickname is too long, maximum length is 64")
	}
	if update.AvatarURL != nil {
		if len(*update.AvatarURL) > 2<<20 {
			return errors.New("avatar is too large, maximum is 2MB")
		}
	}
	if update.Email != nil && *update.Email != "" {
		if len(*update.Email) > 256 {
			return errors.New("email is too long, maximum length is 256")
		}
		if !util.ValidateEmail(*update.Email) {
			return errors.New("invalid email format")
		}
	}
	return nil
}
