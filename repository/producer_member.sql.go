// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: producer_member.sql

package repository

import (
	"context"
)

const getProducerMember = `-- name: GetProducerMember :many
SELECT
    pm.id,
    m.name,
    pm.bond_level_curent,
    pm.discography_disc_total
FROM
    producer_members pm
    JOIN members m ON pm.member_id = m.id
ORDER BY
    m.group_id, m.phase, m.first_name
`

type GetProducerMemberRow struct {
	ID                   int64
	Name                 string
	BondLevelCurent      int64
	DiscographyDiscTotal int64
}

func (q *Queries) GetProducerMember(ctx context.Context, db DBTX) ([]GetProducerMemberRow, error) {
	rows, err := db.QueryContext(ctx, getProducerMember)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProducerMemberRow
	for rows.Next() {
		var i GetProducerMemberRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BondLevelCurent,
			&i.DiscographyDiscTotal,
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

const registProducerMember = `-- name: RegistProducerMember :exec
;

INSERT INTO producer_members (
    producer_id,
    member_id,
    bond_level_curent,
    bond_level_collection_max,
    bond_level_scene_max,
    discography_disc_total,
    discography_disc_total_max
)
VALUES (?, ?, 0 ,0 ,0 ,0 ,0)
`

type RegistProducerMemberParams struct {
	ProducerID int64
	MemberID   int64
}

func (q *Queries) RegistProducerMember(ctx context.Context, db DBTX, arg RegistProducerMemberParams) error {
	_, err := db.ExecContext(ctx, registProducerMember, arg.ProducerID, arg.MemberID)
	return err
}

const updateProducerMember = `-- name: UpdateProducerMember :exec
UPDATE
    producer_members
SET
    bond_level_curent = ?,
    discography_disc_total = ?
WHERE
    id = ?
`

type UpdateProducerMemberParams struct {
	BondLevelCurent      int64
	DiscographyDiscTotal int64
	ID                   int64
}

func (q *Queries) UpdateProducerMember(ctx context.Context, db DBTX, arg UpdateProducerMemberParams) error {
	_, err := db.ExecContext(ctx, updateProducerMember, arg.BondLevelCurent, arg.DiscographyDiscTotal, arg.ID)
	return err
}
