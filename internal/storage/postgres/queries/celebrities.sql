-- name: AddCelebrity :one
INSERT INTO celebrities (name, birth_date)
VALUES ($1, $2)
RETURNING *;

-- name: GetCelebrityById :one
SELECT *
FROM celebrities
WHERE id = $1;

-- name: ListCelebrities :many
SELECT *
FROM celebrities
LIMIT $1
OFFSET $2;

-- name: UpdateCelebrityById :one
UPDATE celebrities
SET name = $2, birth_date = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCelebrityById :exec
DELETE FROM celebrities
WHERE id = $1;