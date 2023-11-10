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

package errmsg

const (
	UserNotFoundErr                   = "user not found"
	FailedToGetUserErr                = "failed to get user"
	FailedToCreateUserErr             = "failed to create user"
	FailedToGetSystemSettingsErr      = "failed to get system settings"
	UnauthorizedErr                   = "unauthorized"
	MalformedSystemSettingRequestErr  = "malformed post system setting request"
	MalformedSignInRequestErr         = "malformed sign in request"
	MalformedSignUpRequestErr         = "malformed sign up request"
	InvalidSystemSettingErr           = "invalid system setting"
	InvalidUsernameErr                = "invalid username"
	FailedToUpsertSystemSettingErr    = "failed to upsert system setting"
	FailedToUpsertUserSettingErr      = "failed to upsert user setting"
	FailedToUnmarshalSystemSettingErr = "failed to unmarshal system setting"
	IncorrectCredentialsErr           = "incorrect credentials, please try again"
	FailedToGenerateAccessTokenErr    = "failed to generate access token"
	FailedToUpsertAccessTokenErr      = "failed to upsert access token"
	FailedToGetAccessTokensErr        = "failed to get access tokens"
	FailedToGeneratePasswordHashErr   = "failed to generate password hash"
	SignUpDisabledErr                 = "failed to sign up, disabled"
	LoginDisabledErr                  = "failed to login, password login disabled"
)
