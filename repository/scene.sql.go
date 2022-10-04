// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: scene.sql

package repository

import (
	"context"
	"database/sql"
)

const getAllScenes = `-- name: GetAllScenes :many
SELECT s.id FROM scenes s
`

func (q *Queries) GetAllScenes(ctx context.Context, db DBTX) ([]int64, error) {
	rows, err := db.QueryContext(ctx, getAllScenes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getScenesWithColor = `-- name: GetScenesWithColor :many
SELECT
	s.id,
	p.name AS photograph,
	p.abbreviation,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.performance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.performance_max,
	s.expected_value,
	s.ssr_plus,
	pm.bond_level_curent AS bonds,
	pm.discography_disc_total AS discography,
	ps.have
FROM
	scenes s
	JOIN photograph p ON s.photograph_id = p.id
	JOIN color_types c ON s.color_type_id = c.id
	JOIN members m ON s.member_id = m.id
	LEFT OUTER JOIN producer_members pm ON s.member_id = pm.member_id
	LEFT OUTER JOIN producer_scenes ps ON s.id = ps.scene_id
WHERE
	c.name LIKE ?
ORDER BY
	s.expected_value desc, total desc
`

type GetScenesWithColorRow struct {
	ID             int64
	Photograph     string
	Abbreviation   string
	Member         string
	Color          string
	Total          int64
	VocalMax       int64
	DanceMax       int64
	PerformanceMax int64
	ExpectedValue  sql.NullString
	SsrPlus        int64
	Bonds          int64
	Discography    int64
	Have           int64
}

func (q *Queries) GetScenesWithColor(ctx context.Context, db DBTX, name string) ([]GetScenesWithColorRow, error) {
	rows, err := db.QueryContext(ctx, getScenesWithColor, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetScenesWithColorRow
	for rows.Next() {
		var i GetScenesWithColorRow
		if err := rows.Scan(
			&i.ID,
			&i.Photograph,
			&i.Abbreviation,
			&i.Member,
			&i.Color,
			&i.Total,
			&i.VocalMax,
			&i.DanceMax,
			&i.PerformanceMax,
			&i.ExpectedValue,
			&i.SsrPlus,
			&i.Bonds,
			&i.Discography,
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

const registScene = `-- name: RegistScene :exec
;

INSERT INTO scenes (
	photograph_id,
	member_id,
	color_type_id,
	vocal_max,
	dance_max,
	performance_max,
	center_skill_name,
	expected_value,
	ssr_plus
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type RegistSceneParams struct {
	PhotographID    int64
	MemberID        int64
	ColorTypeID     int64
	VocalMax        int64
	DanceMax        int64
	PerformanceMax  int64
	CenterSkillName sql.NullString
	ExpectedValue   sql.NullString
	SsrPlus         int64
}

func (q *Queries) RegistScene(ctx context.Context, db DBTX, arg RegistSceneParams) error {
	_, err := db.ExecContext(ctx, registScene,
		arg.PhotographID,
		arg.MemberID,
		arg.ColorTypeID,
		arg.VocalMax,
		arg.DanceMax,
		arg.PerformanceMax,
		arg.CenterSkillName,
		arg.ExpectedValue,
		arg.SsrPlus,
	)
	return err
}
