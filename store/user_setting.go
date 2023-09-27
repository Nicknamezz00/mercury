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
	"strings"
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

func (s *Store) UpsertUserSetting(ctx context.Context, upsert *UserSetting) (*UserSetting, error) {
	stmt := `
		INSERT INTO user_setting (user_id, key, value) 
		VALUES (?, ?, ?)
		ON CONFLICT (user_id, key) DO UPDATE 
		SET value = EXCLUDED.value
	`
	if _, err := s.db.ExecContext(ctx, stmt, upsert.UserID, upsert.Key, upsert.Value); err != nil {
		return nil, err
	}
	userSetting := upsert
	s.userSettingCache.Store(userSetting.UserID, userSetting)
	return userSetting, nil
}

func (s *Store) ListUserSettings(ctx context.Context, find *FindUserSetting) ([]*UserSetting, error) {
	where, args := []string{"1 = 1"}, []any{}
	if v := find.Key; v != "" {
		where, args = append(where, "key = ?"), append(args, v)
	}
	if v := find.UserID; v != nil {
		where, args = append(where, "user_id = ?"), append(args, *v)
	}
	query := `
		SELECT
			user_id, key, value
		FROM user_setting
		WHERE ` + strings.Join(where, " AND ")
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	userSettingList := make([]*UserSetting, 0)
	for rows.Next() {
		var userSetting UserSetting
		if err := rows.Scan(
			&userSetting.UserID,
			&userSetting.Key,
			&userSetting.Value,
		); err != nil {
			return nil, err
		}
		userSettingList = append(userSettingList, &userSetting)
	}
	if err := rows.Err(); err != nil {
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
	cacheKey := getUserSettingCacheKey(userSetting.UserID, userSetting.Key)
	s.userSettingCache.Store(cacheKey, userSetting)
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
