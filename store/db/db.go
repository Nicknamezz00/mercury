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
	"embed"
	"fmt"
	"github.com/Nicknamezz00/mercury/server/profile"
	"github.com/Nicknamezz00/mercury/server/version"
	"github.com/pkg/errors"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"time"
)

type DB struct {
	DBInstance *sql.DB
	profile    *profile.Profile
}

//go:embed migration
var migrationFS embed.FS

//go:embed seed
var seedFS embed.FS

var (
	minorDirRegexp       = regexp.MustCompile(`^migration/prod/[0-9]+\.[0-9]+$`)
	latestSchemaFileName = "LATEST__SCHEMA.sql"
)

func New(profile *profile.Profile) *DB {
	return &DB{
		profile: profile,
	}
}

func (d *DB) Open() error {
	if d.profile.DSN == "" {
		return errors.New("empty dsn")
	}
	sqlite, err := sql.Open("sqlite", d.profile.DSN+"?_pragma=foreign_keys(0)&_pragma=busy_timeout(10000)&_pragma=journal_mode(WAL)")
	if err != nil {
		return errors.Wrapf(err, "failed to open db with dsn: %s", d.profile.DSN)
	}
	d.DBInstance = sqlite
	return nil
}

func (d *DB) Migrate(ctx context.Context) error {
	// production
	if d.profile.Mode == "production" {
		_, err := os.Stat(d.profile.DSN)
		if err != nil {
			// db not exist, create a new one with the latest schema.
			return d.newDBWithLatestSchema(ctx, err)
		} else {
			// db exists, check whether it is necessary to migrate.
			currentVersion := version.GetCurrentVersion(d.profile.Mode)
			migrationHistory, err := d.FindMigrationHistoryList(ctx, &FindMigrationHistory{})
			if err != nil {
				return errors.Wrap(err, "fail to find migration history")
			}
			if len(migrationHistory) == 0 {
				_, err := d.UpsertMigrationHistory(ctx, &UpsertMigrationHistory{
					Version: currentVersion,
				})
				if err != nil {
					return errors.Wrap(err, "failed to upsert migration history")
				}
				return nil
			}
			mhVersionList := []string{}
			for _, mh := range migrationHistory {
				mhVersionList = append(mhVersionList, mh.Version)
			}
			sort.Sort(version.SortVersion(mhVersionList))
			latestMHVersion := mhVersionList[len(mhVersionList)-1]
			if version.IsVersionGreaterThan(version.GetSchemaVersion(currentVersion), latestMHVersion) {
				return d.migrateDB(ctx, latestMHVersion, currentVersion)
			}
		}
	} else {
		// non-production mode, always migrate.
		return d.migrateNonProductionDB(ctx)
	}
	return nil
}

func (d *DB) migrateDB(ctx context.Context, latestMHVerion string, currentVersion string) error {
	minorVersion := getMinorVersionList()
	// backup
	bytes, err := os.ReadFile(d.profile.DSN)
	if err != nil {
		return errors.Wrap(err, "failed to read raw db file")
	}
	backupPath := fmt.Sprintf("%s/mercury_%s_%d_backup.db", d.profile.Data, d.profile.Version, time.Now().Unix())
	if err := os.WriteFile(backupPath, bytes, 0644); err != nil {
		return errors.Wrap(err, "failed to write backup db file")
	}
	println("successfully copy a backup database")
	println("start migration...")
	for _, mv := range minorVersion {
		v := mv + ".0"
		if version.IsVersionGreaterThan(v, latestMHVerion) && version.IsVersionGreaterOrEqualThan(currentVersion, v) {
			println("applying migration for", v)
			if err := d.applyMigrationForMinorVersion(ctx, mv); err != nil {
				return errors.Wrap(err, "failed to apply migration for minor version")
			}
		}
	}
	println("end migration")
	if err := os.Remove(backupPath); err != nil {
		println(fmt.Sprintf("failed to remove temporary backup database file, error: %v", err))
	}
	return nil
}

func (d *DB) newDBWithLatestSchema(ctx context.Context, err error) error {
	if errors.Is(err, os.ErrNotExist) {
		if err := d.applyLatestSchema(ctx); err != nil {
			return errors.Wrap(err, "failed to apply latest schema")
		}
	} else {
		return errors.Wrap(err, "failed to get db file stat")
	}
	return nil
}

func (d *DB) migrateNonProductionDB(ctx context.Context) error {
	if _, err := os.Stat(d.profile.DSN); errors.Is(err, os.ErrNotExist) {
		if err := d.applyLatestSchema(ctx); err != nil {
			return errors.Wrap(err, "failed to apply latest schema")
		}
		if d.profile.Mode == "demo" {
			if err := d.seed(ctx); err != nil {
				return errors.Wrap(err, "failed to seed database")
			}
		}
	}
	return nil
}

func getMinorVersionList() []string {
	mv := []string{}
	if err := fs.WalkDir(migrationFS, "migration", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && minorDirRegexp.MatchString(path) {
			mv = append(mv, d.Name())
		}
		return nil
	}); err != nil {
		panic(err)
	}
	sort.Sort(version.SortVersion(mv))
	return mv
}

func (d *DB) applyLatestSchema(ctx context.Context) error {
	schemaMode := "dev"
	if d.profile.Mode == "production" {
		schemaMode = "production"
	}
	latestFile := fmt.Sprintf("%s/%s/%s", "migration", schemaMode, latestSchemaFileName)
	buf, err := migrationFS.ReadFile(latestFile)
	if err != nil {
		return errors.Wrapf(err, "failed to read latest schema file %q", latestFile)
	}
	stmt := string(buf)
	if err := d.execute(ctx, stmt); err != nil {
		return errors.Wrapf(err, "migration error: %s", stmt)
	}
	return nil
}

func (d *DB) applyMigrationForMinorVersion(ctx context.Context, minorVersion string) error {
	filenames, err := fs.Glob(migrationFS, fmt.Sprintf("%s/%s/*.sql", "migration/production", minorVersion))
	if err != nil {
		return errors.Wrap(err, "failed to read migration files")
	}
	sort.Strings(filenames)
	stmt := ""
	for _, filename := range filenames {
		buf, err := migrationFS.ReadFile(filename)
		if err != nil {
			return errors.Wrapf(err, "failed to read migration file %s", filename)
		}
		stmt += string(buf)
		if err := d.execute(ctx, stmt); err != nil {
			return errors.Wrapf(err, "execute statement error: %s", stmt)
		}
	}
	v := minorVersion + ".0"
	if _, err := d.UpsertMigrationHistory(ctx, &UpsertMigrationHistory{
		Version: v,
	}); err != nil {
		return errors.Wrapf(err, "failed to upsert migration history with version: %s", v)
	}
	return nil
}

func (d *DB) seed(ctx context.Context) error {
	filenames, err := fs.Glob(seedFS, fmt.Sprintf("%s/*.sql", "seed"))
	if err != nil {
		return errors.Wrap(err, "failed to read seed files")
	}
	sort.Strings(filenames)
	for _, filename := range filenames {
		buf, err := seedFS.ReadFile(filename)
		if err != nil {
			return errors.Wrapf(err, "failed to read seed file %s", filename)
		}
		stmt := string(buf)
		if err := d.execute(ctx, stmt); err != nil {
			return errors.Wrapf(err, "execute statement error: %s", stmt)
		}
	}
	return nil
}

func (d *DB) execute(ctx context.Context, stmt string) error {
	tx, err := d.DBInstance.Begin()
	if err != nil {
		return err
	}
	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	if _, err := tx.ExecContext(ctx, stmt); err != nil {
		return errors.Wrap(err, "failed to execute statement")
	}
	return tx.Commit()
}
