
-- name: AddCategory :one
INSERT INTO categories (genre_id, movie_id)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE genre_id = $1 AND movie_id = $2;