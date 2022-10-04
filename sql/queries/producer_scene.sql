-- name: GetProducerScenes :many
SELECT
    ps.producer_id,
    ps.scene_id,
    c.name AS color,
    p.name AS photograph,
    m.name AS member,
    s.ssr_plus,
    ps.have
FROM
    producer_scenes ps
    JOIN scenes s ON ps.scene_id = s.id
    JOIN photograph p on s.photograph_id = p.id
    JOIN members m on s.member_id = m.id
    JOIN color_types c ON s.color_type_id = c.id
ORDER BY
    p.id,
    m.phase,
    m.first_name
;

-- name: RegistProducerScene :exec
INSERT OR IGNORE INTO producer_scenes (
	producer_id,
	scene_id
) VALUES (?, ?)
;

-- name: InsertOrUpdateProducerScene :exec
INSERT OR REPLACE INTO producer_scenes (
	producer_id,
	scene_id,
    have
) VALUES (?, ?, ?)
;