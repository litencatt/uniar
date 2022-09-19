-- name: GetProducerOffice :one
SELECT office_bonus FROM producer_offices WHERE id = 1;

-- name: UpdateProducerOffice :exec
UPDATE producer_offices SET office_bonus = ? WHERE id = 1;