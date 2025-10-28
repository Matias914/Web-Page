-- noinspection SqlResolve
-- name: GetGenre :one
SELECT *
FROM genres
WHERE id = $1;

-- noinspection SqlResolve
-- name: AddGenre :one
INSERT INTO genres (name)
VALUES ($1)
RETURNING *;

-- noinspection SqlResolve
-- name: UpdateGenre :one
UPDATE genres
SET name = $2
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: DeleteGenre :one
DELETE FROM genres
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: ListGenres :many
SELECT *
FROM genres
LIMIT $1
OFFSET $2;

-- noinspection SqlResolve
-- name: ListMovieGenres :many
SELECT gen.*
FROM movies AS mov
LEFT JOIN categories AS cat
ON (cat.movie_id = mov.id)
LEFT JOIN genres AS gen
ON (cat.genre_id = gen.id)
WHERE mov.id = $1
LIMIT $2
OFFSET $3;