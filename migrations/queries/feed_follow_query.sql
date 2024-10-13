-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow
WHERE feed_id = $1
    AND user_id = $2;
-- name: CreateFeedFollow :exec
INSERT INTO feed_follow (user_id, feed_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetFollowedFeeds :many
SELECT *
FROM feed_follow
WHERE user_id = $1
LIMIT $2 OFFSET $3;
-- name: GetFeedFollowers :many
SELECT *
FROM feed_follow
WHERE feed_id = $1;
-- name: GetFeedFollowerCount :one
SELECT COUNT(user_id)
FROM feed_follow
WHERE feed_id = $1;