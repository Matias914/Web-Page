-- name: AddRating :one
INSERT INTO ratings (user_id, movie_id, rating)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRating :one
SELECT *
FROM ratings
WHERE user_id = $1 AND movie_id = $2;

-- name: ListUsersRatingsFromMovieId :many
SELECT *
FROM ratings AS rat
JOIN users AS usr
ON (usr.id = rat.user_id)
WHERE rat.movie_id = $1
ORDER BY rat.rating DESC
LIMIT $2
OFFSET $3;

-- name: ListMoviesRatingsFromUserId :many
SELECT *
FROM ratings AS rat
JOIN movies AS mov
ON (mov.id = rat.movie_id)
WHERE rat.user_id = $1
ORDER BY rat.created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateRating :one
UPDATE ratings
SET rating = $3
WHERE user_id = $1 AND movie_id = $2
RETURNING *;

-- name: DeleteRating :exec
DELETE FROM ratings
WHERE user_id = $1 AND movie_id = $2;