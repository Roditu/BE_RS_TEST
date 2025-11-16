-- name: CreateTask :one
INSERT INTO task (
  todo, exp, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTaskByUserId :one
SELECT *
FROM task
WHERE user_id = $1
LIMIT 1;

-- name: GetUndoneTaskByUserId :one
SELECT *
FROM task
WHERE user_id = $1 AND status = "UNFINISHED"
LIMIT 1;

-- name: ListTasksByUser :many
SELECT * FROM task
WHERE user_id = $1
ORDER BY task_id DESC;

-- name: UpdateTaskStatus :one
UPDATE task
SET status = $2
WHERE task_id = $1
RETURNING *;

-- name: GetCompleteTask :one
UPDATE task
SET status = "COMPLETE"
WHERE user_id = $1
RETURNING *;