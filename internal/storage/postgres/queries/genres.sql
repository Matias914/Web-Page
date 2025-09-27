-- name: AddGenre :one
INSERT INTO genres (name)
VALUES ($1)
RETURNING *;

-- name: GetGenreById :one
SELECT *
FROM genres
WHERE id = $1;

-- name: ListGenres :many
SELECT *
FROM genres
LIMIT $1
OFFSET $2;

-- name: UpdateGenreById :one
UPDATE genres
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteGenreById :exec
DELETE FROM genres
WHERE id = $1;