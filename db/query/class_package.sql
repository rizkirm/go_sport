-- name: CreateClassPackage :one
INSERT INTO sport.class_package (
  code,
  name,
  price,
  customer_type,
  type,
  description,
  path_photo
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetClassPackage :one
SELECT * FROM sport.class_package
WHERE id = $1 LIMIT 1;

-- name: ListClassPackage :many
SELECT * FROM sport.class_package
WHERE class_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;