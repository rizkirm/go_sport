-- name: CreateNews :one
INSERT INTO sport.news (
  code,
  title,
  description,
  path_photo
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetNews :one
SELECT * FROM sport.news
WHERE id = $1 LIMIT 1;

-- name: ListNews :many
SELECT * FROM sport.news
WHERE client_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;