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

package store

import (
	"context"
	"database/sql"
	"github.com/Nicknamezz00/mercury/internal/types"
	storepb "github.com/Nicknamezz00/mercury/proto/gen/store"
)

type Driver interface {
	GetDB() *sql.DB
	Close() error
	Migrate(ctx context.Context) error
	Vacuum(ctx context.Context) error
	BackupTo(ctx context.Context, filename string) error

	// User methods.
	CreateUser(ctx context.Context, create *User) (*User, error)
	UpdateUser(ctx context.Context, update *UpdateUser) (*User, error)
	ListUsers(ctx context.Context, find *FindUser) ([]*User, error)
	DeleteUser(ctx context.Context, delete *DeleteUser) error

	// User setting methods.
	UpsertUserSetting(ctx context.Context, upsert *UserSetting) (*UserSetting, error)
	ListUserSettings(ctx context.Context, find *FindUserSetting) ([]*UserSetting, error)
	UpsertUserSettingV1(ctx context.Context, upsert *storepb.UserSetting) (*storepb.UserSetting, error)
	ListUserSettingsV1(ctx context.Context, find *FindUserSettingV1) ([]*storepb.UserSetting, error)

	// Message methods.
	CreateMessage(ctx context.Context, create *Message) (*Message, error)
	ListMessages(ctx context.Context, find *FindMessage) ([]*Message, error)
	UpdateMessage(ctx context.Context, update *UpdateMessage) error
	DeleteMessage(ctx context.Context, delete *DeleteMessage) error
	FindMessagesVisibilityList(ctx context.Context, messageIDs []int32) ([]types.Visibility, error)

	// MessageOrganizer methods.
	UpsertMessageOrganizer(ctx context.Context, upsert *MessageOrganizer) (*MessageOrganizer, error)
	ListMessageOrganizer(ctx context.Context, find *FindMessageOrganizer) ([]*MessageOrganizer, error)
	DeleteMessageOrganizer(ctx context.Context, delete *DeleteMessageOrganizer) error

	// Resource methods.
	CreateResource(ctx context.Context, create *Resource) (*Resource, error)
	ListResources(ctx context.Context, find *FindResource) ([]*Resource, error)
	UpdateResource(ctx context.Context, update *UpdateResource) (*Resource, error)
	DeleteResource(ctx context.Context, delete *DeleteResource) error

	// MessageRelation methods.
	UpsertMessageRelation(ctx context.Context, create *MessageRelation) (*MessageRelation, error)
	ListMessageRelations(ctx context.Context, find *FindMessageRelation) ([]*MessageRelation, error)
	DeleteMessageRelation(ctx context.Context, delete *DeleteMessageRelation) error

	// Tag methods.
	UpsertTag(ctx context.Context, upsert *Tag) (*Tag, error)
	ListTags(ctx context.Context, find *FindTag) ([]*Tag, error)
	DeleteTag(ctx context.Context, delete *DeleteTag) error

	// SystemSetting methods.
	UpsertSystemSetting(ctx context.Context, upsert *SystemSetting) (*SystemSetting, error)
	ListSystemSettings(ctx context.Context, find *FindSystemSetting) ([]*SystemSetting, error)
}
