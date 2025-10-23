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

-- name: DeleteRole :exec
DELETE FROM roles
WHERE movie_id = $1 AND celebrity_id = $2;