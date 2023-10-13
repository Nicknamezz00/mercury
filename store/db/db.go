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
	"database/sql"
	"github.com/Nicknamezz00/mercury/internal/constants"
	"github.com/Nicknamezz00/mercury/server/profile"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/Nicknamezz00/mercury/store/db/sqlite"
	"github.com/pkg/errors"
)

type DB struct {
	DBInstance *sql.DB
	profile    *profile.Profile
}

func NewDBDriver(profile *profile.Profile) (store.Driver, error) {
	var (
		d   store.Driver
		err error
	)
	if profile.Mode == constants.PRODUCTION && profile.Driver != constants.SQLITE {
		return nil, errors.New("only sqlite is supported in production mode")
	}

	switch profile.Driver {
	case constants.SQLITE:
		d, err = sqlite.NewDB(profile)
	case constants.MYSQL:
		// TODO: Mysql driver
		//d, err = mysql.NewDB(profile)
	default:
		return nil, errors.New("unsupported db")
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db driver")
	}
	return d, nil
}
