// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: class_package.sql

package db

import (
	"context"
)

const createClassPackage = `-- name: CreateClassPackage :one
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
) RETURNING id, class_id, code, name, price, customer_type, type, description, path_photo, created_at, updated_at, deleted_at
`

type CreateClassPackageParams struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	CustomerType int32  `json:"customer_type"`
	Type         int32  `json:"type"`
	Description  string `json:"description"`
	PathPhoto    string `json:"path_photo"`
}

func (q *Queries) CreateClassPackage(ctx context.Context, arg CreateClassPackageParams) (SportClassPackage, error) {
	row := q.db.QueryRowContext(ctx, createClassPackage,
		arg.Code,
		arg.Name,
		arg.Price,
		arg.CustomerType,
		arg.Type,
		arg.Description,
		arg.PathPhoto,
	)
	var i SportClassPackage
	err := row.Scan(
		&i.ID,
		&i.ClassID,
		&i.Code,
		&i.Name,
		&i.Price,
		&i.CustomerType,
		&i.Type,
		&i.Description,
		&i.PathPhoto,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getClassPackage = `-- name: GetClassPackage :one
SELECT id, class_id, code, name, price, customer_type, type, description, path_photo, created_at, updated_at, deleted_at FROM sport.class_package
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetClassPackage(ctx context.Context, id int32) (SportClassPackage, error) {
	row := q.db.QueryRowContext(ctx, getClassPackage, id)
	var i SportClassPackage
	err := row.Scan(
		&i.ID,
		&i.ClassID,
		&i.Code,
		&i.Name,
		&i.Price,
		&i.CustomerType,
		&i.Type,
		&i.Description,
		&i.PathPhoto,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listClassPackage = `-- name: ListClassPackage :many
SELECT id, class_id, code, name, price, customer_type, type, description, path_photo, created_at, updated_at, deleted_at FROM sport.class_package
WHERE class_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListClassPackageParams struct {
	ClassID int32 `json:"class_id"`
	Limit   int32 `json:"limit"`
	Offset  int32 `json:"offset"`
}

func (q *Queries) ListClassPackage(ctx context.Context, arg ListClassPackageParams) ([]SportClassPackage, error) {
	rows, err := q.db.QueryContext(ctx, listClassPackage, arg.ClassID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SportClassPackage{}
	for rows.Next() {
		var i SportClassPackage
		if err := rows.Scan(
			&i.ID,
			&i.ClassID,
			&i.Code,
			&i.Name,
			&i.Price,
			&i.CustomerType,
			&i.Type,
			&i.Description,
			&i.PathPhoto,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
