-- name: CreateTask :one
INSERT INTO tasks (title, description, status)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetTasksFiltered :many
SELECT id, title, description, status
FROM tasks
WHERE (cardinality(COALESCE($1::uuid[], '{}')) = 0 OR id = ANY($1::uuid[]))
  AND (cardinality(COALESCE($2::task_status[], '{}')) = 0 OR status = ANY($2::task_status[]))
ORDER BY id;
