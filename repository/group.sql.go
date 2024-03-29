// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: group.sql

package repository

import (
	"context"
)

const getGroup = `-- name: GetGroup :many
SELECT id, name FROM groups
`

type GetGroupRow struct {
	ID   int64
	Name string
}

func (q *Queries) GetGroup(ctx context.Context, db DBTX) ([]GetGroupRow, error) {
	rows, err := db.QueryContext(ctx, getGroup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetGroupRow
	for rows.Next() {
		var i GetGroupRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const getGroupNameById = `-- name: GetGroupNameById :one
SELECT name FROM groups WHERE id = ?
`

func (q *Queries) GetGroupNameById(ctx context.Context, db DBTX, id int64) (string, error) {
	row := db.QueryRowContext(ctx, getGroupNameById, id)
	var name string
	err := row.Scan(&name)
	return name, err
}
