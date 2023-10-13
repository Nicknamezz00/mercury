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
	"github.com/Nicknamezz00/mercury/internal/constants"
	"github.com/Nicknamezz00/mercury/server/profile"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/pkg/errors"
	"modernc.org/sqlite"
)

type DB struct {
	db      *sql.DB
	profile *profile.Profile
}

// NewDB opens a specific db by DSN.
func NewDB(profile *profile.Profile) (store.Driver, error) {
	if profile.DSN == "" {
		return nil, errors.New("dsn required")
	}

	sqliteDB, err := sql.Open(constants.SQLITE, profile.DSN+"?_pragma=foreign_keys(0)&_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)")
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open db with dsn: %s", profile.DSN)
	}
	d := DB{
		db:      sqliteDB,
		profile: profile,
	}
	return &d, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) BackupTo(ctx context.Context, filename string) error {
	conn, err := d.db.Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "fail to open new connection")
	}
	defer func(conn *sql.Conn) {
		_ = conn.Close()
	}(conn)

	err = conn.Raw(func(driverConn any) error {
		type backuper interface {
			NewBackup(string) (*sqlite.Backup, error)
		}
		backupConn, ok := driverConn.(backuper)
		if !ok {
			return errors.New("db connection is not a sqlite backuper")
		}

		bck, err := backupConn.NewBackup(filename)
		if err != nil {
			return errors.Wrap(err, "fail to create sqlite backup")
		}

		for more := true; more; {
			more, err = bck.Step(-1)
			if err != nil {
				return errors.Wrap(err, "fail to execute sqlite backup")
			}
		}
		return bck.Finish()
	})

	if err != nil {
		return errors.Wrap(err, "fail to backup")
	}
	return nil
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}

func (d *DB) Vacuum(ctx context.Context) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)
	if err := d.vacuum(ctx, tx); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if _, err := d.db.Exec("VACUUM"); err != nil {
		return err
	}
	return nil
}

func (d *DB) vacuum(ctx context.Context, tx *sql.Tx) error {
	if err := vacuumMessage(ctx, tx); err != nil {
		return err
	}
	if err := vacuumUserSetting(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMessageOrganizer(ctx, tx); err != nil {
		return err
	}
	if err := vacuumResource(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMessageRelations(ctx, tx); err != nil {
		return err
	}
	if err := vacuumTag(ctx, tx); err != nil {
		return err
	}
	return nil
}
