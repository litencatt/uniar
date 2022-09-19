// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: music.sql

package repository

import (
	"context"
	"database/sql"
)

const getMusicList = `-- name: GetMusicList :many
SELECT
	l.name AS live,
	m.name AS music,
	c.name AS TYPE,
	m.length,
	m.music_bonus AS bonus,
	m.master
FROM
	music m
	JOIN lives l ON m.live_id = l.id
	JOIN color_types c ON m.color_type_id = c.id
ORDER BY
	l.id
`

type GetMusicListRow struct {
	Live   string
	Music  string
	TYPE   string
	Length int64
	Bonus  sql.NullInt64
	Master int64
}

func (q *Queries) GetMusicList(ctx context.Context, db DBTX) ([]GetMusicListRow, error) {
	rows, err := db.QueryContext(ctx, getMusicList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMusicListRow
	for rows.Next() {
		var i GetMusicListRow
		if err := rows.Scan(
			&i.Live,
			&i.Music,
			&i.TYPE,
			&i.Length,
			&i.Bonus,
			&i.Master,
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

const getMusicListWithColor = `-- name: GetMusicListWithColor :many
;

SELECT
	l.name AS live,
	m.name AS music,
	c.name AS TYPE,
	m.length,
	m.music_bonus AS bonus,
	m.master
FROM
	music m
	JOIN lives l ON m.live_id = l.id
	JOIN color_types c ON m.color_type_id = c.id
WHERE
	c.name = ?
ORDER BY
	l.id
`

type GetMusicListWithColorRow struct {
	Live   string
	Music  string
	TYPE   string
	Length int64
	Bonus  sql.NullInt64
	Master int64
}

func (q *Queries) GetMusicListWithColor(ctx context.Context, db DBTX, name string) ([]GetMusicListWithColorRow, error) {
	rows, err := db.QueryContext(ctx, getMusicListWithColor, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMusicListWithColorRow
	for rows.Next() {
		var i GetMusicListWithColorRow
		if err := rows.Scan(
			&i.Live,
			&i.Music,
			&i.TYPE,
			&i.Length,
			&i.Bonus,
			&i.Master,
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

const registMusic = `-- name: RegistMusic :exec
;

INSERT INTO music (
	name,
	normal,
	pro,
	master,
	length,
	color_type_id,
	live_id
) VALUES (?, ?, ?, ?, ?, ?, ?)
`

type RegistMusicParams struct {
	Name        string
	Normal      int64
	Pro         int64
	Master      int64
	Length      int64
	ColorTypeID int64
	LiveID      int64
}

func (q *Queries) RegistMusic(ctx context.Context, db DBTX, arg RegistMusicParams) error {
	_, err := db.ExecContext(ctx, registMusic,
		arg.Name,
		arg.Normal,
		arg.Pro,
		arg.Master,
		arg.Length,
		arg.ColorTypeID,
		arg.LiveID,
	)
	return err
}