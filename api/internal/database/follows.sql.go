// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: follows.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createFollow = `-- name: CreateFollow :exec
INSERT INTO follows (
    follower, followee
) VALUES ($1, $2)
`

type CreateFollowParams struct {
	Follower uuid.UUID `json:"follower"`
	Followee uuid.UUID `json:"followee"`
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) error {
	_, err := q.exec(ctx, q.createFollowStmt, createFollow, arg.Follower, arg.Followee)
	return err
}

const deleteFollow = `-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower = $1 AND followee = $2
`

type DeleteFollowParams struct {
	Follower uuid.UUID `json:"follower"`
	Followee uuid.UUID `json:"followee"`
}

func (q *Queries) DeleteFollow(ctx context.Context, arg DeleteFollowParams) error {
	_, err := q.exec(ctx, q.deleteFollowStmt, deleteFollow, arg.Follower, arg.Followee)
	return err
}

const followsUser = `-- name: FollowsUser :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower = $1 AND followee = $2
) AS follows_user
`

type FollowsUserParams struct {
	Follower uuid.UUID `json:"follower"`
	Followee uuid.UUID `json:"followee"`
}

func (q *Queries) FollowsUser(ctx context.Context, arg FollowsUserParams) (bool, error) {
	row := q.queryRow(ctx, q.followsUserStmt, followsUser, arg.Follower, arg.Followee)
	var follows_user bool
	err := row.Scan(&follows_user)
	return follows_user, err
}

const getFollowerCount = `-- name: GetFollowerCount :one
SELECT COUNT(*)
FROM follows
WHERE followee = $1
`

func (q *Queries) GetFollowerCount(ctx context.Context, followee uuid.UUID) (int64, error) {
	row := q.queryRow(ctx, q.getFollowerCountStmt, getFollowerCount, followee)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowers = `-- name: GetFollowers :many
SELECT
    follows.follower,
    users.name
FROM follows
INNER JOIN users ON follows.follower = users.id
WHERE follows.followee = $1
LIMIT $2 OFFSET $3
`

type GetFollowersParams struct {
	Followee uuid.UUID `json:"followee"`
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
}

type GetFollowersRow struct {
	Follower uuid.UUID `json:"follower"`
	Name     string    `json:"name"`
}

func (q *Queries) GetFollowers(ctx context.Context, arg GetFollowersParams) ([]GetFollowersRow, error) {
	rows, err := q.query(ctx, q.getFollowersStmt, getFollowers, arg.Followee, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFollowersRow
	for rows.Next() {
		var i GetFollowersRow
		if err := rows.Scan(&i.Follower, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowing = `-- name: GetFollowing :many
SELECT
    follows.followee,
    users.name
FROM follows
INNER JOIN users ON follows.followee = users.id
WHERE follows.follower = $1
LIMIT $2 OFFSET $3
`

type GetFollowingParams struct {
	Follower uuid.UUID `json:"follower"`
	Limit    int32     `json:"limit"`
	Offset   int32     `json:"offset"`
}

type GetFollowingRow struct {
	Followee uuid.UUID `json:"followee"`
	Name     string    `json:"name"`
}

func (q *Queries) GetFollowing(ctx context.Context, arg GetFollowingParams) ([]GetFollowingRow, error) {
	rows, err := q.query(ctx, q.getFollowingStmt, getFollowing, arg.Follower, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFollowingRow
	for rows.Next() {
		var i GetFollowingRow
		if err := rows.Scan(&i.Followee, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowingCount = `-- name: GetFollowingCount :one
SELECT COUNT(*)
FROM follows
WHERE follower = $1
`

func (q *Queries) GetFollowingCount(ctx context.Context, follower uuid.UUID) (int64, error) {
	row := q.queryRow(ctx, q.getFollowingCountStmt, getFollowingCount, follower)
	var count int64
	err := row.Scan(&count)
	return count, err
}
