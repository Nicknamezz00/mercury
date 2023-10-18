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
	"github.com/Nicknamezz00/mercury/common/log"
	"github.com/Nicknamezz00/mercury/internal/constants"
	"github.com/Nicknamezz00/mercury/internal/types"
	"github.com/Nicknamezz00/mercury/server/profile"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type SystemStatus struct {
	Host    *User           `json:"host"`
	Profile profile.Profile `json:"profile"`
	DBSize  int64           `json:"dbSize"`

	AllowSignUp                        bool              `json:"allowSignUp"`
	DisablePasswordLogin               bool              `json:"disablePasswordLogin"`
	DisablePublicMessages              bool              `json:"disablePublicMessages"`
	MaxUploadSizeMiB                   int               `json:"maxUploadSizeMiB"`
	AutoBackupInterval                 int               `json:"autoBackupInterval"`
	AdditionalStyle                    string            `json:"additionalStyle"`
	AdditionalScript                   string            `json:"additionalScript"`
	CustomizedProfile                  CustomizedProfile `json:"customizedProfile"`
	StorageServiceId                   int32             `json:"storageServiceId"`
	LocalStoragePath                   string            `json:"localStoragePath"`
	MessageDisplayWithUpdatedTimestamp bool              `json:"messageDisplayWithUpdatedTimestamp"`
}

func (s *APIV1Service) registerSystemRoutes(g *echo.Group) {
	g.GET("/ping", s.PingSystem)
	g.GET("/status", s.GetSystemStatus)
	g.GET("/system/vacuum", s.ExecVacuum)
}

func (*APIV1Service) PingSystem(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

func (s *APIV1Service) GetSystemStatus(c echo.Context) error {
	ctx := c.Request().Context()
	systemStatus := SystemStatus{
		Profile:          *s.Profile,
		AllowSignUp:      true,
		MaxUploadSizeMiB: 32,
		CustomizedProfile: CustomizedProfile{
			Name:       "mercury",
			Locale:     "en",
			Appearance: constants.APPEARANCE_SYSTEM,
		},
		StorageServiceId: DefaultStorage,
		LocalStoragePath: constants.LOCAL_STORAGE_PATH,
	}
	hostUserType := types.RoleHost
	hostUser, err := s.Store.GetUser(ctx, &store.FindUser{
		Role: &hostUserType,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to find host user").SetInternal(err)
	}
	if hostUser != nil {
		systemStatus.Host = &User{
			ID: hostUser.ID,
		}
	}

	systemSettingList, err := s.Store.ListSystemSettings(ctx, &store.FindSystemSetting{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to find system setting list").SetInternal(err)
	}
	for _, s := range systemSettingList {
		if s.Name == SystemSettingServerIDName.String() || s.Name == SystemSettingSecretSessionName.String() {
			continue
		}
		var value any
		err := json.Unmarshal([]byte(s.Value), &value)
		if err != nil {
			log.Warn("failed to unmarshal system setting value", zap.String("setting name", s.Name))
			continue
		}

		switch s.Name {
		case SystemSettingAllowSignUpName.String():
			systemStatus.AllowSignUp = value.(bool)
		case SystemSettingDisablePasswordLoginName.String():
			systemStatus.DisablePasswordLogin = value.(bool)
		case SystemSettingDisablePublicMessagesName.String():
			systemStatus.DisablePublicMessages = value.(bool)
		case SystemSettingMaxUploadSizeMiBName.String():
			systemStatus.MaxUploadSizeMiB = int(value.(float64))
		case SystemSettingAdditionalStyleName.String():
			systemStatus.AdditionalStyle = value.(string)
		case SystemSettingAdditionalScriptName.String():
			systemStatus.AdditionalScript = value.(string)
		case SystemSettingCustomizedProfileName.String():
			customizedProfile := CustomizedProfile{}
			if err := json.Unmarshal([]byte(s.Value), &customizedProfile); err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to unmarshal system setting customized profile value").SetInternal(err)
			}
		case SystemSettingLocalStoragePathName.String():
			systemStatus.LocalStoragePath = value.(string)
		case SystemSettingMessageDisplayWithUpdatedTimestampName.String():
			systemStatus.MessageDisplayWithUpdatedTimestamp = value.(bool)
		case SystemSettingAutoBackupIntervalName.String():
			systemStatus.AutoBackupInterval = int(value.(float64))
		default:
			log.Warn("unknown syste setting name", zap.String("setting name", s.Name))
		}
	}
	return c.JSON(http.StatusOK, systemStatus)
}

func (s *APIV1Service) ExecVacuum(c echo.Context) error {
	ctx := c.Request().Context()
	uid, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "missing user in session")
	}
	user, err := s.Store.GetUser(ctx, &store.FindUser{
		ID: &uid,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to find user").SetInternal(err)
	}
	if user == nil || user.Role != types.RoleHost {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}
	if err := s.Store.Vacuum(ctx); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to vacuum database").SetInternal(err)
	}
	return c.JSON(http.StatusOK, true)
}
