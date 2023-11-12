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

package types

// RowStatus type
type RowStatus string

const (
	Normal   RowStatus = "NORMAL"
	Archived RowStatus = "ARCHIVED"
)

func (r RowStatus) String() string {
	return string(r)
}

// Visibility type
type Visibility string

const (
	Public    Visibility = "PUBLIC"
	Protected Visibility = "PROTECTED"
	Private   Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	switch v {
	case Public:
		return "PUBLIC"
	case Protected:
		return "PROTECTED"
	case Private:
		return "PRIVATE"
	}
	return "PRIVATE"
}

// MessageRelationType type
type MessageRelationType string

const (
	MessageRelationReference  MessageRelationType = "REFERENCE"
	MessageRelationAdditional MessageRelationType = "ADDITIONAL"
)

// Role type
type Role string

const (
	RoleHost  Role = "HOST"
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

func (e Role) String() string {
	switch e {
	case RoleHost:
		return "HOST"
	case RoleAdmin:
		return "ADMIN"
	case RoleUser:
		return "USER"
	}
	return "USER"
}

// SystemSettingName type
type SystemSettingName string

func (s SystemSettingName) String() string {
	return string(s)
}

// UserSetting type
type UserSettingKey string

const (
	UserSettingLocaleKey UserSettingKey = "locale"
	// UserSettingAppearanceKey is the key type for user appearance.
	UserSettingAppearanceKey UserSettingKey = "appearance"
	// UserSettingMessageVisibilityKey is the key type for user preference memo default visibility.
	UserSettingMessageVisibilityKey UserSettingKey = "message-visibility"
)

func (key UserSettingKey) String() string {
	switch key {
	case UserSettingLocaleKey:
		return "locale"
	case UserSettingAppearanceKey:
		return "appearance"
	case UserSettingMessageVisibilityKey:
		return "message-visibility"
	}
	return ""
}

var (
	UserSettingLocaleValue = []string{
		"de",
		"en",
		"es",
		"fr",
		"hi",
		"hr",
		"it",
		"ja",
		"ko",
		"nl",
		"pl",
		"pt-BR",
		"ru",
		"sl",
		"sv",
		"tr",
		"uk",
		"vi",
		"zh-Hans",
		"zh-Hant",
	}
	UserSettingAppearanceValue     = []string{"system", "light", "dark"}
	UserSettingMemoVisibilityValue = []Visibility{Private, Protected, Public}
)
