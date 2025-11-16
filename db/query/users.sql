-- name: CreateUser :one
INSERT INTO users (
  username, password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: GetUserByUserId :one
SELECT *
FROM users
WHERE user_id = $1;

-- name: AddUserExp :one
UPDATE users
SET exp = exp + $2
WHERE user_id = $1
RETURNING *;