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
	"github.com/Nicknamezz00/mercury/api/auth"
	"github.com/Nicknamezz00/mercury/common/util"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

const (
	// The key name used to store user id in the context
	// user id is extracted from the jwt token subject field.
	userIDContextKey = "user-id"
)

func JWTMiddleware(s *APIV1Service, next echo.HandlerFunc, secret string) echo.HandlerFunc {
	return func(c echo.Context) error {
		//ctx := c.Request().Context()
		path := c.Request().URL.Path
		method := c.Request().Method

		if s.defaultAuthSkipper(c) {
			return next(c)
		}
		if util.HasPrefixes(path, "/api/v1/ping", "/api/v1/idp", "/api/v1/status", "/api/v1/user") && path != "/api/v1/user/me" && method == http.MethodGet {
			return next(c)
		}

		accessToken := findAccessToken(c)
		if accessToken == "" {
			// Allow public endpoints.
			if util.HasPrefixes(path, "/o") {
				return next(c)
			}
			if util.HasPrefixes(path, "/api/v1/message") && method == http.MethodGet {
				return next(c)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "missing access token")
		}

		uid := 0
		//uid, err := getUserIDFromAccessToken(accessToken, secret)
		//if err != nil {
		//	removeAccessTokenAndCookies(c)
		//	return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired access token")
		//}
		//
		//if !validateAccessToken(accessToken, accessTokens) {
		//	removeAccessTokenAndCookies(c)
		//	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid access token.")
		//}

		c.Set(userIDContextKey, uid)
		return next(c)
	}
}

func findAccessToken(c echo.Context) string {
	// Check the HTTP request header first.
	accessToken, _ := extractTokenFromHeader(c)
	if accessToken == "" {
		// Check the cookie.
		cookie, _ := c.Cookie(auth.AccessTokenCookieName)
		if cookie != nil {
			accessToken = cookie.Value
		}
	}
	return accessToken
}

func extractTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}

func (s *APIV1Service) defaultAuthSkipper(c echo.Context) bool {
	path := c.Path()
	return util.HasPrefixes(path, "/api/v1/auth")
}

func getUserIDFromAccessToken(accessToken, secret string) (int32, error) {
	claims := &auth.ClaimsMessage{}
	_, err := jwt.ParseWithClaims(accessToken, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, errors.Errorf("unexpected access token signing method=%v, expect %v", t.Header["alg"], jwt.SigningMethodHS256)
		}
		if kid, ok := t.Header["kid"].(string); ok {
			if kid == "v1" {
				return []byte(secret), nil
			}
		}
		return nil, errors.Errorf("unexpected access token kid=%v", t.Header["kid"])
	})
	if err != nil {
		return 0, errors.Wrap(err, "Invalid or expired access token")
	}
	// We either have a valid access token or we will attempt to generate new access token.
	userID, err := util.ConvertStringToInt32(claims.Subject)
	if err != nil {
		return 0, errors.Wrap(err, "Malformed ID in the token")
	}
	return userID, nil
}
