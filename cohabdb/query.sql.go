// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package cohabdb

import (
	"context"
	"database/sql"
)

const expireSession = `-- name: ExpireSession :exec
UPDATE sessions
SET is_logged_in = false
WHERE id = ?
`

func (q *Queries) ExpireSession(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, expireSession, id)
	return err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, created_at, is_logged_in, google_force_approval, contact_groups_json, selected_resource_name FROM sessions
WHERE ID = ? LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id int64) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.IsLoggedIn,
		&i.GoogleForceApproval,
		&i.ContactGroupsJson,
		&i.SelectedResourceName,
	)
	return i, err
}

const getToken = `-- name: GetToken :one
SELECT token FROM users u
INNER JOIN sessions s
ON u.id = s.user_id
WHERE s.id = ? LIMIT 1
`

func (q *Queries) GetToken(ctx context.Context, id int64) (sql.NullString, error) {
	row := q.db.QueryRowContext(ctx, getToken, id)
	var token sql.NullString
	err := row.Scan(&token)
	return token, err
}

const getUser = `-- name: GetUser :one
SELECT id, sub, name, picture, token FROM users
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Sub,
		&i.Name,
		&i.Picture,
		&i.Token,
	)
	return i, err
}

const getUserBySession = `-- name: GetUserBySession :one
SELECT u.id, u.sub, u.name, u.picture, u.token FROM users u
INNER JOIN sessions s
WHERE u.id = s.user_id
AND s.id = ? LIMIT 1
`

func (q *Queries) GetUserBySession(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserBySession, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Sub,
		&i.Name,
		&i.Picture,
		&i.Token,
	)
	return i, err
}

const insertSession = `-- name: InsertSession :one
INSERT INTO sessions (
  id, user_id
) VALUES (
  ?, ?
)
RETURNING id, user_id, created_at, is_logged_in, google_force_approval, contact_groups_json, selected_resource_name
`

type InsertSessionParams struct {
	ID     int64
	UserID int64
}

func (q *Queries) InsertSession(ctx context.Context, arg InsertSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, insertSession, arg.ID, arg.UserID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.IsLoggedIn,
		&i.GoogleForceApproval,
		&i.ContactGroupsJson,
		&i.SelectedResourceName,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users
(
  sub,
  name,
  picture
) VALUES (
  ?, ?, ?
)
RETURNING id, sub, name, picture, token
`

type InsertUserParams struct {
	Sub     string
	Name    sql.NullString
	Picture sql.NullString
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.Sub, arg.Name, arg.Picture)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Sub,
		&i.Name,
		&i.Picture,
		&i.Token,
	)
	return i, err
}

const updateContactGroupsJSON = `-- name: UpdateContactGroupsJSON :exec
UPDATE sessions
SET contact_groups_json = ?
WHERE id = ?
`

type UpdateContactGroupsJSONParams struct {
	ContactGroupsJson sql.NullString
	ID                int64
}

func (q *Queries) UpdateContactGroupsJSON(ctx context.Context, arg UpdateContactGroupsJSONParams) error {
	_, err := q.db.ExecContext(ctx, updateContactGroupsJSON, arg.ContactGroupsJson, arg.ID)
	return err
}

const updateGoogleForceApproval = `-- name: UpdateGoogleForceApproval :exec
UPDATE sessions
SET google_force_approval = ?
WHERE id = ?
`

type UpdateGoogleForceApprovalParams struct {
	GoogleForceApproval bool
	ID                  int64
}

func (q *Queries) UpdateGoogleForceApproval(ctx context.Context, arg UpdateGoogleForceApprovalParams) error {
	_, err := q.db.ExecContext(ctx, updateGoogleForceApproval, arg.GoogleForceApproval, arg.ID)
	return err
}

const updateSelectedResourceName = `-- name: UpdateSelectedResourceName :exec
UPDATE sessions
SET selected_resource_name = ?
WHERE id = ?
`

type UpdateSelectedResourceNameParams struct {
	SelectedResourceName sql.NullString
	ID                   int64
}

func (q *Queries) UpdateSelectedResourceName(ctx context.Context, arg UpdateSelectedResourceNameParams) error {
	_, err := q.db.ExecContext(ctx, updateSelectedResourceName, arg.SelectedResourceName, arg.ID)
	return err
}

const updateTokenBySession = `-- name: UpdateTokenBySession :exec
UPDATE users
SET token = ?
WHERE (
  SELECT user_id
  FROM sessions
  WHERE sessions.id = ?
  AND users.id = user_id
)
`

type UpdateTokenBySessionParams struct {
	Token sql.NullString
	ID    int64
}

func (q *Queries) UpdateTokenBySession(ctx context.Context, arg UpdateTokenBySessionParams) error {
	_, err := q.db.ExecContext(ctx, updateTokenBySession, arg.Token, arg.ID)
	return err
}

const upsertSession = `-- name: UpsertSession :one
INSERT INTO sessions (
  id, user_id
) VALUES (
  ?, ?
)
ON CONFLICT(id, user_id) DO UPDATE SET
  id=excluded.id,
  created_at=strftime('%s','now'),
  is_logged_in=true
RETURNING id, user_id, created_at, is_logged_in, google_force_approval, contact_groups_json, selected_resource_name
`

type UpsertSessionParams struct {
	ID     int64
	UserID int64
}

func (q *Queries) UpsertSession(ctx context.Context, arg UpsertSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, upsertSession, arg.ID, arg.UserID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.IsLoggedIn,
		&i.GoogleForceApproval,
		&i.ContactGroupsJson,
		&i.SelectedResourceName,
	)
	return i, err
}

const upsertUser = `-- name: UpsertUser :one
INSERT INTO users
(
  sub,
  name,
  picture
) VALUES (
  ?, ?, ?
)
ON CONFLICT(sub) DO UPDATE SET
  name=excluded.name,
  picture=excluded.picture
RETURNING id, sub, name, picture, token
`

type UpsertUserParams struct {
	Sub     string
	Name    sql.NullString
	Picture sql.NullString
}

func (q *Queries) UpsertUser(ctx context.Context, arg UpsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, upsertUser, arg.Sub, arg.Name, arg.Picture)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Sub,
		&i.Name,
		&i.Picture,
		&i.Token,
	)
	return i, err
}
