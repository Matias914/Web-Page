-- noinspection SqlResolve
-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- noinspection SqlResolve
-- name: AddUser :one
INSERT INTO users (username, password, mail)
VALUES ($1, $2, $3)
RETURNING *;

-- noinspection SqlResolve
-- name: UpdateUser :one
UPDATE users
SET username = $2, mail = $3
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: ListUsers :many
SELECT *
FROM users
LIMIT $1
OFFSET $2;