-- name: SignUp :one
INSERT INTO users (id, name)
VALUES ($1, $2)
RETURNING *;
-- name: SignIn :one
SELECT *
FROM users
WHERE name = $1;
-- name: GetUserByApiKey :one
SELECT *
FROM users
WHERE api_key = $1;
-- name: GetAllUsers :many
SELECT *
FROM users;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;