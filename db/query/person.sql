-- name: GetAuthor :one
SELECT * FROM person
WHERE id = $1 LIMIT 1;

-- name: CreateAuthor :one
INSERT INTO person (
  name, ambition
) VALUES (
  $1, $2
)
RETURNING *;