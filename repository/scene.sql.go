// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: scene.sql

package repository

import (
	"context"
	"database/sql"
)

const getScenes = `-- name: GetScenes :many
SELECT
	p.name AS photograph,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.peformance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.peformance_max,
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
	JOIN producer_members pm ON s.member_id = pm.member_id
	JOIN producer_scenes ps ON s.photograph_id = ps.photograph_id AND s.member_id = ps.member_id
ORDER BY
	s.expected_value desc, total desc
`

type GetScenesRow struct {
	Photograph    string
	Member        string
	Color         string
	Total         int32
	VocalMax      int32
	DanceMax      int32
	PeformanceMax int32
	ExpectedValue sql.NullString
	SsrPlus       bool
	Bonds         int32
	Discography   int32
	Have          sql.NullBool
}

func (q *Queries) GetScenes(ctx context.Context, db DBTX) ([]GetScenesRow, error) {
	rows, err := db.QueryContext(ctx, getScenes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetScenesRow
	for rows.Next() {
		var i GetScenesRow
		if err := rows.Scan(
			&i.Photograph,
			&i.Member,
			&i.Color,
			&i.Total,
			&i.VocalMax,
			&i.DanceMax,
			&i.PeformanceMax,
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

const getScenesWithColor = `-- name: GetScenesWithColor :many
SELECT
	p.name AS photograph,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.peformance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.peformance_max,
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
	JOIN producer_members pm ON s.member_id = pm.member_id
	JOIN producer_scenes ps ON s.photograph_id = ps.photograph_id AND s.member_id = ps.member_id
WHERE
	c.name = ?
ORDER BY
	s.expected_value desc, total desc
`

type GetScenesWithColorRow struct {
	Photograph    string
	Member        string
	Color         string
	Total         int32
	VocalMax      int32
	DanceMax      int32
	PeformanceMax int32
	ExpectedValue sql.NullString
	SsrPlus       bool
	Bonds         int32
	Discography   int32
	Have          sql.NullBool
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
			&i.Photograph,
			&i.Member,
			&i.Color,
			&i.Total,
			&i.VocalMax,
			&i.DanceMax,
			&i.PeformanceMax,
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
INSERT INTO scenes (
	photograph_id,
	member_id,
	color_type_id,
	vocal_max,
	dance_max,
	peformance_max,
	center_skill_name,
	expected_value,
	ssr_plus
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type RegistSceneParams struct {
	PhotographID    int32
	MemberID        int32
	ColorTypeID     int32
	VocalMax        int32
	DanceMax        int32
	PeformanceMax   int32
	CenterSkillName sql.NullString
	ExpectedValue   sql.NullString
	SsrPlus         bool
}

func (q *Queries) RegistScene(ctx context.Context, db DBTX, arg RegistSceneParams) error {
	_, err := db.ExecContext(ctx, registScene,
		arg.PhotographID,
		arg.MemberID,
		arg.ColorTypeID,
		arg.VocalMax,
		arg.DanceMax,
		arg.PeformanceMax,
		arg.CenterSkillName,
		arg.ExpectedValue,
		arg.SsrPlus,
	)
	return err
}
