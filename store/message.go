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
	"fmt"
	"github.com/Nicknamezz00/mercury/common/util"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type Visibility string

const (
	Public    Visibility = "PUBLIC"
	Protected Visibility = "PROTECTED"
	PRIVATE   Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	switch v {
	case Public:
		return "PUBLIC"
	case Protected:
		return "PROTECTED"
	case PRIVATE:
		return "PRIVATE"
	}
	return "PRIVATE"
}

type Message struct {
	ID               int32
	RowStatus        RowStatus
	CreatorID        int32
	CreatedTimestamp int64
	UpdatedTimestamp int64

	Content        string
	Visibility     Visibility
	Pinned         bool
	ResourceIDList []int32
	RelationList   []*MessageRelation
}

type FindMessage struct {
	ID                     *int32
	RowStatus              *RowStatus
	CreatorID              *int32
	CreatedTimestampAfter  *int64
	CreatedTimestampBefore *int64

	Pinned         *bool
	ContentSearch  []string
	VisibilityList []Visibility

	Limit                   *int
	Offset                  *int
	OrderByUpdatedTimestamp bool
}

type UpdateMessage struct {
	ID               int32
	RowStatus        *RowStatus
	CreatedTimestamp *int64
	UpdatedTimestamp *int64
	Content          *string
	Visibility       *Visibility
}

type DeleteMessage struct {
	ID int32
}

func (s *Store) CreateMessage(ctx context.Context, m *Message) (*Message, error) {
	if m.CreatedTimestamp == 0 {
		m.CreatedTimestamp = time.Now().Unix()
	}
	stmt := `
		INSERT INTO message (
			creator_id,
			created_timestamp,
		  content,
		  visibility
		)
		VALUES (?, ?, ?, ?)
		RETURNING id, created_timestamp, updated_timestamp, row_status
	`
	err := s.db.QueryRowContext(ctx, stmt, m.ID, m.CreatedTimestamp, m.UpdatedTimestamp, m.RowStatus).Scan(&m.ID, &m.CreatedTimestamp, &m.UpdatedTimestamp, &m.RowStatus)
	if err != nil {
		return nil, err
	}
	res := m
	return res, nil
}

func (s *Store) ListMessages(ctx context.Context, find *FindMessage) ([]*Message, error) {
	where, args := []string{"1 = 1"}, []any{}
	if v := find.ID; v != nil {
		where, args = append(where, "message.id = ?"), append(args, *v)
	}
	if v := find.CreatorID; v != nil {
		where, args = append(where, "message.creator_id = ?"), append(args, *v)
	}
	if v := find.RowStatus; v != nil {
		where, args = append(where, "message.row_status = ?"), append(args, *v)
	}
	if v := find.CreatedTimestampBefore; v != nil {
		where, args = append(where, "message.created_timestamp < ?"), append(args, *v)
	}
	if v := find.CreatedTimestampAfter; v != nil {
		where, args = append(where, "message.created_timestamp > ?"), append(args, *v)
	}
	if v := find.Pinned; v != nil {
		where, args = append(where, "message_organizer.pinned = 1"), append(args, *v)
	}
	if v := find.ContentSearch; len(v) != 0 {
		for _, s := range v {
			where, args = append(where, "message.content LIKE ?"), append(args, "%"+s+"%")
		}
	}
	if v := find.VisibilityList; len(v) != 0 {
		var lst []string
		for _, vs := range v {
			lst = append(lst, fmt.Sprintf("$%d", len(args)+1))
			args = append(args, vs)
		}
		where = append(where, fmt.Sprintf("message.visiblitiy in (%s)", strings.Join(lst, ",")))
	}
	orders := []string{"pinned DESC"}
	if find.OrderByUpdatedTimestamp {
		orders = append(orders, "updated_timestamp DESC")
	} else {
		orders = append(orders, "created_timestamp DESC")
	}
	orders = append(orders, "id DESC")
	query := formatQueryStatement(where, orders)
	if find.Limit != nil {
		query = fmt.Sprintf("%s LIMIT %d", query, *find.Limit)
		if find.Offset != nil {
			query = fmt.Sprintf("%s OFFSET %d", query, *find.Offset)
		}
	}
	rows, err := s.db.QueryContext(ctx, query, args)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	messages := make([]*Message, 0)
	for rows.Next() {
		var (
			m               Message
			mResourceIDList sql.NullString
			mRelationList   sql.NullString
		)
		err = rows.Scan(&m.ID, &m.CreatorID, &m.CreatedTimestamp, &m.UpdatedTimestamp, &m.RowStatus, &m.Content, &m.Visibility, &m.Pinned, &mResourceIDList, &mRelationList)
		if err != nil {
			return nil, err
		}
		if mResourceIDList.Valid {
			idStringList := strings.Split(mResourceIDList.String, ",")
			m.ResourceIDList = make([]int32, 0, len(idStringList))
			for _, idString := range idStringList {
				id, err := util.ConvertStringToInt32(idString)
				if err != nil {
					return nil, err
				}
				m.ResourceIDList = append(m.ResourceIDList, id)
			}
		}
		if mRelationList.Valid {
			m.RelationList = make([]*MessageRelation, 0)
			relatedTypeList := strings.Split(mRelationList.String, ",")
			for _, rt := range relatedTypeList {
				relationAndType := strings.Split(rt, ":")
				if len(relationAndType) != 2 {
					return nil, errors.Errorf("invalid relation format")
				}
				relatedID, err := util.ConvertStringToInt32(relationAndType[0])
				if err != nil {
					return nil, err
				}
				m.RelationList = append(m.RelationList, &MessageRelation{
					MessageID:      m.ID,
					RelatedMessage: relatedID,
					Type:           MessageRelationType(relationAndType[1]),
				})
			}
		}
		messages = append(messages, &m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func formatQueryStatement(where []string, orders []string) string {
	query := `
	SELECT 
    message.id AS id,
		message.creator_id AS creator_id,
		message.created_timestamp AS created_timestamp,
		message.updated_timestamp AS updated_timestamp,
		message.row_status AS row_status,
		message.content AS content,
		message.visibility AS visibility,
		CASE WHEN message_organizer.pinned = 1 THEN 1 ELSE 0 END AS pinned,
		GROUP_CONCAT(message_resource.resource_id) AS resource_id_list,
		(
			SELECT 
			    GROUP_CONCAT(related_message_id || ':' || type)
			FROM 
			    message_relation
			WHERE
			    message_relation.message_id = message.id
			GROUP BY
			    message_relation.message_id
		) AS relation_list
	FROM 
	    message
	LEFT JOIN 
	    message_organizer ON message.id = message_organizer.message_id
	LEFT JOIN 
	    message_resource ON message.id = message_resource.message_id
	WHERE ` + strings.Join(where, " AND ") + ` 
	GROUP BY message.id
	ORDER BY ` + strings.Join(orders, ", ") + ` 
	`
	return query
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
	set, args := []string{}, []any{}
	if v := update.CreatedTimestamp; v != nil {
		set, args = append(set, "created_timestamp = ?"), append(args, *v)
	}
	if v := update.UpdatedTimestamp; v != nil {
		set, args = append(set, "updated_timestamp = ?"), append(args, *v)
	}
	if v := update.RowStatus; v != nil {
		set, args = append(set, "row_status = ?"), append(args, *v)
	}
	if v := update.Content; v != nil {
		set, args = append(set, "content = ?"), append(args, *v)
	}
	if v := update.Visibility; v != nil {
		set, args = append(set, "visibility = ?"), append(args, *v)
	}
	args = append(args, update.ID)
	stmt := `
	UPDATE message
	SET ` + strings.Join(set, ", ") + `
	WHERE id = ?
	`
	if _, err := s.db.ExecContext(ctx, stmt, args...); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteMessage(ctx context.Context, delete *DeleteMessage) error {
	where, args := []string{"id = ?"}, []any{delete.ID}
	stmt := `DELETE FROM message WHERE ` + strings.Join(where, " AND ")
	result, err := s.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		return err
	}
	if err := s.Vacuum(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Store) FindMessageVisibilityList(ctx context.Context, messageIDs []int32) ([]Visibility, error) {
	args := make([]any, 0, len(messageIDs))
	list := make([]string, 0, len(messageIDs))
	for _, id := range messageIDs {
		args = append(args, id)
		list = append(list, "?")
	}
	where := fmt.Sprintf("id in (%s)", strings.Join(list, ","))
	query := `SELECT DISTINCT(visibility) FROM message WHERE ` + where
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)
	visList := make([]Visibility, 0)
	for rows.Next() {
		var v Visibility
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		visList = append(visList, v)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return visList, nil
}

func vacuumMessage(ctx context.Context, tx *sql.Tx) error {
	stmt := `
	DELETE FROM
    message
	WHERE
		creator_id NOT IN (
			SELECT 
	      id
			FROM
				user
		)`
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
