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
	"github.com/Nicknamezz00/mercury/internal/types"
)

type MessageRelation struct {
	MessageID        int32
	RelatedMessageID int32
	Type             types.MessageRelationType
}

type FindMessageRelation struct {
	MessageID        *int32
	RelatedMessageID *int32
	Type             *types.MessageRelationType
}

type DeleteMessageRelation struct {
	MessageID        *int32
	RelatedMessageID *int32
	Type             *types.MessageRelationType
}

func (s *Store) UpsertMessageRelation(ctx context.Context, create *MessageRelation) (*MessageRelation, error) {
	return s.driver.UpsertMessageRelation(ctx, create)
}

func (s *Store) ListMessageRelations(ctx context.Context, find *FindMessageRelation) ([]*MessageRelation, error) {
	return s.driver.ListMessageRelations(ctx, find)
}

func (s *Store) GetMessageRelation(ctx context.Context, find *FindMessageRelation) (*MessageRelation, error) {
	list, err := s.ListMessageRelations(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return list[0], nil
}

func (s *Store) DeleteMessageRelation(ctx context.Context, delete *DeleteMessageRelation) error {
	return s.driver.DeleteMessageRelation(ctx, delete)
}
