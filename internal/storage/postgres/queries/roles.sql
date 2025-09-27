-- name: AddRole :one
INSERT INTO roles (movie_id, celebrity_id, role)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRoleByMovieId :one
SELECT *
FROM roles
WHERE movie_id = $1 AND celebrity_id = $2;

-- name: ListMoviesFromCelebrityId :many
SELECT *
FROM roles AS rol
JOIN movies AS mov
ON (mov.id = rol.movie_id)
WHERE rol.celebrity_id = $1
ORDER BY mov.released_at DESC
LIMIT $2
OFFSET $3;

-- name: ListCelebritiesFromMovieId :many
SELECT *
FROM roles AS rol
JOIN celebrities AS cel
ON (cel.id = rol.celebrity_id)
WHERE rol.movie_id = $1
ORDER BY cel.name
LIMIT $2
OFFSET $3;

-- name: DeleteRoleByMovieId :exec
DELETE FROM roles
WHERE movie_id = $1 AND celebrity_id = $2;