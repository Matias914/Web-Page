-- name: GetGenre :one
SELECT *
FROM genres
WHERE id = $1;

-- name: AddGenre :one
INSERT INTO genres (name)
VALUES ($1)
RETURNING *;

-- name: UpdateGenre :one
UPDATE genres
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = $1;

-- name: ListGenres :many
SELECT *
FROM genres
LIMIT $1
OFFSET $2;

-- name: ListMovieGenres :many
SELECT gen.*
FROM categories AS cat
JOIN genres AS gen
ON (cat.genre_id = gen.id)
WHERE cat.movie_id = $1
LIMIT $2
OFFSET $3;