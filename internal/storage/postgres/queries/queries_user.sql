-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1;

-- name: AddUser :one
INSERT INTO users (username, mail)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET username = $2,
    mail = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT *
FROM users
LIMIT $1
OFFSET $2;