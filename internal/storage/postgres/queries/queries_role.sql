-- name: GetRole :one
SELECT *
FROM roles
WHERE movie_id = $1 AND celebrity_id = $2;

-- name: AddRole :one
INSERT INTO roles (movie_id, celebrity_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateRole :one
UPDATE roles
SET role = $3
WHERE movie_id = $1 AND celebrity_id = $2
RETURNING *;

-- name: DeleteRole :one
DELETE FROM roles
WHERE movie_id = $1 AND celebrity_id = $2
RETURNING *;

-- name: ListCelebrityRoles :many
SELECT rol.*
FROM celebrities AS cel
LEFT JOIN roles AS rol
ON (rol.celebrity_id = cel.id)
WHERE cel.id = $1
LIMIT $2
OFFSET $3;

-- name: ListMovieRoles :many
SELECT rol.*
FROM movies AS mov
LEFT JOIN roles AS rol
ON (rol.movie_id = mov.id)
WHERE mov.id = $1
LIMIT $2
OFFSET $3;