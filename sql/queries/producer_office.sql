-- name: GetProducerOffice :one
SELECT * FROM producer_offices WHERE producer_id = ?;

-- name: UpdateProducerOffice :exec
UPDATE producer_offices SET office_bonus = ? WHERE producer_id = ?;

-- name: RegistProducerOffice :exec
INSERT INTO producer_offices (
    producer_id,
    office_bonus
)
VALUES (?, 0);