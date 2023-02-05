-- name: GetProducer :one
SELECT
    *
FROM
    producers p
WHERE
    p.identity_id = ?
;

-- name: RegistProducer :exec
INSERT OR REPLACE INTO producers
    (provider_id, identity_id, display_name)
VALUES
    (?, ?, ?);

-- name: UpdateProducer :exec
UPDATE
    producers
SET
    display_name = ?
WHERE
    id = ?
;