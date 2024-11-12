// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: watch-list.sql

package repository

import (
	"context"

	"github.com/google/uuid"
)

const addToList = `-- name: AddToList :one
INSERT INTO watch_list (
  id, user_id, type, media_id, poster, title, status, episodes, duration  
) VALUES ( uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8 )
RETURNING id, user_id, type, media_id, poster, title, status, episodes, duration, created_at, updated_at
`

type AddToListParams struct {
	UserID   uuid.UUID  `json:"userId"`
	Type     ValidTypes `binding:"required" json:"type"`
	MediaID  *string    `binding:"required" json:"mediaId"`
	Poster   string     `binding:"required" json:"poster"`
	Title    string     `binding:"required" json:"title"`
	Status   Status     `binding:"required" json:"status"`
	Episodes *int32     `json:"episodes"`
	Duration *int32     `json:"duration"`
}

func (q *Queries) AddToList(ctx context.Context, arg AddToListParams) (*WatchList, error) {
	row := q.db.QueryRow(ctx, addToList,
		arg.UserID,
		arg.Type,
		arg.MediaID,
		arg.Poster,
		arg.Title,
		arg.Status,
		arg.Episodes,
		arg.Duration,
	)
	var i WatchList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Type,
		&i.MediaID,
		&i.Poster,
		&i.Title,
		&i.Status,
		&i.Episodes,
		&i.Duration,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const deleteWatchList = `-- name: DeleteWatchList :one
DELETE FROM watch_list
  WHERE id = $1
RETURNING id, user_id, type, media_id, poster, title, status, episodes, duration, created_at, updated_at
`

func (q *Queries) DeleteWatchList(ctx context.Context, id uuid.UUID) (*WatchList, error) {
	row := q.db.QueryRow(ctx, deleteWatchList, id)
	var i WatchList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Type,
		&i.MediaID,
		&i.Poster,
		&i.Title,
		&i.Status,
		&i.Episodes,
		&i.Duration,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getUserWatchList = `-- name: GetUserWatchList :many
SELECT id, user_id, type, media_id, poster, title, status, episodes, duration, created_at, updated_at FROM watch_list
WHERE user_id = $1
`

func (q *Queries) GetUserWatchList(ctx context.Context, userID uuid.UUID) ([]*WatchList, error) {
	rows, err := q.db.Query(ctx, getUserWatchList, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*WatchList
	for rows.Next() {
		var i WatchList
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Type,
			&i.MediaID,
			&i.Poster,
			&i.Title,
			&i.Status,
			&i.Episodes,
			&i.Duration,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getWatchListByMediaID = `-- name: GetWatchListByMediaID :one
SELECT id, user_id, type, media_id, poster, title, status, episodes, duration, created_at, updated_at FROM watch_list
WHERE media_id = $1
`

func (q *Queries) GetWatchListByMediaID(ctx context.Context, mediaID *string) (*WatchList, error) {
	row := q.db.QueryRow(ctx, getWatchListByMediaID, mediaID)
	var i WatchList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Type,
		&i.MediaID,
		&i.Poster,
		&i.Title,
		&i.Status,
		&i.Episodes,
		&i.Duration,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateWatchListStatus = `-- name: UpdateWatchListStatus :one
UPDATE watch_list
SET
    status = COALESCE($1, status),
    updated_at = now()
WHERE id = $2
RETURNING id, user_id, type, media_id, poster, title, status, episodes, duration, created_at, updated_at
`

type UpdateWatchListStatusParams struct {
	Status Status    `binding:"required" json:"status"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) UpdateWatchListStatus(ctx context.Context, arg UpdateWatchListStatusParams) (*WatchList, error) {
	row := q.db.QueryRow(ctx, updateWatchListStatus, arg.Status, arg.ID)
	var i WatchList
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Type,
		&i.MediaID,
		&i.Poster,
		&i.Title,
		&i.Status,
		&i.Episodes,
		&i.Duration,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
