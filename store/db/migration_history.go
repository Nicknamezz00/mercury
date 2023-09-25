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

package db

import (
	"context"
	"database/sql"
	"strings"
)

type MigrationHistory struct {
	Version          string
	CreatedTimestamp int64
}

type FindMigrationHistory struct {
	Version *string
}

type UpsertMigrationHistory struct {
	Version string
}

func (d *DB) FindMigrationHistoryList(ctx context.Context, find *FindMigrationHistory) ([]*MigrationHistory, error) {
	where, args := []string{"1 = 1"}, []any{}
	if v := find.Version; v != nil {
		where, args = append(where, "version = ?"), append(args, *v)
	}
	query := `
		SELECT
			version, created_timestamp
		FROM
			migration_history
		WHERE` + strings.Join(where, " AND ") + `
		ORDER BY created_timestamp DESC
	`
	rows, err := d.DBInstance.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	list := make([]*MigrationHistory, 0)
	for rows.Next() {
		var m MigrationHistory
		if err := rows.Scan(&m.Version, m.CreatedTimestamp); err != nil {
			return nil, err
		}
		list = append(list, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) UpsertMigrationHistory(ctx context.Context, upsert *UpsertMigrationHistory) (*MigrationHistory, error) {
	stmt := `
		INSERT INTO migration_history (version)
		VALUES (?)
		ON CONFLICT (version) DO UPDATE 
		SET 
		  version=EXCLUDED.version
		RETURNING version, created_timestamp
  `
	var m MigrationHistory
	if err := d.DBInstance.QueryRowContext(ctx, stmt, upsert.Version).Scan(&m.Version, &m.CreatedTimestamp); err != nil {
		return nil, err
	}
	return &m, nil
}
