-- name: GetCelebrity :one
SELECT *
FROM celebrities
WHERE id = $1;

-- name: AddCelebrity :one
INSERT INTO celebrities (name, birth_date)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateCelebrity :one
UPDATE celebrities
SET name = $2, birth_date = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCelebrity :exec
DELETE FROM celebrities
WHERE id = $1;

-- name: ListCelebrities :many
SELECT *
FROM celebrities
LIMIT $1
OFFSET $2;

-- name: ListMovieCelebrities :many
SELECT cel.*
FROM roles AS rol
JOIN celebrities AS cel
ON (cel.id = rol.celebrity_id)
WHERE rol.movie_id = $1
ORDER BY cel.name
LIMIT $2
OFFSET $3;