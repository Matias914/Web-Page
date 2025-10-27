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

-- name: DeleteGenre :one
DELETE FROM genres
WHERE id = $1
RETURNING *;

-- name: ListGenres :many
SELECT *
FROM genres
LIMIT $1
OFFSET $2;

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