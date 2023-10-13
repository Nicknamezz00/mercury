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

import "context"

type Tag struct {
	Name      string
	CreatorID int32
}

type FindTag struct {
	CreatorID int32
}

type DeleteTag struct {
	Name      string
	CreatorID int32
}

func (s *Store) UpsertTag(ctx context.Context, upsert *Tag) (*Tag, error) {
	return s.driver.UpsertTag(ctx, upsert)
}

func (s *Store) ListTags(ctx context.Context, find *FindTag) ([]*Tag, error) {
	return s.driver.ListTags(ctx, find)
}

func (s *Store) DeleteTag(ctx context.Context, delete *DeleteTag) error {
	return s.driver.DeleteTag(ctx, delete)
}
