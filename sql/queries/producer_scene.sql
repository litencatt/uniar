-- name: GetProducerScenes :many
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
    JOIN scenes s ON ps.photograph_id = s.photograph_id AND ps.member_id = s.member_id
    JOIN photograph p on ps.photograph_id = p.id
    JOIN members m on ps.member_id = m.id
    JOIN color_types c ON s.color_type_id = c.id
ORDER BY
    p.id,
    m.phase,
    m.first_name
;

-- name: RegistProducerScene :exec
INSERT OR IGNORE INTO producer_scenes (
	producer_id,
	photograph_id,
    member_id
) VALUES (?, ?, ?)
;

-- name: InsertOrUpdateProducerScene :exec
INSERT OR REPLACE INTO producer_scenes (
	producer_id,
	photograph_id,
    member_id,
    have
) VALUES (?, ?, ?, ?)
;