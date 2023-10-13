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

package sqlite

import (
	"context"
	"database/sql"
	"github.com/Nicknamezz00/mercury/store"
	"strings"
)

func (d *DB) CreateUser(ctx context.Context, create *store.User) (*store.User, error) {
	fields := []string{"`username`", "`role`", "`email`", "`nickname`", "`password_hash`"}
	placeholder := []string{"?", "?", "?", "?", "?"}
	args := []any{create.Username, create.Role, create.Email, create.Nickname, create.PasswordHash}

	if create.AvatarURL != "" {
		fields = append(fields, "`avatar_url`")
		placeholder = append(placeholder, "?")
		args = append(args, create.AvatarURL)
	}
	if create.RowStatus != "" {
		fields = append(fields, "`row_status`")
		placeholder = append(placeholder, "?")
		args = append(args, create.RowStatus)
	}
	if create.CreatedTimestamp != 0 {
		fields = append(fields, "`created_timestamp`")
		placeholder = append(placeholder, "?")
		args = append(args, create.CreatedTimestamp)
	}
	if create.UpdatedTimestamp != 0 {
		fields = append(fields, "`updated_timestamp`")
		placeholder = append(placeholder, "?")
		args = append(args, create.UpdatedTimestamp)
	}
	if create.ID != 0 {
		fields = append(fields, "`id`")
		placeholder = append(placeholder, "?")
		args = append(args, create.ID)
	}
	stmt := `
		INSERT INTO user (` + strings.Join(fields, ", ") + `) 
		VALUES (` + strings.Join(placeholder, ", ") + `) 
		RETURNING id, avatar_url, created_timestamp, updated_timestamp, row_status
		`
	err := d.db.QueryRowContext(ctx, stmt, args...).Scan(&create.ID, &create.AvatarURL, &create.CreatedTimestamp, &create.UpdatedTimestamp, &create.RowStatus)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (d *DB) UpdateUser(ctx context.Context, update *store.UpdateUser) (*store.User, error) {
	var (
		args []any
		set  []string
	)
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
	user := &store.User{}
	err := d.db.QueryRowContext(ctx, query, args...).Scan(&user.ID, &user.Username, &user.Role, &user.Email, &user.Nickname, &user.PasswordHash, &user.AvatarURL, &user.CreatedTimestamp, &user.UpdatedTimestamp, &user.RowStatus)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *DB) ListUsers(ctx context.Context, find *store.FindUser) ([]*store.User, error) {
	var args []any
	where := []string{"1 = 1"}
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
			id,
			username,
			role,
			email,
			nickname,
			password_hash,
			avatar_url,
			created_timestamp,
			updated_timestamp,
			row_status
		FROM user
		WHERE ` + strings.Join(where, " AND ") + `
		ORDER BY created_timestamp DESC, row_status DESC
	`
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	list := make([]*store.User, 0)
	for rows.Next() {
		var user store.User
		err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.Email, &user.Nickname, &user.PasswordHash, &user.AvatarURL, &user.CreatedTimestamp, &user.UpdatedTimestamp, &user.RowStatus)
		if err != nil {
			return nil, err
		}
		list = append(list, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) DeleteUser(ctx context.Context, delete *store.DeleteUser) error {
	result, err := d.db.ExecContext(ctx, `DELETE FROM user WHERE id = ?`, delete.ID)
	if err != nil {
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		return err
	}
	if err := d.Vacuum(ctx); err != nil {
		return err
	}
	return nil
}
