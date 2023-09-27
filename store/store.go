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
	"github.com/Nicknamezz00/mercury/server/profile"
	"github.com/pkg/errors"
	"modernc.org/sqlite"
	"sync"
)

type Store struct {
	Profile            *profile.Profile
	db                 *sql.DB
	systemSettingCache sync.Map
	userSettingCache   sync.Map
	userCache          sync.Map
}

func New(db *sql.DB, profile *profile.Profile) *Store {
	return &Store{
		Profile: profile,
		db:      db,
	}
}

func (s *Store) GetDB() *sql.DB {
	return s.db
}

type backup interface {
	NewBackup(string) (*sqlite.Backup, error)
}

func (s *Store) BackupTo(ctx context.Context, filename string) error {
	conn, err := s.db.Conn(ctx)
	if err != nil {
		return errors.Errorf("fail to get conn %s", err)
	}
	defer func(conn *sql.Conn) {
		_ = conn.Close()
	}(conn)

	err = conn.Raw(func(driverConn any) error {
		backupConn, ok := driverConn.(backup)
		if !ok {
			return errors.Errorf("db connection is not a sqlite backup")
		}
		backup, err := backupConn.NewBackup(filename)
		if err != nil {
			return errors.Errorf("fail to create sqlite backup %s", err)
		}
		nxt := true
		for nxt {
			nxt, err = backup.Step(-1)
			if err != nil {
				return errors.Errorf("fail to execute sqlite backup %s", err)
			}
		}
		return backup.Finish()
	})
	if err != nil {
		return errors.Errorf("fail to backup %s", err)
	}
	return nil
}

func (s *Store) Vacuum(ctx context.Context) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)
	if err := s.vacuum(ctx, tx); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if _, err := s.db.Exec("VACUUM"); err != nil {
		return err
	}
	return nil
}

func (s *Store) vacuum(ctx context.Context, tx *sql.Tx) error {
	// TODO: implement each table
	if err := vacuumMessage(ctx, tx); err != nil {
		return err
	}
	if err := vacuumUserSetting(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMessageOrganizer(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMessageResource(ctx, tx); err != nil {
		return err
	}
	if err := vacuumMessageRelation(ctx, tx); err != nil {
		return err
	}
	return nil
}
