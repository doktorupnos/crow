-- name: CreateFollow :exec
INSERT INTO follows (
    follower, followee
) VALUES ($1, $2);

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower = $1 AND followee = $2;

-- name: GetFollowerCount :one
SELECT COUNT(*)
FROM follows
WHERE followee = $1;

-- name: GetFollowingCount :one
SELECT COUNT(*)
FROM follows
WHERE follower = $1;

-- name: GetFollowers :many
SELECT
    follows.follower,
    users.name
FROM follows
INNER JOIN users ON follows.follower = users.id
WHERE follows.followee = $1
LIMIT $2 OFFSET $3;

-- name: GetFollowing :many
SELECT
    follows.followee,
    users.name
FROM follows
INNER JOIN users ON follows.followee = users.id
WHERE follows.follower = $1
LIMIT $2 OFFSET $3;

-- name: FollowsUser :one
SELECT EXISTS(
    SELECT 1
    FROM follows
    WHERE follower = $1 AND followee = $2
) AS follows_user;
