-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, title, description, url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdateFeed :one
UPDATE feeds
SET title = $1,
    description = $2,
    url = $3,
    updated_at = timezone('utc', now())
WHERE user_id = $4
RETURNING *;
-- name: DeleteFeed :exec
DELETE FROM feeds
WHERE id = $1
    AND user_id = $2;
-- name: GetAllFeeds :many
SELECT *
FROM feeds
LIMIT $1 OFFSET $2;
-- name: GetAllFeedsByUserID :many
SELECT *
FROM feeds
WHERE user_id = $1
LIMIT $2 OFFSET $3;
-- name: GetFeedByID :one
SELECT *
FROM feeds
WHERE id = $1;
-- name: GetNextFeedToFetch :many
SELECT id,
    url,
    last_publisher_published_at
FROM feeds
WHERE last_fetched_at < $1
    OR last_fetched_at IS NULL
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $2;
-- name: UpdateLastFetchedAt :exec
UPDATE feeds
SET last_fetched_at = timezone('utc', now()),
    updated_at = timezone('utc', now())
WHERE id = $1;
-- name: UpdateLastPublisherPublishedAt :exec
UPDATE feeds
SET last_publisher_published_at = $1,
    updated_at = timezone('utc', now())
WHERE id = $2;