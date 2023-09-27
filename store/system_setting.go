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
	"strings"
)

type SystemSetting struct {
	Name        string
	Value       string
	Description string
}

type FindSystemSetting struct {
	Name string
}

func (s *Store) UpsertSystemSetting(ctx context.Context, upsert *SystemSetting) (*SystemSetting, error) {
	stmt := `
		INSERT INTO system_setting (
			name, value, description
		)
		VALUES (?, ?, ?)
		ON CONFLICT(name) DO UPDATE 
		SET
			value = EXCLUDED.value,
			description = EXCLUDED.description
	`
	if _, err := s.db.ExecContext(ctx, stmt, upsert.Name, upsert.Value, upsert.Description); err != nil {
		return nil, err
	}
	ss := upsert
	return ss, nil
}

func (s *Store) ListSystemSettings(ctx context.Context, find *FindSystemSetting) ([]*SystemSetting, error) {
	where, args := []string{"1 = 1"}, []any{}
	if find.Name != "" {
		where, args = append(where, "name = ?"), append(args, find.Name)
	}
	query := `
		SELECT
			name,
			value,
			description
		FROM system_setting
		WHERE ` + strings.Join(where, " AND ")
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	list := []*SystemSetting{}
	for rows.Next() {
		ss := &SystemSetting{}
		if err := rows.Scan(
			&ss.Name,
			&ss.Value,
			&ss.Description,
		); err != nil {
			return nil, err
		}
		list = append(list, ss)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for _, ss := range list {
		s.systemSettingCache.Store(ss.Name, ss)
	}
	return list, nil
}

func (s *Store) GetSystemSetting(ctx context.Context, find *FindSystemSetting) (*SystemSetting, error) {
	if find.Name != "" {
		if cache, ok := s.systemSettingCache.Load(find.Name); ok {
			return cache.(*SystemSetting), nil
		}
	}
	list, err := s.ListSystemSettings(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	ss := list[0]
	s.systemSettingCache.Store(ss.Name, ss)
	return ss, nil
}

func (s *Store) GetSystemSettingWithDefault(ctx *context.Context, settingName, defaultValue string) string {
	if setting, err := s.GetSystemSetting(*ctx, &FindSystemSetting{
		Name: settingName,
	}); err != nil && setting.Value != "" {
		return setting.Value
	}
	return defaultValue
}
