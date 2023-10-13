-- drop all tables first
DROP TABLE IF EXISTS migration_history;
DROP TABLE IF EXISTS system_setting;
DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS user_setting;
DROP TABLE IF EXISTS message;
DROP TABLE IF EXISTS message_organizer;
DROP TABLE IF EXISTS message_relation;
DROP TABLE IF EXISTS resource;
DROP TABLE IF EXISTS tag;
DROP TABLE IF EXISTS activity;
DROP TABLE IF EXISTS storage;
DROP TABLE IF EXISTS idp;

-- migration_history
CREATE TABLE migration_history
(
	version    TEXT   NOT NULL PRIMARY KEY,
	created_timestamp BIGINT NOT NULL DEFAULT (strftime('%s', 'now'))
);

-- system_setting
CREATE TABLE system_setting
(
	name        TEXT NOT NULL,
	value       TEXT NOT NULL,
	description TEXT NOT NULL DEFAULT '',
	UNIQUE (name)
);

-- user
CREATE TABLE user
(
	id            INTEGER PRIMARY KEY AUTOINCREMENT,
	created_timestamp    BIGINT NOT NULL                                              DEFAULT (strftime('%s', 'now')),
	updated_timestamp    BIGINT NOT NULL                                              DEFAULT (strftime('%s', 'now')),
	row_status    TEXT   NOT NULL CHECK (row_status IN ('NORMAL', 'ARCHIVED')) DEFAULT 'NORMAL',
	username      TEXT   NOT NULL UNIQUE,
	role          TEXT   NOT NULL CHECK (role IN ('HOST', 'ADMIN', 'USER'))    DEFAULT 'USER',
	email         TEXT   NOT NULL                                              DEFAULT '',
	nickname      TEXT   NOT NULL                                              DEFAULT '',
	password_hash TEXT   NOT NULL,
	avatar_url    TEXT   NOT NULL                                              DEFAULT ''
);

CREATE INDEX idx_user_username ON user (username);

-- user_setting
CREATE TABLE user_setting
(
	user_id INTEGER NOT NULL,
	key     TEXT    NOT NULL,
	value   TEXT    NOT NULL,
	UNIQUE (user_id, key)
);

-- message
CREATE TABLE message
(
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	creator_id INTEGER NOT NULL,
	created_timestamp BIGINT  NOT NULL                                                          DEFAULT (strftime('%s', 'now')),
	updated_timestamp BIGINT  NOT NULL                                                          DEFAULT (strftime('%s', 'now')),
	row_status TEXT    NOT NULL CHECK (row_status IN ('NORMAL', 'ARCHIVED'))             DEFAULT 'NORMAL',
	content    TEXT    NOT NULL                                                          DEFAULT '',
	visibility TEXT    NOT NULL CHECK (visibility IN ('PUBLIC', 'PROTECTED', 'PRIVATE')) DEFAULT 'PRIVATE'
);

CREATE INDEX idx_message_creator_id ON message (creator_id);
CREATE INDEX idx_message_content ON message (content);
CREATE INDEX idx_message_visibility ON message (visibility);

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

-- message_organizer
CREATE TABLE message_organizer
(
	message_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	pinned  INTEGER NOT NULL CHECK (pinned IN (0, 1)) DEFAULT 0,
	UNIQUE (message_id, user_id)
);

-- message_relation
CREATE TABLE message_relation
(
	message_id         INTEGER NOT NULL,
	related_message_id INTEGER NOT NULL,
	type            TEXT    NOT NULL,
	UNIQUE (message_id, related_message_id, type)
);

-- resource
CREATE TABLE resource
(
	id            INTEGER PRIMARY KEY AUTOINCREMENT,
	creator_id    INTEGER NOT NULL,
	created_timestamp    BIGINT  NOT NULL DEFAULT (strftime('%s', 'now')),
	updated_timestamp    BIGINT  NOT NULL DEFAULT (strftime('%s', 'now')),
	filename      TEXT    NOT NULL DEFAULT '',
	blob          BLOB             DEFAULT NULL,
	external_link TEXT    NOT NULL DEFAULT '',
	type          TEXT    NOT NULL DEFAULT '',
	size          INTEGER NOT NULL DEFAULT 0,
	internal_path TEXT    NOT NULL DEFAULT '',
	message_id       INTEGER
);

CREATE INDEX idx_resource_creator_id ON resource (creator_id);

CREATE INDEX idx_resource_message_id ON resource (message_id);

-- tag
CREATE TABLE tag
(
	name       TEXT    NOT NULL,
	creator_id INTEGER NOT NULL,
	UNIQUE (name, creator_id)
);

-- activity
CREATE TABLE activity
(
	id         INTEGER PRIMARY KEY AUTOINCREMENT,
	creator_id INTEGER NOT NULL,
	created_timestamp BIGINT  NOT NULL                                            DEFAULT (strftime('%s', 'now')),
	type       TEXT    NOT NULL                                            DEFAULT '',
	level      TEXT    NOT NULL CHECK (level IN ('INFO', 'WARN', 'ERROR')) DEFAULT 'INFO',
	payload    TEXT    NOT NULL                                            DEFAULT '{}'
);

-- storage
CREATE TABLE storage
(
	id     INTEGER PRIMARY KEY AUTOINCREMENT,
	name   TEXT NOT NULL,
	type   TEXT NOT NULL,
	config TEXT NOT NULL DEFAULT '{}'
);

-- idp
CREATE TABLE idp
(
	id                INTEGER PRIMARY KEY AUTOINCREMENT,
	name              TEXT NOT NULL,
	type              TEXT NOT NULL,
	identifier_filter TEXT NOT NULL DEFAULT '',
	config            TEXT NOT NULL DEFAULT '{}'
);
