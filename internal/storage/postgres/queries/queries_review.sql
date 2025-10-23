-- name: GetReview :one
SELECT *
FROM reviews
WHERE user_id = $1 AND movie_id = $2;

-- name: AddReview :one
INSERT INTO reviews (user_id, movie_id, comment, created_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateReview :one
UPDATE reviews
SET comment = $3
WHERE user_id = $1 AND movie_id = $2
RETURNING *;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE user_id = $1 AND movie_id = $2;

-- name: ListMovieReviews :many
SELECT rev.*
FROM reviews AS rev
JOIN users AS usr
ON (rev.user_id = usr.id)
WHERE rev.movie_id = $1
LIMIT $2
OFFSET $3;

-- name: ListUserReviews :many
SELECT rev.*
FROM reviews AS rev
JOIN users AS usr
ON (rev.user_id = usr.id)
WHERE rev.user_id = $1
LIMIT $2
OFFSET $3;