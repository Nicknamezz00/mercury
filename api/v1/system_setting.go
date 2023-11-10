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
	"github.com/Nicknamezz00/mercury/internal/errmsg"
	"github.com/Nicknamezz00/mercury/internal/types"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"net/http"
)

const (
	SystemSettingServerIDName                           types.SystemSettingName = "server-id"
	SystemSettingSecretSessionName                      types.SystemSettingName = "secret-session"
	SystemSettingAllowSignUpName                        types.SystemSettingName = "allow-signup"
	SystemSettingDisablePasswordLoginName               types.SystemSettingName = "disable-password-login"
	SystemSettingDisablePublicMessagesName              types.SystemSettingName = "disable-public-messages"
	SystemSettingMaxUploadSizeMiBName                   types.SystemSettingName = "max-upload-size-mib"
	SystemSettingAdditionalStyleName                    types.SystemSettingName = "additional-style"
	SystemSettingAdditionalScriptName                   types.SystemSettingName = "additional-script"
	SystemSettingCustomizedProfileName                  types.SystemSettingName = "customized-profile"
	SystemSettingStorageServiceIDName                   types.SystemSettingName = "storage-service-id"
	SystemSettingLocalStoragePathName                   types.SystemSettingName = "local-storage-path"
	SystemSettingMessageDisplayWithUpdatedTimestampName types.SystemSettingName = "message-display-with-updated-timestamp"
	SystemSettingAutoBackupIntervalName                 types.SystemSettingName = "auto-backup-interval"
)

type CustomizedProfile struct {
	Name        string `json:"name"`
	LogoURL     string `json:"logoURL"`
	Description string `json:"description"`
	Locale      string `json:"locale"`
	Appearance  string `json:"appearance"`
	ExternalURL string `json:"externalURL"`
}

type SystemSetting struct {
	Name        types.SystemSettingName `json:"name"`
	Value       string                  `json:"value"`
	Description string                  `json:"description"`
}

type UpsertSystemSettingRequest struct {
	Name        types.SystemSettingName `json:"name"`
	Value       string                  `json:"value"`
	Description string                  `json:"description"`
}

const systemSettingUnmarshalError = `failed to unmarshal value from system setting "%v"`

func (r UpsertSystemSettingRequest) Validate() interface{} {
	switch settingName := r.Name; settingName {
	case SystemSettingServerIDName:
		return errors.Errorf("updating %v is not allowed", settingName)
	case SystemSettingAllowSignUpName:
		var value bool
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingDisablePublicMessagesName:
		var value bool
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingMaxUploadSizeMiBName:
		var value int
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingAdditionalStyleName:
		var value string
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingAdditionalScriptName:
		var value string
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingCustomizedProfileName:
		customizedProfile := CustomizedProfile{
			Name:        "mercury",
			LogoURL:     "",
			Description: "",
			Locale:      "en",
			Appearance:  "system",
			ExternalURL: "",
		}
		if err := json.Unmarshal([]byte(r.Value), &customizedProfile); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingStorageServiceIDName:
		// Note: 0 is the default value(database) for storage service ID.
		value := 0
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
		return nil
	case SystemSettingLocalStoragePathName:
		value := ""
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	case SystemSettingAutoBackupIntervalName:
		var value int
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
		if value < 0 {
			return errors.New("must be positive")
		}
	case SystemSettingMessageDisplayWithUpdatedTimestampName:
		var value bool
		if err := json.Unmarshal([]byte(r.Value), &value); err != nil {
			return errors.Errorf(systemSettingUnmarshalError, settingName)
		}
	default:
		return errors.New(errmsg.InvalidSystemSettingErr)
	}
	return nil
}

func (s *APIV1Service) registerSystemSettingRoutes(g *echo.Group) {
	g.GET("/system/setting", s.GetSystemSettingList)
	g.POST("/system/setting", s.CreateSystemSetting)
}

func (s *APIV1Service) GetSystemSettingList(c echo.Context) error {
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
	if user == nil || user.Role != types.RoleHost {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UnauthorizedErr)
	}
	list, err := s.Store.ListSystemSettings(ctx, &store.FindSystemSetting{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToGetSystemSettingsErr).SetInternal(err)
	}
	slist := make([]*SystemSetting, 0, len(list))
	for _, setting := range list {
		slist = append(slist, convertSystemSettingFromStore(setting))
	}
	return c.JSON(http.StatusOK, slist)
}

func (s *APIV1Service) CreateSystemSetting(c echo.Context) error {
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
	if user == nil || user.Role != types.RoleHost {
		return echo.NewHTTPError(http.StatusUnauthorized, errmsg.UnauthorizedErr)
	}
	systemSettingUpsert := &UpsertSystemSettingRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(systemSettingUpsert); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.MalformedSystemSettingRequestErr).SetInternal(err)
	}
	if err := systemSettingUpsert.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidSystemSettingErr).SetInternal(err)
	}

	if systemSettingUpsert.Name == SystemSettingDisablePasswordLoginName {
		var disable bool
		if err := json.Unmarshal([]byte(systemSettingUpsert.Value), &disable); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errmsg.InvalidSystemSettingErr).SetInternal(err)
		}
		// TODO: disable
	}

	systemSetting, err := s.Store.UpsertSystemSetting(ctx, &store.SystemSetting{
		Name:        systemSettingUpsert.Name.String(),
		Value:       systemSettingUpsert.Value,
		Description: systemSettingUpsert.Description,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, errmsg.FailedToUpsertSystemSettingErr).SetInternal(err)
	}
	return c.JSON(http.StatusOK, convertSystemSettingFromStore(systemSetting))
}

func convertSystemSettingFromStore(systemSetting *store.SystemSetting) *SystemSetting {
	return &SystemSetting{
		Name:        types.SystemSettingName(systemSetting.Name),
		Value:       systemSetting.Value,
		Description: systemSetting.Description,
	}
}
