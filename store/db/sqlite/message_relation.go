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

func (d *DB) UpsertMessageRelation(ctx context.Context, create *store.MessageRelation) (*store.MessageRelation, error) {
	stmt := `
		INSERT INTO message_relation (
			message_id,
			related_message_id,
			type
		)
		VALUES (?, ?, ?)
		RETURNING message_id, related_message_id, type
	`
	messageRelation := &store.MessageRelation{}
	err := d.db.QueryRowContext(ctx, stmt, create.MessageID, create.RelatedMessageID, create.Type).Scan(&messageRelation.MessageID, &messageRelation.RelatedMessageID, &messageRelation.Type)
	if err != nil {
		return nil, err
	}

	return messageRelation, nil
}

func (d *DB) ListMessageRelations(ctx context.Context, find *store.FindMessageRelation) ([]*store.MessageRelation, error) {
	var args []any
	where := []string{"TRUE"}
	if find.MessageID != nil {
		where, args = append(where, "message_id = ?"), append(args, find.MessageID)
	}
	if find.RelatedMessageID != nil {
		where, args = append(where, "related_message_id = ?"), append(args, find.RelatedMessageID)
	}
	if find.Type != nil {
		where, args = append(where, "type = ?"), append(args, find.Type)
	}

	rows, err := d.db.QueryContext(ctx, `
		SELECT
			message_id,
			related_message_id,
			type
		FROM message_relation
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var list []*store.MessageRelation
	for rows.Next() {
		messageRelation := &store.MessageRelation{}
		err := rows.Scan(&messageRelation.MessageID, &messageRelation.RelatedMessageID, &messageRelation.Type)
		if err != nil {
			return nil, err
		}
		list = append(list, messageRelation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) DeleteMessageRelation(ctx context.Context, delete *store.DeleteMessageRelation) error {
	var args []any
	where := []string{"TRUE"}
	if delete.MessageID != nil {
		where, args = append(where, "message_id = ?"), append(args, delete.MessageID)
	}
	if delete.RelatedMessageID != nil {
		where, args = append(where, "related_message_id = ?"), append(args, delete.RelatedMessageID)
	}
	if delete.Type != nil {
		where, args = append(where, "type = ?"), append(args, delete.Type)
	}
	stmt := `
		DELETE FROM message_relation
		WHERE ` + strings.Join(where, " AND ")
	result, err := d.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func vacuumMessageRelations(ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.ExecContext(ctx, `
		DELETE FROM message_relation
		WHERE message_id NOT IN (SELECT id FROM message) OR related_message_id NOT IN (SELECT id FROM message)
	`); err != nil {
		return err
	}
	return nil
}
