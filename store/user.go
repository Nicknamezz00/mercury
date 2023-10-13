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
	"github.com/Nicknamezz00/mercury/internal/types"
)

type User struct {
	ID               int32
	RowStatus        types.RowStatus
	CreatedTimestamp int64
	UpdatedTimestamp int64
	Username         string
	Role             types.Role
	Email            string
	Nickname         string
	PasswordHash     string
	AvatarURL        string
}

type FindUser struct {
	ID        *int32
	RowStatus *types.RowStatus
	Username  *string
	Role      *types.Role
	Email     *string
	Nickname  *string
}

type UpdateUser struct {
	ID               int32
	UpdatedTimestamp *int64
	RowStatus        *types.RowStatus
	Username         *string
	Role             *types.Role
	Email            *string
	Nickname         *string
	Password         *string
	AvatarURL        *string
	PasswordHash     *string
}

type DeleteUser struct {
	ID int32
}

func (s *Store) CreateUser(ctx context.Context, create *User) (*User, error) {
	user, err := s.driver.CreateUser(ctx, create)
	if err != nil {
		return nil, err
	}
	s.userCache.Store(user.ID, user)
	return user, nil
}

func (s *Store) UpdateUser(ctx context.Context, update *UpdateUser) (*User, error) {
	user, err := s.driver.UpdateUser(ctx, update)
	if err != nil {
		return nil, err
	}
	s.userCache.Store(user.ID, user)
	return user, nil
}

func (s *Store) ListUsers(ctx context.Context, find *FindUser) ([]*User, error) {
	list, err := s.driver.ListUsers(ctx, find)
	if err != nil {
		return nil, err
	}
	for _, user := range list {
		s.userCache.Store(user.ID, user)
	}
	return list, nil
}

func (s *Store) GetUser(ctx context.Context, find *FindUser) (*User, error) {
	if find.ID != nil {
		if cache, ok := s.userCache.Load(*find.ID); ok {
			return cache.(*User), nil
		}
	}
	list, err := s.ListUsers(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	user := list[0]
	s.userCache.Store(user.ID, user)
	return user, nil
}

func (s *Store) DeleteUser(ctx context.Context, delete *DeleteUser) error {
	err := s.driver.DeleteUser(ctx, delete)
	if err != nil {
		return err
	}
	s.userCache.Delete(delete.ID)
	return nil
}
