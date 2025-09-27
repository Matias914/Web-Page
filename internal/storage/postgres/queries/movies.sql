-- name: AddMovie :one
INSERT INTO movies (title, synopsis, released_at, poster_url, duration_minutes)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetMovieById :one
SELECT *
FROM movies
WHERE id = $1;

-- name: ListMoviesByReleaseDate :many
SELECT *
FROM movies
ORDER BY released_at DESC
LIMIT $1
OFFSET $2;

-- name: UpdateMovieById :one
UPDATE movies
SET title = $2,
    synopsis = $3,
    released_at = $4,
    poster_url = $5,
    duration_minutes = $6
WHERE id = $1
RETURNING *;

-- name: DeleteMovieById :exec
DELETE FROM movies
WHERE id = $1;