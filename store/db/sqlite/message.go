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
	"fmt"
	"github.com/Nicknamezz00/mercury/common/util"
	"github.com/Nicknamezz00/mercury/internal/types"
	"github.com/Nicknamezz00/mercury/store"
	"github.com/pkg/errors"
	"strings"
)

func (d *DB) CreateMessage(ctx context.Context, create *store.Message) (*store.Message, error) {
	fields := []string{"`creator_id`", "`content`", "`visibility`"}
	placeholder := []string{"?", "?", "?"}
	args := []any{create.CreatorID, create.Content, create.Visibility}

	if create.ID != 0 {
		fields = append(fields, "`id`")
		placeholder = append(placeholder, "?")
		args = append(args, create.ID)
	}
	if create.CreatedTimestamp != 0 {
		fields = append(fields, "`created_timestamp`")
		placeholder = append(placeholder, "?")
		args = append(args, create.CreatedTimestamp)
	}
	if create.UpdatedTimestamp != 0 {
		fields = append(fields, "`updated_timestamp`")
		placeholder = append(placeholder, "?")
		args = append(args, create.UpdatedTimestamp)
	}
	if create.RowStatus != "" {
		fields = append(fields, "`row_status`")
		placeholder = append(placeholder, "?")
		args = append(args, create.RowStatus)
	}

	stmt := `
		INSERT INTO message (` + strings.Join(fields, ", ") + `) 
		VALUES (` + strings.Join(placeholder, ", ") + `) 
		RETURNING id, created_timestamp, updated_timestamp, row_status
	`
	err := d.db.QueryRowContext(ctx, stmt, args...).Scan(&create.ID, &create.CreatedTimestamp, &create.UpdatedTimestamp, &create.RowStatus)
	if err != nil {
		return nil, err
	}
	return create, nil
}

func (d *DB) ListMessages(ctx context.Context, find *store.FindMessage) ([]*store.Message, error) {
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
	if v := find.ContentSearch; len(v) != 0 {
		for _, s := range v {
			where, args = append(where, "message.content LIKE ?"), append(args, "%"+s+"%")
		}
	}
	if v := find.VisibilityList; len(v) != 0 {
		var list []string
		for _, visibility := range v {
			list = append(list, fmt.Sprintf("$%d", len(args)+1))
			args = append(args, visibility)
		}
		where = append(where, fmt.Sprintf("message.visibility in (%s)", strings.Join(list, ",")))
	}
	if v := find.Pinned; v != nil {
		where = append(where, "message_organizer.pinned = 1")
	}
	if v := find.HasParent; v != nil {
		if *v {
			where = append(where, "parent_id IS NOT NULL")
		} else {
			where = append(where, "parent_id IS NULL")
		}
	}

	orders := []string{"pinned DESC"}
	if find.OrderByUpdatedTimestamp {
		orders = append(orders, "updated_timestamp DESC")
	} else {
		orders = append(orders, "created_timestamp DESC")
	}
	orders = append(orders, "id DESC")

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
		(
			SELECT
				related_message_id
			FROM
				message_relation
			WHERE
				message_relation.message_id = message.id AND message_relation.type = 'COMMENT'
			LIMIT 1
		) AS parent_id,
		GROUP_CONCAT(resource.id) AS resource_id_list,
		(
			SELECT
				GROUP_CONCAT(message_id || ':' || related_message_id || ':' || type)
			FROM
				message_relation
			WHERE
				message_relation.message_id = message.id OR message_relation.related_message_id = message.id
		) AS relation_list
	FROM
		message
	LEFT JOIN
		message_organizer ON message.id = message_organizer.message_id
	LEFT JOIN
		resource ON message.id = resource.message_id
	WHERE ` + strings.Join(where, " AND ") + `
	GROUP BY message.id
	ORDER BY ` + strings.Join(orders, ", ") + `
	`
	if find.Limit != nil {
		query = fmt.Sprintf("%s LIMIT %d", query, *find.Limit)
		if find.Offset != nil {
			query = fmt.Sprintf("%s OFFSET %d", query, *find.Offset)
		}
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	list := make([]*store.Message, 0)
	for rows.Next() {
		var message store.Message
		var messageResourceIDList sql.NullString
		var messageRelationList sql.NullString
		if err := rows.Scan(
			&message.ID,
			&message.CreatorID,
			&message.CreatedTimestamp,
			&message.UpdatedTimestamp,
			&message.RowStatus,
			&message.Content,
			&message.Visibility,
			&message.Pinned,
			&message.ParentID,
			&messageResourceIDList,
			&messageRelationList,
		); err != nil {
			return nil, err
		}

		if messageResourceIDList.Valid {
			idStringList := strings.Split(messageResourceIDList.String, ",")
			message.ResourceIDList = make([]int32, 0, len(idStringList))
			for _, idString := range idStringList {
				id, err := util.ConvertStringToInt32(idString)
				if err != nil {
					return nil, err
				}
				message.ResourceIDList = append(message.ResourceIDList, id)
			}
		}
		if messageRelationList.Valid {
			message.RelationList = make([]*store.MessageRelation, 0)
			relatedMessageTypeList := strings.Split(messageRelationList.String, ",")
			for _, relatedMessageType := range relatedMessageTypeList {
				relatedMessageTypeList := strings.Split(relatedMessageType, ":")
				if len(relatedMessageTypeList) != 3 {
					return nil, errors.New("invalid relation format")
				}
				messageID, err := util.ConvertStringToInt32(relatedMessageTypeList[0])
				if err != nil {
					return nil, err
				}
				relatedMessageID, err := util.ConvertStringToInt32(relatedMessageTypeList[1])
				if err != nil {
					return nil, err
				}
				relationType := types.MessageRelationType(relatedMessageTypeList[2])
				message.RelationList = append(message.RelationList, &store.MessageRelation{
					MessageID:        messageID,
					RelatedMessageID: relatedMessageID,
					Type:             relationType,
				})
			}
		}
		list = append(list, &message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) UpdateMessage(ctx context.Context, update *store.UpdateMessage) error {
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
	if _, err := d.db.ExecContext(ctx, stmt, args...); err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteMessage(ctx context.Context, delete *store.DeleteMessage) error {
	where, args := []string{"id = ?"}, []any{delete.ID}
	stmt := `DELETE FROM message WHERE ` + strings.Join(where, " AND ")
	result, err := d.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		return err
	}

	if err := d.Vacuum(ctx); err != nil {
		// Prevent linter warning.
		return err
	}
	return nil
}

func (d *DB) FindMessagesVisibilityList(ctx context.Context, messageIDs []int32) ([]types.Visibility, error) {
	args := make([]any, 0, len(messageIDs))
	list := make([]string, 0, len(messageIDs))
	for _, messageID := range messageIDs {
		args = append(args, messageID)
		list = append(list, "?")
	}

	where := fmt.Sprintf("id in (%s)", strings.Join(list, ","))
	query := `SELECT DISTINCT(visibility) FROM message WHERE ` + where
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	visibilityList := make([]types.Visibility, 0)
	for rows.Next() {
		var visibility types.Visibility
		if err := rows.Scan(&visibility); err != nil {
			return nil, err
		}
		visibilityList = append(visibilityList, visibility)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return visibilityList, nil
}

func vacuumMessage(ctx context.Context, tx *sql.Tx) error {
	stmt := `
	DELETE FROM
    message
	WHERE
		creator_id NOT IN (SELECT id FROM user)
	`
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
