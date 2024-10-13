-- name: CreatePost :exec
INSERT INTO posts (
        id,
        feed_id,
        title,
        description,
        url,
        published_at
    )
VALUES ($1, $2, $3, $4, $5, $6);
-- name: GetPostsByUser :many
SELECT *
FROM posts
    LEFT JOIN feed_follow ON posts.feed_id = feed_follow.feed_id
WHERE feed_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2 OFFSET $3;