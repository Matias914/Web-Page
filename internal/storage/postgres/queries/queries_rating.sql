-- name: GetRating :one
SELECT *
FROM ratings
WHERE user_id = $1 AND movie_id = $2;

-- name: AddRating :one
INSERT INTO ratings (user_id, movie_id, rating)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateRating :one
UPDATE ratings
SET rating = $3
WHERE user_id = $1 AND movie_id = $2
RETURNING *;

-- name: DeleteRating :one
DELETE FROM ratings
WHERE user_id = $1 AND movie_id = $2
RETURNING *;

-- name: ListMovieRatings :many
SELECT rat.*
FROM movies AS mov
LEFT JOIN ratings AS rat
ON (rat.movie_id = mov.id)
WHERE mov.id = $1
ORDER BY rat.rating DESC
LIMIT $2
OFFSET $3;

-- name: ListUserRatings :many
SELECT rat.*
FROM users AS usr
LEFT JOIN ratings AS rat
ON (rat.user_id = usr.id)
WHERE usr.id = $1
ORDER BY rat.created_at DESC
LIMIT $2
OFFSET $3;