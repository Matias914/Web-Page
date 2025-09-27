-- name: AddCategory :one
INSERT INTO categories (genre_id, movie_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetCategoryById :one
SELECT *
FROM categories
WHERE genre_id = $1 AND movie_id = $2;

-- name: ListGenresByMovieId :many
SELECT *
FROM categories AS cat
JOIN genres AS gen
ON (cat.id = gen.genre_id)
WHERE cat.movie_id = $1
LIMIT $2
OFFSET $3;

-- name: ListMoviesByGenreId :many
SELECT *
FROM categories AS cat
JOIN movies AS mov
ON (mov.id = cat.movie_id)
WHERE cat.genre_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateCategoryById :one
UPDATE categories
SET genre_id = $3, movie_id = $4
WHERE genre_id = $1 AND movie_id = $2
RETURNING *;

-- name: DeleteCategoryById :exec
DELETE FROM categories
WHERE genre_id = $1 AND movie_id = $2;