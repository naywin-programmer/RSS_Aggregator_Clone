-- +goose Up
CREATE TABLE users (
    id UUID NOT NULL,
    name text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (timezone('utc', now())),
    updated_at TIMESTAMP NOT NULL DEFAULT (timezone('utc', now())),
    PRIMARY KEY(id)
);
-- +goose Down
DROP TABLE users;