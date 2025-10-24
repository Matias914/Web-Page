-- name: GetCategory :one
SELECT *
FROM categories
WHERE genre_id = $1 AND movie_id = $2;

-- name: AddCategory :exec
INSERT INTO categories (genre_id, movie_id)
VALUES ($1, $2);

-- name: DeleteCategory :one
DELETE FROM categories
WHERE genre_id = $1 AND movie_id = $2
RETURNING *;