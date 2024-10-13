-- +goose Up
ALTER TABLE feeds
ADD COLUMN last_publisher_published_at TIMESTAMP;
-- +goose Down
ALTER TABLE feeds DROP COLUMN last_publisher_published_at;