// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: producer_scene.sql

package repository

import (
	"context"
	"database/sql"
)

const getProducerScenes = `-- name: GetProducerScenes :many
;

SELECT
    ps.producer_id,
    ps.photograph_id,
    ps.member_id,
    c.name AS color,
    p.name AS photograph,
    m.name AS member,
    s.ssr_plus,
    ps.have
FROM
    producer_scenes ps
    JOIN scenes s ON ps.photograph_id = s.photograph_id AND ps.member_id = s.member_id AND ps.ssr_plus = s.ssr_plus
    JOIN photograph p on ps.photograph_id = p.id
    JOIN members m on ps.member_id = m.id
    JOIN color_types c ON s.color_type_id = c.id
WHERE
    p.name LIKE ?
    AND m.name LIKE ?
ORDER BY
    p.id,
    m.phase,
    m.first_name
`

type GetProducerScenesParams struct {
	Name   string
	Name_2 string
}

type GetProducerScenesRow struct {
	ProducerID   int64
	PhotographID int64
	MemberID     int64
	Color        string
	Photograph   string
	Member       string
	SsrPlus      int64
	Have         sql.NullInt64
}

func (q *Queries) GetProducerScenes(ctx context.Context, db DBTX, arg GetProducerScenesParams) ([]GetProducerScenesRow, error) {
	rows, err := db.QueryContext(ctx, getProducerScenes, arg.Name, arg.Name_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProducerScenesRow
	for rows.Next() {
		var i GetProducerScenesRow
		if err := rows.Scan(
			&i.ProducerID,
			&i.PhotographID,
			&i.MemberID,
			&i.Color,
			&i.Photograph,
			&i.Member,
			&i.SsrPlus,
			&i.Have,
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

const getProducerScenesByGroupId = `-- name: GetProducerScenesByGroupId :many
SELECT
    ps.producer_id,
    ps.photograph_id,
    ps.member_id,
    c.name AS color,
    p.name AS photograph,
    m.name AS member,
    s.ssr_plus,
    CASE
      WHEN ps.have IS NULL then false
      WHEN ps.have = 0 then false
      WHEN ps.have = 1 then true
    end as ps_have
FROM
    producer_scenes ps
    JOIN scenes s ON ps.photograph_id = s.photograph_id AND ps.member_id = s.member_id AND ps.ssr_plus = s.ssr_plus
    JOIN photograph p on ps.photograph_id = p.id
    JOIN members m on ps.member_id = m.id
    JOIN color_types c ON s.color_type_id = c.id
WHERE
    m.group_id = ?
ORDER BY
    p.id,
    m.phase,
    m.first_name
`

type GetProducerScenesByGroupIdRow struct {
	ProducerID   int64
	PhotographID int64
	MemberID     int64
	Color        string
	Photograph   string
	Member       string
	SsrPlus      int64
	PsHave       interface{}
}

func (q *Queries) GetProducerScenesByGroupId(ctx context.Context, db DBTX, groupID int64) ([]GetProducerScenesByGroupIdRow, error) {
	rows, err := db.QueryContext(ctx, getProducerScenesByGroupId, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProducerScenesByGroupIdRow
	for rows.Next() {
		var i GetProducerScenesByGroupIdRow
		if err := rows.Scan(
			&i.ProducerID,
			&i.PhotographID,
			&i.MemberID,
			&i.Color,
			&i.Photograph,
			&i.Member,
			&i.SsrPlus,
			&i.PsHave,
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

const initProducerSceneAll = `-- name: InitProducerSceneAll :exec
;

DELETE FROM producer_scenes
WHERE
	producer_id = ?
    AND member_id = ?
`

type InitProducerSceneAllParams struct {
	ProducerID int64
	MemberID   int64
}

func (q *Queries) InitProducerSceneAll(ctx context.Context, db DBTX, arg InitProducerSceneAllParams) error {
	_, err := db.ExecContext(ctx, initProducerSceneAll, arg.ProducerID, arg.MemberID)
	return err
}

const registProducerScene = `-- name: RegistProducerScene :exec
;

INSERT OR REPLACE INTO producer_scenes (
	producer_id,
	photograph_id,
    member_id,
    ssr_plus,
    have
) VALUES (?, ?, ?, ?, ?)
`

type RegistProducerSceneParams struct {
	ProducerID   int64
	PhotographID int64
	MemberID     int64
	SsrPlus      int64
	Have         sql.NullInt64
}

func (q *Queries) RegistProducerScene(ctx context.Context, db DBTX, arg RegistProducerSceneParams) error {
	_, err := db.ExecContext(ctx, registProducerScene,
		arg.ProducerID,
		arg.PhotographID,
		arg.MemberID,
		arg.SsrPlus,
		arg.Have,
	)
	return err
}

const updateProducerScene = `-- name: UpdateProducerScene :exec
;

UPDATE
    producer_scenes
SET
    have = ?
WHERE
	producer_id = ?
	AND photograph_id = ?
    AND member_id = ?
`

type UpdateProducerSceneParams struct {
	Have         sql.NullInt64
	ProducerID   int64
	PhotographID int64
	MemberID     int64
}

func (q *Queries) UpdateProducerScene(ctx context.Context, db DBTX, arg UpdateProducerSceneParams) error {
	_, err := db.ExecContext(ctx, updateProducerScene,
		arg.Have,
		arg.ProducerID,
		arg.PhotographID,
		arg.MemberID,
	)
	return err
}
