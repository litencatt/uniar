// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: live.sql

package repository

import (
	"context"
)

const getLiveList = `-- name: GetLiveList :many
SELECT id, name FROM lives
`

type GetLiveListRow struct {
	ID   int64
	Name string
}

func (q *Queries) GetLiveList(ctx context.Context, db DBTX) ([]GetLiveListRow, error) {
	rows, err := db.QueryContext(ctx, getLiveList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLiveListRow
	for rows.Next() {
		var i GetLiveListRow
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

const registLive = `-- name: RegistLive :exec
INSERT INTO lives (name) VALUES (?)
`

func (q *Queries) RegistLive(ctx context.Context, db DBTX, name string) error {
	_, err := db.ExecContext(ctx, registLive, name)
	return err
}
