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
	"fmt"
	storepb "github.com/Nicknamezz00/mercury/proto/gen/store"
)

type UserSetting struct {
	UserID int32
	Key    string
	Value  string
}

type FindUserSetting struct {
	UserID *int32
	Key    string
}

func (s *Store) GetUserAccessTokens(ctx context.Context, userID int32) ([]*storepb.AccessTokensUserSetting_AccessToken, error) {
	us, err := s.GetUserSettingV1(ctx, &FindUserSettingV1{
		UserID: &userID,
		Key:    storepb.UserSettingKey_USER_SETTING_ACCESS_TOKENS,
	})
	if err != nil {
		return nil, err
	}
	if us == nil {
		return []*storepb.AccessTokensUserSetting_AccessToken{}, nil
	}
	accessTokensUserSetting := us.GetAccessTokens()
	return accessTokensUserSetting.AccessTokens, nil
}

func (s *Store) UpsertUserSetting(ctx context.Context, upsert *UserSetting) (*UserSetting, error) {
	userSetting, err := s.driver.UpsertUserSetting(ctx, upsert)
	if err != nil {
		return nil, err
	}
	cacheKey := getUserSettingCacheKey(userSetting.UserID, userSetting.Key)
	s.userSettingCache.Store(cacheKey, userSetting)
	return userSetting, nil
}

func (s *Store) ListUserSettings(ctx context.Context, find *FindUserSetting) ([]*UserSetting, error) {
	userSettingList, err := s.driver.ListUserSettings(ctx, find)
	if err != nil {
		return nil, err
	}
	for _, userSetting := range userSettingList {
		cacheKey := getUserSettingCacheKey(userSetting.UserID, userSetting.Key)
		s.userSettingCache.Store(cacheKey, userSetting)
	}
	return userSettingList, nil
}

func (s *Store) GetUserSetting(ctx context.Context, find *FindUserSetting) (*UserSetting, error) {
	if find.UserID != nil {
		cacheKey := getUserSettingCacheKey(*find.UserID, find.Key)
		if cache, ok := s.userSettingCache.Load(cacheKey); ok {
			return cache.(*UserSetting), nil
		}
	}
	list, err := s.ListUserSettings(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	userSetting := list[0]
	return userSetting, nil
}

func getUserSettingCacheKey(id int32, key string) string {
	return fmt.Sprintf("%d-%s", id, key)
}

func vacuumUserSetting(ctx context.Context, tx *sql.Tx) error {
	stmt := `
		DELETE FROM
			user_setting
		WHERE
			user_id NOT IN (SELECT id FROM user)
	`
	if _, err := tx.ExecContext(ctx, stmt); err != nil {
		return err
	}
	return nil
}

type FindUserSettingV1 struct {
	UserID *int32
	Key    storepb.UserSettingKey
}

func (s *Store) UpsertUserSettingV1(ctx context.Context, upsert *storepb.UserSetting) (*storepb.UserSetting, error) {
	userSettingMessage, err := s.driver.UpsertUserSettingV1(ctx, upsert)
	if err != nil {
		return nil, err
	}
	cacheKey := getUserSettingV1CacheKey(userSettingMessage.UserId, userSettingMessage.Key.String())
	s.userSettingCache.Store(cacheKey, userSettingMessage)
	return userSettingMessage, nil
}

func (s *Store) ListUserSettingsV1(ctx context.Context, find *FindUserSettingV1) ([]*storepb.UserSetting, error) {
	userSettingList, err := s.driver.ListUserSettingsV1(ctx, find)
	if err != nil {
		return nil, err
	}
	for _, userSetting := range userSettingList {
		s.userSettingCache.Store(getUserSettingV1CacheKey(userSetting.UserId, userSetting.Key.String()), userSetting)
	}
	return userSettingList, nil
}

func (s *Store) GetUserSettingV1(ctx context.Context, find *FindUserSettingV1) (*storepb.UserSetting, error) {
	if find.UserID != nil {
		if cache, ok := s.userSettingCache.Load(getUserSettingV1CacheKey(*find.UserID, find.Key.String())); ok {
			return cache.(*storepb.UserSetting), nil
		}
	}

	list, err := s.ListUserSettingsV1(ctx, find)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, nil
	}

	userSetting := list[0]
	s.userSettingCache.Store(getUserSettingV1CacheKey(userSetting.UserId, userSetting.Key.String()), userSetting)
	return userSetting, nil
}

func getUserSettingV1CacheKey(userID int32, key string) any {
	return fmt.Sprintf("%d-%s-v1", userID, key)
}
