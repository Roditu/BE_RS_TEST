-- name: GetPerson :one
SELECT * FROM person
WHERE id = $1 LIMIT 1;

-- name: CreatePerson :one
INSERT INTO person (
  name, ambition
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPersonByName :one
SELECT id, name, ambition
FROM person
WHERE name = $1;