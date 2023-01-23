-- name: CreateClass :one
INSERT INTO sport.class (
  client_id,
  code,
  name,
  description
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetClass :one
SELECT * FROM sport.class
WHERE id = $1 LIMIT 1;

-- name: ListClass :many
SELECT * FROM sport.class
ORDER BY id
LIMIT $1
OFFSET $2;

