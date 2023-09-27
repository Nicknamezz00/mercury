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

type MessageOrganizer struct {
	MessageID int32
	UserID    int32
	Pinned    bool
}

type FindMessageOrganizer struct {
	MessageID int32
	UserID    int32
}

type DeleteMessageOrganizer struct {
	MessageID *int32
	UserID    *int32
}

func (s *Store) UpsertMessageOrganizer(ctx context.Context, upsert *MessageOrganizer) (*MessageOrganizer, error) {
	stmt := `
		INSERT INTO message_organizer (
			message_id, user_id, pinned                               
		)	
		VALUES (?, ?, ?)
		ON CONFLICT (message_id, user_id) DO UPDATE
    SET
   		pinned = EXCLUDED.pinned
	`
	if _, err := s.db.ExecContext(ctx, stmt, upsert.MessageID, upsert.UserID, upsert.Pinned); err != nil {
		return nil, err
	}
	m := upsert
	return m, nil
}

func (s *Store) GetMessageOrganizer(ctx context.Context, find *FindMessageOrganizer) (*MessageOrganizer, error) {
	where, args := []string{}, []any{}
	if find.MessageID != 0 {
		where, args = append(where, "message_id = ?"), append(args, find.MessageID)
	}
	if find.UserID != 0 {
		where, args = append(where, "user_id = ?"), append(args, find.UserID)
	}
	query := `
		SELECT
			message_id, user_id, pinned
		FROM
			message_organizer
		WHERE
	` + strings.Join(where, " AND ")
	rows := s.db.QueryRowContext(ctx, query, args...)
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, nil
	}
	m := &MessageOrganizer{}
	if err := rows.Scan(&m.MessageID, &m.UserID, &m.Pinned); err != nil {
		return nil, err
	}
	return m, nil
}

func (s *Store) DeleteMessageOrganizer(ctx context.Context, delete *DeleteMessageOrganizer) error {
	where, args := []string{}, []any{}
	if v := delete.MessageID; v != nil {
		where, args = append(where, "message_id = ?"), append(args, *v)
	}
	if v := delete.UserID; v != nil {
		where, args = append(where, "user_id = ?"), append(args, *v)
	}
	stmt := `DELETE FROM message_organizer WHERE ` + strings.Join(where, " AND ")
	if _, err := s.db.ExecContext(ctx, stmt, args...); err != nil {
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
	if _, err := tx.ExecContext(ctx, stmt); err != nil {
		return err
	}
	return nil
}
