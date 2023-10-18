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
	RowStatus *types.RowStatus
}
