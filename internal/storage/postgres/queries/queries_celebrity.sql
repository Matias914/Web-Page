-- noinspection SqlResolve
-- name: GetCelebrity :one
SELECT *
FROM celebrities
WHERE id = $1;

-- noinspection SqlResolve
-- name: AddCelebrity :one
INSERT INTO celebrities (name, birth_date)
VALUES ($1, $2)
RETURNING *;

-- noinspection SqlResolve
-- name: UpdateCelebrity :one
UPDATE celebrities
SET name = $2, birth_date = $3
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: DeleteCelebrity :one
DELETE FROM celebrities
WHERE id = $1
RETURNING *;

-- noinspection SqlResolve
-- name: ListCelebrities :many
SELECT *
FROM celebrities
LIMIT $1
OFFSET $2;

-- noinspection SqlResolve
-- name: ListMovieCelebrities :many
SELECT cel.*
FROM movies AS mov
LEFT JOIN roles AS rol
ON (mov.id = rol.movie_id)
LEFT JOIN celebrities AS cel
ON (cel.id = rol.celebrity_id)
WHERE mov.id = $1
ORDER BY cel.name
LIMIT $2
OFFSET $3;