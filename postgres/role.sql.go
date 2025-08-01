// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: role.sql

package db

import (
	"context"
)

const getRoleByName = `-- name: GetRoleByName :one
SELECT role_id, name, description, created_at, updated_at FROM roles WHERE name = $1 LIMIT 1
`

func (q *Queries) GetRoleByName(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRow(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(
		&i.RoleID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listRoles = `-- name: ListRoles :many
SELECT role_id, name, description, created_at, updated_at FROM roles
`

func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.RoleID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
