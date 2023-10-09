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

package teststore

import (
	"context"
	"fmt"
	"github.com/Nicknamezz00/mercury/store"
	db2 "github.com/Nicknamezz00/mercury/store/db"
	"github.com/Nicknamezz00/mercury/test"
	"testing"
)

func NewTestingStore(ctx context.Context, t *testing.T) *store.Store {
	profile := test.NewTestingProfile(t)
	db := db2.New(profile)
	if err := db.Open(); err != nil {
		fmt.Printf("failed to open db, error: %+v\n", err)
	}
	if err := db.Migrate(ctx); err != nil {
		fmt.Printf("failed to migrate db, error: %+v\n", err)
	}
	return store.New(db.DBInstance, profile)
}