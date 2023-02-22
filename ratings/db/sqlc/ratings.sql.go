// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: ratings.sql

package db

import (
	"context"
	"database/sql"
)

const createRating = `-- name: CreateRating :one
INSERT INTO ratings(
    movie_id,
    user_id,
    score
)    VALUES(
        $1, $2, $3
    ) RETURNING rating_id, score, user_id, movie_id
`

type CreateRatingParams struct {
	MovieID sql.NullInt64 `json:"movie_id"`
	UserID  sql.NullInt64 `json:"user_id"`
	Score   sql.NullInt32 `json:"score"`
}

func (q *Queries) CreateRating(ctx context.Context, arg CreateRatingParams) (Rating, error) {
	row := q.db.QueryRowContext(ctx, createRating, arg.MovieID, arg.UserID, arg.Score)
	var i Rating
	err := row.Scan(
		&i.RatingID,
		&i.Score,
		&i.UserID,
		&i.MovieID,
	)
	return i, err
}

const deleteRating = `-- name: DeleteRating :exec
DELETE FROM ratings
WHERE movie_id = $1 AND user_id = $2
`

type DeleteRatingParams struct {
	MovieID sql.NullInt64 `json:"movie_id"`
	UserID  sql.NullInt64 `json:"user_id"`
}

func (q *Queries) DeleteRating(ctx context.Context, arg DeleteRatingParams) error {
	_, err := q.db.ExecContext(ctx, deleteRating, arg.MovieID, arg.UserID)
	return err
}

const getRatingDetail = `-- name: GetRatingDetail :one
SELECT rating_id, score, user_id, movie_id FROM ratings
WHERE movie_id = $1 AND user_id = $2
`

type GetRatingDetailParams struct {
	MovieID sql.NullInt64 `json:"movie_id"`
	UserID  sql.NullInt64 `json:"user_id"`
}

func (q *Queries) GetRatingDetail(ctx context.Context, arg GetRatingDetailParams) (Rating, error) {
	row := q.db.QueryRowContext(ctx, getRatingDetail, arg.MovieID, arg.UserID)
	var i Rating
	err := row.Scan(
		&i.RatingID,
		&i.Score,
		&i.UserID,
		&i.MovieID,
	)
	return i, err
}

const getUserRatings = `-- name: GetUserRatings :many
SELECT score FROM ratings
WHERE user_id = $1
`

func (q *Queries) GetUserRatings(ctx context.Context, userID sql.NullInt64) ([]sql.NullInt32, error) {
	rows, err := q.db.QueryContext(ctx, getUserRatings, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullInt32
	for rows.Next() {
		var score sql.NullInt32
		if err := rows.Scan(&score); err != nil {
			return nil, err
		}
		items = append(items, score)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listMovieRatings = `-- name: ListMovieRatings :many
SELECT score FROM ratings
WHERE movie_id = $1
`

func (q *Queries) ListMovieRatings(ctx context.Context, movieID sql.NullInt64) ([]sql.NullInt32, error) {
	rows, err := q.db.QueryContext(ctx, listMovieRatings, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullInt32
	for rows.Next() {
		var score sql.NullInt32
		if err := rows.Scan(&score); err != nil {
			return nil, err
		}
		items = append(items, score)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateRating = `-- name: UpdateRating :one
UPDATE ratings
SET score = $3
WHERE movie_id = $1 AND user_id = $2
RETURNING rating_id, score, user_id, movie_id
`

type UpdateRatingParams struct {
	MovieID sql.NullInt64 `json:"movie_id"`
	UserID  sql.NullInt64 `json:"user_id"`
	Score   sql.NullInt32 `json:"score"`
}

func (q *Queries) UpdateRating(ctx context.Context, arg UpdateRatingParams) (Rating, error) {
	row := q.db.QueryRowContext(ctx, updateRating, arg.MovieID, arg.UserID, arg.Score)
	var i Rating
	err := row.Scan(
		&i.RatingID,
		&i.Score,
		&i.UserID,
		&i.MovieID,
	)
	return i, err
}
