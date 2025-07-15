-- name: CreateTask :one
INSERT INTO tasks (title, description, status)
VALUES ($1, $2, $3)
RETURNING id;

-- name: GetTasksFiltered :many
SELECT id, title, description, status
FROM tasks
WHERE (cardinality(COALESCE($1::uuid[], '{}')) = 0 OR id = ANY($1::uuid[]))
  AND (cardinality(COALESCE($2::task_status[], '{}')) = 0 OR status = ANY($2::task_status[]))
ORDER BY id
LIMIT $3 OFFSET $4;

-- name: UpdateTask :one
UPDATE tasks
SET
  title = COALESCE(sqlc.narg('title'), title),
  description = COALESCE(sqlc.narg('description'), description),
  status = COALESCE(sqlc.narg('status'), status)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
