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

func (d *DB) CreateResource(ctx context.Context, create *store.Resource) (*store.Resource, error) {
	fields := []string{"`filename`", "`blob`", "`external_link`", "`type`", "`size`", "`creator_id`", "`internal_path`"}
	placeholder := []string{"?", "?", "?", "?", "?", "?", "?"}
	args := []any{create.Filename, create.Blob, create.ExternalLink, create.Type, create.Size, create.CreatorID, create.InternalPath}

	if create.ID != 0 {
		fields = append(fields, "`id`")
		placeholder = append(placeholder, "?")
		args = append(args, create.ID)
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
	if create.MessageID != nil {
		fields = append(fields, "`message_id`")
		placeholder = append(placeholder, "?")
		args = append(args, *create.MessageID)
	}

	stmt := `
	INSERT INTO resource (` + strings.Join(fields, ", ") + `)
	VALUES (` + strings.Join(placeholder, ", ") + `) 
	RETURNING id, created_timestamp, updated_timestamp
	`
	err := d.db.QueryRowContext(ctx, stmt, args...).Scan(&create.ID, &create.CreatedTimestamp, &create.UpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (d *DB) ListResources(ctx context.Context, find *store.FindResource) ([]*store.Resource, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := find.ID; v != nil {
		where, args = append(where, "id = ?"), append(args, *v)
	}
	if v := find.CreatorID; v != nil {
		where, args = append(where, "creator_id = ?"), append(args, *v)
	}
	if v := find.Filename; v != nil {
		where, args = append(where, "filename = ?"), append(args, *v)
	}
	if v := find.MessageID; v != nil {
		where, args = append(where, "message_id = ?"), append(args, *v)
	}
	if find.HasRelatedMessage {
		where = append(where, "message_id IS NOT NULL")
	}

	fields := []string{"id", "filename", "external_link", "type", "size", "creator_id", "created_timestamp", "updated_timestamp", "internal_path", "message_id"}
	if find.GetBlob {
		fields = append(fields, "blob")
	}

	query := fmt.Sprintf(`
		SELECT
			%s
		FROM resource
		WHERE %s
		GROUP BY id
		ORDER BY created_timestamp DESC
	`, strings.Join(fields, ", "), strings.Join(where, " AND "))
	if find.Limit != nil {
		query = fmt.Sprintf("%s LIMIT %d", query, *find.Limit)
		if find.Offset != nil {
			query = fmt.Sprintf("%s OFFSET %d", query, *find.Offset)
		}
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	list := make([]*store.Resource, 0)
	for rows.Next() {
		resource := store.Resource{}
		var messageID sql.NullInt32
		dests := []any{
			&resource.ID,
			&resource.Filename,
			&resource.ExternalLink,
			&resource.Type,
			&resource.Size,
			&resource.CreatorID,
			&resource.CreatedTimestamp,
			&resource.UpdatedTimestamp,
			&resource.InternalPath,
			&messageID,
		}
		if find.GetBlob {
			dests = append(dests, &resource.Blob)
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		if messageID.Valid {
			resource.MessageID = &messageID.Int32
		}
		list = append(list, &resource)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) UpdateResource(ctx context.Context, update *store.UpdateResource) (*store.Resource, error) {
	set, args := []string{}, []any{}

	if v := update.UpdatedTimestamp; v != nil {
		set, args = append(set, "updated_timestamp = ?"), append(args, *v)
	}
	if v := update.Filename; v != nil {
		set, args = append(set, "filename = ?"), append(args, *v)
	}
	if v := update.InternalPath; v != nil {
		set, args = append(set, "internal_path = ?"), append(args, *v)
	}
	if v := update.MessageID; v != nil {
		set, args = append(set, "message_id = ?"), append(args, *v)
	}
	if v := update.Blob; v != nil {
		set, args = append(set, "blob = ?"), append(args, v)
	}

	args = append(args, update.ID)
	fields := []string{"id", "filename", "external_link", "type", "size", "creator_id", "created_timestamp", "updated_timestamp", "internal_path"}
	stmt := `
		UPDATE resource
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING ` + strings.Join(fields, ", ")
	resource := store.Resource{}
	dests := []any{
		&resource.ID,
		&resource.Filename,
		&resource.ExternalLink,
		&resource.Type,
		&resource.Size,
		&resource.CreatorID,
		&resource.CreatedTimestamp,
		&resource.UpdatedTimestamp,
		&resource.InternalPath,
	}
	if err := d.db.QueryRowContext(ctx, stmt, args...).Scan(dests...); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (d *DB) DeleteResource(ctx context.Context, delete *store.DeleteResource) error {
	stmt := `DELETE FROM resource WHERE id = ?`
	result, err := d.db.ExecContext(ctx, stmt, delete.ID)
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

func vacuumResource(ctx context.Context, tx *sql.Tx) error {
	stmt := `
	DELETE FROM 
		resource 
	WHERE 
		creator_id NOT IN (SELECT id FROM user
	)`
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
