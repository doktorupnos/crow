-- name: CreatePost :one
INSERT INTO posts (
    id, created_at, updated_at, body, user_id
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPosts :many
SELECT
    posts.*,
    users.name AS user_name
FROM posts
INNER JOIN users ON posts.user_id = users.id
ORDER BY posts.created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetPostsByUser :many
SELECT
    posts.*,
    users.name AS user_name
FROM posts
INNER JOIN users ON posts.user_id = users.id
WHERE users.name = $1
ORDER BY posts.created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1 AND user_id = $2;
