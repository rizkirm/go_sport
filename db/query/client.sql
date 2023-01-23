-- name: CreateClient :one
INSERT INTO sport.client (
    code, name, domain_local, domain_production, whatsapp_number, whatsapp_message, whatsapp_link,
    fb_link, ig_link, path_logo
) VALUES (
    $1,$2,$3,$4,$5,$6,$7,$8,$9,$10
) RETURNING *;

-- name: GetClient :one
SELECT * FROM sport.client
WHERE id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM sport.client
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateClient :one
UPDATE sport.client
SET 
name = $2,
domain_local = $3, 
domain_production = $4, 
whatsapp_number = $5,
whatsapp_message = $6,
whatsapp_link = $7,
fb_link = $8,
ig_link = $9,
path_logo = $10
WHERE id = $1
RETURNING *;

-- name: DeleteClient :exec
DELETE FROM sport.client 
WHERE id = $1;