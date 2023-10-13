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

type Message struct {
	ID               int32
	RowStatus        types.RowStatus
	CreatorID        int32
	CreatedTimestamp int64
	UpdatedTimestamp int64

	Content    string
	Visibility types.Visibility

	ParentID       *int32
	Pinned         bool
	ResourceIDList []int32
	RelationList   []*MessageRelation
}

type FindMessage struct {
	ID                     *int32
	RowStatus              *types.RowStatus
	CreatorID              *int32
	CreatedTimestampAfter  *int64
	CreatedTimestampBefore *int64

	ContentSearch  []string
	VisibilityList []types.Visibility
	Pinned         *bool
	HasParent      *bool

	Limit                   *int
	Offset                  *int
	OrderByUpdatedTimestamp bool
}

type UpdateMessage struct {
	ID               int32
	RowStatus        *types.RowStatus
	CreatedTimestamp *int64
	UpdatedTimestamp *int64
	Content          *string
	Visibility       *types.Visibility
}

type DeleteMessage struct {
	ID int32
}

func (s *Store) CreateMessage(ctx context.Context, create *Message) (*Message, error) {
	return s.driver.CreateMessage(ctx, create)
}

func (s *Store) ListMessages(ctx context.Context, find *FindMessage) ([]*Message, error) {
	return s.driver.ListMessages(ctx, find)
}

func (s *Store) GetMessage(ctx context.Context, find *FindMessage) (*Message, error) {
	list, err := s.ListMessages(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}
	return list[0], nil
}

func (s *Store) UpdateMessage(ctx context.Context, update *UpdateMessage) error {
	return s.driver.UpdateMessage(ctx, update)
}

func (s *Store) DeleteMessage(ctx context.Context, delete *DeleteMessage) error {
	return s.driver.DeleteMessage(ctx, delete)
}

func (s *Store) FindMessageVisibilityList(ctx context.Context, messageIDs []int32) ([]types.Visibility, error) {
	return s.driver.FindMessagesVisibilityList(ctx, messageIDs)
}
