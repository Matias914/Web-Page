-- name: GetMovie :one
SELECT * FROM movies WHERE id = $1;

-- name: ListMovies :many
SELECT * FROM movies ORDER BY created_at DESC;

-- name: CreateMovie :one
INSERT INTO movies (title, director, release_year)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateMovie :one
UPDATE movies
SET title = $1, director = $2, release_year = $3
WHERE id = $4
RETURNING *;

-- name: DeleteMovie :exec
DELETE FROM movies WHERE id = $1;

-- name: SearchMovies :many
SELECT * FROM movies
WHERE title ILIKE '%' || $1 || '%'
OR director ILIKE '%' || $1 || '%';