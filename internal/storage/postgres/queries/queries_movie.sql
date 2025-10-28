-- noinspection SqlResolve
-- name: GetMovie :one
SELECT *
FROM movies
WHERE id = $1;

-- noinspection SqlResolve
-- name: AddMovie :one
INSERT INTO movies (title, synopsis, released_at, duration_minutes, poster_url)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- noinspection SqlResolve
-- name: UpdateMovie :one
UPDATE movies
SET title = $2,
    synopsis = $3,
    released_at = $4,
    duration_minutes = $5,
    poster_url = $6
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: DeleteMovie :one
DELETE FROM movies
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: ListMovies :many
SELECT *
FROM movies
ORDER BY released_at DESC
LIMIT $1
OFFSET $2;

-- noinspection SqlResolve
-- name: ListGenreMovies :many
SELECT mov.*
FROM genres AS gen
LEFT JOIN categories AS cat
ON (cat.genre_id = gen.id)
LEFT JOIN movies AS mov
ON (cat.movie_id = mov.id)
WHERE gen.id = $1
LIMIT $2
OFFSET $3;

-- noinspection SqlResolve
-- name: ListCelebrityMovies :many
SELECT mov.*
FROM roles AS rol
JOIN movies AS mov
ON (mov.id = rol.movie_id)
WHERE rol.celebrity_id = $1
ORDER BY mov.released_at DESC
LIMIT $2
OFFSET $3;

