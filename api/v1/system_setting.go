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

import "github.com/Nicknamezz00/mercury/internal/types"

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
