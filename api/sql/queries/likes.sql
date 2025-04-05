-- name: CreateLike :exec
INSERT INTO likes (
    user_id, post_id
) VALUES ($1, $2);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE user_id = $1 AND post_id = $2;

-- name: GetLikesForPost :one
SELECT COUNT(*)
FROM likes
WHERE post_id = $1;

-- name: UserLikesPost :one
SELECT EXISTS(
    SELECT 1
    FROM likes
    WHERE user_id = $1 AND post_id = $2
) AS has_liked;
