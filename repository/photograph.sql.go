// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: photograph.sql

package repository

import (
	"context"
)

const getPhotographByGroupId = `-- name: GetPhotographByGroupId :many
;

SELECT
    id,
    name
FROM
    photograph
WHERE
    group_id = ?
ORDER BY
	name_for_order ASC
`

type GetPhotographByGroupIdRow struct {
	ID   int64
	Name string
}

func (q *Queries) GetPhotographByGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetPhotographByGroupIdRow, error) {
	rows, err := db.QueryContext(ctx, getPhotographByGroupId, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPhotographByGroupIdRow
	for rows.Next() {
		var i GetPhotographByGroupIdRow
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

const getPhotographList = `-- name: GetPhotographList :many
;

SELECT
    id,
    name
FROM
    photograph
WHERE
    group_id = ?
    AND photo_type = ?
ORDER BY
    group_id, id ASC
`

type GetPhotographListParams struct {
	GroupID   int64
	PhotoType string
}

type GetPhotographListRow struct {
	ID   int64
	Name string
}

func (q *Queries) GetPhotographList(ctx context.Context, db DBTX, arg GetPhotographListParams) ([]GetPhotographListRow, error) {
	rows, err := db.QueryContext(ctx, getPhotographList, arg.GroupID, arg.PhotoType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPhotographListRow
	for rows.Next() {
		var i GetPhotographListRow
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

const getPhotographListAll = `-- name: GetPhotographListAll :many
SELECT
    name,
    photo_type,
    released_at
FROM
    photograph
ORDER BY released_at ASC
`

type GetPhotographListAllRow struct {
	Name       string
	PhotoType  string
	ReleasedAt interface{}
}

func (q *Queries) GetPhotographListAll(ctx context.Context, db DBTX) ([]GetPhotographListAllRow, error) {
	rows, err := db.QueryContext(ctx, getPhotographListAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPhotographListAllRow
	for rows.Next() {
		var i GetPhotographListAllRow
		if err := rows.Scan(&i.Name, &i.PhotoType, &i.ReleasedAt); err != nil {
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

const getPhotographListByPhotoType = `-- name: GetPhotographListByPhotoType :many
;

SELECT
    id,
    name
FROM
    photograph
WHERE
    photo_type = ?
ORDER BY
    group_id, id ASC
`

type GetPhotographListByPhotoTypeRow struct {
	ID   int64
	Name string
}

func (q *Queries) GetPhotographListByPhotoType(ctx context.Context, db DBTX, photoType string) ([]GetPhotographListByPhotoTypeRow, error) {
	rows, err := db.QueryContext(ctx, getPhotographListByPhotoType, photoType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPhotographListByPhotoTypeRow
	for rows.Next() {
		var i GetPhotographListByPhotoTypeRow
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

const getSsrPlusReleasedPhotographList = `-- name: GetSsrPlusReleasedPhotographList :many
;

SELECT
	photograph_id
FROM
	scenes
WHERE
	ssr_plus = 1
GROUP BY
	photograph_id
`

func (q *Queries) GetSsrPlusReleasedPhotographList(ctx context.Context, db DBTX) ([]int64, error) {
	rows, err := db.QueryContext(ctx, getSsrPlusReleasedPhotographList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var photograph_id int64
		if err := rows.Scan(&photograph_id); err != nil {
			return nil, err
		}
		items = append(items, photograph_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const registPhotograph = `-- name: RegistPhotograph :exec
INSERT INTO photograph (name, group_id, photo_type) VALUES (?, ?, ?)
`

type RegistPhotographParams struct {
	Name      string
	GroupID   int64
	PhotoType string
}

func (q *Queries) RegistPhotograph(ctx context.Context, db DBTX, arg RegistPhotographParams) error {
	_, err := db.ExecContext(ctx, registPhotograph, arg.Name, arg.GroupID, arg.PhotoType)
	return err
}
