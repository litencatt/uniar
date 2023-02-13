// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: producer.sql

package repository

import (
	"context"
	"database/sql"
)

const getProducer = `-- name: GetProducer :one
SELECT
    id, provider_id, identity_id, display_name, created_at
FROM
    producers p
WHERE
    p.identity_id = ?
`

func (q *Queries) GetProducer(ctx context.Context, db DBTX, identityID string) (Producer, error) {
	row := db.QueryRowContext(ctx, getProducer, identityID)
	var i Producer
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.IdentityID,
		&i.DisplayName,
		&i.CreatedAt,
	)
	return i, err
}

const registProducer = `-- name: RegistProducer :exec
;

INSERT OR REPLACE INTO producers
    (provider_id, identity_id, display_name)
VALUES
    (?, ?, ?)
`

type RegistProducerParams struct {
	ProviderID  int64
	IdentityID  string
	DisplayName sql.NullString
}

func (q *Queries) RegistProducer(ctx context.Context, db DBTX, arg RegistProducerParams) error {
	_, err := db.ExecContext(ctx, registProducer, arg.ProviderID, arg.IdentityID, arg.DisplayName)
	return err
}

const updateProducer = `-- name: UpdateProducer :exec
UPDATE
    producers
SET
    display_name = ?
WHERE
    id = ?
`

type UpdateProducerParams struct {
	DisplayName sql.NullString
	ID          int64
}

func (q *Queries) UpdateProducer(ctx context.Context, db DBTX, arg UpdateProducerParams) error {
	_, err := db.ExecContext(ctx, updateProducer, arg.DisplayName, arg.ID)
	return err
}