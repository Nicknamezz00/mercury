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

type MessageResource struct {
	MessageID        int32
	ResourceID       int32
	CreatedTimestamp int64
	UpdatedTimestamp int64
}

type UpsertMessageResource struct {
	MessageID        int32
	ResourceID       int32
	CreatedTimestamp int64
	UpdatedTimestamp *int64
}

type FindMessageResource struct {
	MessageID  *int32
	ResourceID *int32
}

type DeleteMessageResource struct {
	MessageID  *int32
	ResourceID *int32
}

func (s *Store) UpsertMessageResource(ctx context.Context, upsert *UpsertMessageResource) (*MessageResource, error) {
	set := []string{"message_id", "resource_id"}
	args := []any{upsert.MessageID, upsert.ResourceID}
	placeholder := []string{"?", "?"}
	if v := upsert.UpdatedTimestamp; v != nil {
		set, args, placeholder = append(set, "updated_timestamp"), append(args, v), append(placeholder, "?")
	}
	query := `
		INSERT INTO message_resource (
			` + strings.Join(set, ",") + `                             
		)
		VALUES (` + strings.Join(placeholder, ",") + `)
		ON CONFLICT (message_id, resource_id) DO UPDATE
		SET
			updated_timestamp = EXCLUDED.updated_timestamp
		RETURNING message_id, resource_id, created_timestamp, updated_timestamp
	`
	mr := &MessageResource{}
	err := s.db.QueryRowContext(ctx, query, args...).Scan(&mr.ResourceID, &mr.MessageID, &mr.CreatedTimestamp, &mr.UpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	return mr, nil
}

func (s *Store) ListMessageResources(ctx context.Context, find *FindMessageResource) ([]*MessageResource, error) {
	where, args := []string{"1 = 1"}, []any{}
	if v := find.MessageID; v != nil {
		where, args = append(where, "message_id = ?"), append(args, *v)
	}
	if v := find.ResourceID; v != nil {
		where, args = append(where, "resource_id = ?"), append(args, *v)
	}
	query := `
		SELECT
			message_id, resource_id, created_timestamp, updated_timestamp
		FROM
			message_resource
		WHERE ` + strings.Join(where, " AND ") + `
		ORDER BY updated_timestamp DESC
	`
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	list := make([]*MessageResource, 0)
	for rows.Next() {
		var mr MessageResource
		if err := rows.Scan(&mr.ResourceID, &mr.MessageID, &mr.CreatedTimestamp, &mr.UpdatedTimestamp); err != nil {
			return nil, err
		}
		list = append(list, &mr)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Store) GetMessageResource(ctx context.Context, find *FindMessageResource) (*MessageResource, error) {
	list, err := s.ListMessageResources(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	m := list[0]
	return m, nil
}

func (s *Store) DeleteMessageResource(ctx context.Context, delete *DeleteMessageResource) error {
	where, args := []string{}, []any{}
	if v := delete.MessageID; v != nil {
		where, args = append(where, "message_id = ?"), append(args, *v)
	}
	if v := delete.ResourceID; v != nil {
		where, args = append(where, "resource_id = ?"), append(args, *v)
	}
	stmt := `DELETE FROM message_resource WHERE ` + strings.Join(where, " AND ")
	r, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err := r.RowsAffected(); err != nil {
		return err
	}
	return nil
}
