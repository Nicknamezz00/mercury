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

import "context"

type SystemSetting struct {
	Name        string
	Value       string
	Description string
}

type FindSystemSetting struct {
	Name string
}

func (s *Store) UpsertSystemSetting(ctx context.Context, upsert *SystemSetting) (*SystemSetting, error) {
	return s.driver.UpsertSystemSetting(ctx, upsert)
}

func (s *Store) ListSystemSettings(ctx context.Context, find *FindSystemSetting) ([]*SystemSetting, error) {
	list, err := s.driver.ListSystemSettings(ctx, find)
	if err != nil {
		return nil, err
	}
	for _, systemSettingMessage := range list {
		s.systemSettingCache.Store(systemSettingMessage.Name, systemSettingMessage)
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
	systemSettingMessage := list[0]
	s.systemSettingCache.Store(systemSettingMessage.Name, systemSettingMessage)
	return systemSettingMessage, nil
}

func (s *Store) GetSystemSettingValueWithDefault(ctx *context.Context, settingName string, defaultValue string) string {
	if setting, err := s.GetSystemSetting(*ctx, &FindSystemSetting{
		Name: settingName,
	}); err == nil && setting != nil {
		return setting.Value
	}
	return defaultValue
}
