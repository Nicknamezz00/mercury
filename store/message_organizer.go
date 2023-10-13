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
	"github.com/pkg/errors"
)

type MessageOrganizer struct {
	MessageID int32
	UserID    int32
	Pinned    bool
}

type FindMessageOrganizer struct {
	MessageID int32
	UserID    int32
}

type DeleteMessageOrganizer struct {
	MessageID *int32
	UserID    *int32
}

func (s *Store) UpsertMessageOrganizer(ctx context.Context, upsert *MessageOrganizer) (*MessageOrganizer, error) {
	return s.driver.UpsertMessageOrganizer(ctx, upsert)
}

func (s *Store) GetMessageOrganizer(ctx context.Context, find *FindMessageOrganizer) (*MessageOrganizer, error) {
	list, err := s.ListMessageOrganizer(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, errors.New("not found")
	}
	return list[0], nil
}

func (s *Store) ListMessageOrganizer(ctx context.Context, find *FindMessageOrganizer) ([]*MessageOrganizer, error) {
	return s.driver.ListMessageOrganizer(ctx, find)
}

func (s *Store) DeleteMessageOrganizer(ctx context.Context, delete *DeleteMessageOrganizer) error {
	return s.driver.DeleteMessageOrganizer(ctx, delete)
}
