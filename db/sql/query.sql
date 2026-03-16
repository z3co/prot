
-- name: CreateTodo :one
INSERT INTO todos (
	folder, description
) VALUES (
	?, ?
)
RETURNING *;

-- name: GetTodosByFolder :many
SELECT * FROM todos
WHERE folder = ?;

-- name: GetUncompleteTodosByFolder :many
SELECT * FROM todos
WHERE folder = ? AND done = 0;

-- name: GetTodoById :one
SELECT * FROM todos
WHERE id = ? LIMIT 1;

-- name: CompleteTodo :exec
UPDATE todos
set done = 1
WHERE id = ?;

-- name: UnCompleteTodo :exec
UPDATE todos
set done = 0
WHERE id = ?;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = ?;
