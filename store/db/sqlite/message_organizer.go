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
	"fmt"
	"github.com/Nicknamezz00/mercury/store"
	"strings"
)

func (d *DB) UpsertMessageOrganizer(ctx context.Context, upsert *store.MessageOrganizer) (*store.MessageOrganizer, error) {
	stmt := `
		INSERT INTO message_organizer (
			message_id,
			user_id,
			pinned
		)
		VALUES (?, ?, ?)
		ON CONFLICT(message_id, user_id) DO UPDATE 
		SET
			pinned = EXCLUDED.pinned
	`
	if _, err := d.db.ExecContext(ctx, stmt, upsert.MessageID, upsert.UserID, upsert.Pinned); err != nil {
		return nil, err
	}
	return upsert, nil
}

func (d *DB) ListMessageOrganizer(ctx context.Context, find *store.FindMessageOrganizer) ([]*store.MessageOrganizer, error) {
	where, args := []string{"1 = 1"}, []any{}
	if find.MessageID != 0 {
		where = append(where, "message_id = ?")
		args = append(args, find.MessageID)
	}
	if find.UserID != 0 {
		where = append(where, "user_id = ?")
		args = append(args, find.UserID)
	}

	query := fmt.Sprintf(`
		SELECT
			message_id,
			user_id,
			pinned
		FROM message_organizer
		WHERE %s
	`, strings.Join(where, " AND "))
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var list []*store.MessageOrganizer
	for rows.Next() {
		messageOrganizer := &store.MessageOrganizer{}
		if err := rows.Scan(&messageOrganizer.MessageID, &messageOrganizer.UserID, &messageOrganizer.Pinned); err != nil {
			return nil, err
		}
		list = append(list, messageOrganizer)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) DeleteMessageOrganizer(ctx context.Context, delete *store.DeleteMessageOrganizer) error {
	var (
		where []string
		args  []any
	)
	if v := delete.MessageID; v != nil {
		where, args = append(where, "message_id = ?"), append(args, *v)
	}
	if v := delete.UserID; v != nil {
		where, args = append(where, "user_id = ?"), append(args, *v)
	}
	stmt := `DELETE FROM message_organizer WHERE ` + strings.Join(where, " AND ")
	if _, err := d.db.ExecContext(ctx, stmt, args...); err != nil {
		return err
	}
	return nil
}

func vacuumMessageOrganizer(ctx context.Context, tx *sql.Tx) error {
	stmt := `
	DELETE FROM 
		message_organizer 
	WHERE 
		message_id NOT IN (SELECT id FROM message)
		OR user_id NOT IN (SELECT id FROM user)
	`
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
