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

type Role string

const (
	RoleHost  Role = "HOST"
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USEr"
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

type User struct {
	ID               int32
	RowStatus        RowStatus
	CreatedTimestamp int64
	UpdatedTimestamp int64
	Username         string
	Role             Role
	Email            string
	Nickname         string
	PasswordHash     string
	AvatarURL        string
}

type FindUser struct {
	ID        *int32
	RowStatus *RowStatus
	Username  *string
	Role      *Role
	Email     *string
	Nickname  *string
}

type UpdateUser struct {
	ID               int32
	UpdatedTimestamp *int64
	RowStatus        *RowStatus
	Username         *string
	Role             *Role
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
	stmt := `
    INSERT INTO user (
      username, role, email, nickname, password_hash
    )
    VALUES (?, ?, ?, ?, ?)
    RETURNING id, avatar_url, created_timestamp, updated_timestamp, row_status
  `
	if err := s.db.QueryRowContext(
		ctx,
		stmt,
		create.Username,
		create.Role,
		create.Email,
		create.Nickname,
		create.PasswordHash,
	).Scan(
		&create.ID,
		&create.AvatarURL,
		&create.CreatedTimestamp,
		&create.UpdatedTimestamp,
		&create.RowStatus,
	); err != nil {
		return nil, err
	}
	user := create
	s.userCache.Store(user.ID, user)
	return user, nil
}

func (s *Store) UpdateUser(ctx context.Context, update *UpdateUser) (*User, error) {
	set, args := []string{}, []any{}
	if v := update.UpdatedTimestamp; v != nil {
		set, args = append(set, "updated_timestamp = ?"), append(args, *v)
	}
	if v := update.RowStatus; v != nil {
		set, args = append(set, "row_status = ?"), append(args, *v)
	}
	if v := update.Username; v != nil {
		set, args = append(set, "username = ?"), append(args, *v)
	}
	if v := update.Email; v != nil {
		set, args = append(set, "email = ?"), append(args, *v)
	}
	if v := update.Nickname; v != nil {
		set, args = append(set, "nickname = ?"), append(args, *v)
	}
	if v := update.AvatarURL; v != nil {
		set, args = append(set, "avatar_url = ?"), append(args, *v)
	}
	if v := update.PasswordHash; v != nil {
		set, args = append(set, "password_hash = ?"), append(args, *v)
	}
	args = append(args, update.ID)

	query := `
		UPDATE user
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING id, username, role, email, nickname, password_hash, avatar_url, created_timestamp, updated_timestamp, row_status
	`
	user := &User{}
	if err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Role,
		&user.Email,
		&user.Nickname,
		&user.PasswordHash,
		&user.AvatarURL,
		&user.CreatedTimestamp,
		&user.UpdatedTimestamp,
		&user.RowStatus,
	); err != nil {
		return nil, err
	}

	s.userCache.Store(user.ID, user)
	return user, nil
}

func (s *Store) ListUsers(ctx context.Context, find *FindUser) ([]*User, error) {
	where, args := []string{"1 = 1"}, []any{}
	if v := find.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := find.Username; v != nil {
		where, args = append(where, "username = ?"), append(args, *v)
	}
	if v := find.Role; v != nil {
		where, args = append(where, "role = ?"), append(args, *v)
	}
	if v := find.Email; v != nil {
		where, args = append(where, "email = ?"), append(args, *v)
	}
	if v := find.Nickname; v != nil {
		where, args = append(where, "nickname = ?"), append(args, *v)
	}
	query := `
		SELECT
			id, username, role, email, nickname, password_hash, avatar_url, created_timestamp, updated_timestamp, row_status
		FROM
			user
		WHERE ` + strings.Join(where, " AND ") + `
		ORDER BY created_timestamp DESC, row_status DESC
	`
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	list := make([]*User, 0)
	for rows.Next() {
		var u User
		if err := rows.Scan(
			&u.ID,
			&u.Username,
			&u.Role,
			&u.Email,
			&u.Nickname,
			&u.PasswordHash,
			&u.AvatarURL,
			&u.CreatedTimestamp,
			&u.UpdatedTimestamp,
			&u.RowStatus,
		); err != nil {
			return nil, err
		}
		list = append(list, &u)
	}
	if err := rows.Err(); err != nil {
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
	result, err := s.db.ExecContext(ctx, `
		DELETE FROM user WHERE id = ?
	`, delete.ID)
	if err != nil {
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		return err
	}
	if err := s.Vacuum(ctx); err != nil {
		return err
	}
	s.userCache.Delete(delete.ID)
	return nil
}
