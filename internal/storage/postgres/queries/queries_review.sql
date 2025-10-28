-- noinspection SqlResolve
-- name: GetReview :one
SELECT *
FROM reviews
WHERE user_id = $1 AND movie_id = $2;

-- noinspection SqlResolve
-- name: AddReview :one
INSERT INTO reviews (user_id, movie_id, comment, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- noinspection SqlResolve
-- name: UpdateReview :one
UPDATE reviews
SET comment = $3
WHERE user_id = $1 AND movie_id = $2
RETURNING *;

-- noinspection SqlResolve
-- name: DeleteReview :one
DELETE FROM reviews
WHERE user_id = $1 AND movie_id = $2
RETURNING *;

-- noinspection SqlResolve
-- name: ListMovieReviews :many
SELECT rev.*
FROM movies AS mov
LEFT JOIN reviews AS rev
ON (rev.movie_id = mov.id)
WHERE mov.id = $1
LIMIT $2
OFFSET $3;

-- noinspection SqlResolve
-- name: ListUserReviews :many
SELECT rev.*
FROM users AS usr
LEFT JOIN reviews AS rev
ON (rev.user_id = usr.id)
WHERE usr.id = $1
LIMIT $2
OFFSET $3;