-- name: CreateClientCMS :one
INSERT INTO sport.client_cms (
  hero_section,
  about_section
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetClientCMS :one
SELECT * FROM sport.client_cms
WHERE id = $1 LIMIT 1;

-- name: ListClientCMS :many
SELECT * FROM sport.client_cms
WHERE client_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;