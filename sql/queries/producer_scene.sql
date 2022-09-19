-- name: GetProducerScenes :many
SELECT
    ps.id,
	c.name AS color,
	p.name AS photograph,
	m.name AS member,
	s.ssr_plus,
	ps.have
FROM
	producer_scenes ps
	JOIN photograph p ON ps.photograph_id = p.id
	JOIN members m ON ps.member_id = m.id
	JOIN scenes s ON ps.photograph_id = s.photograph_id
		AND ps.member_id = s.member_id
	JOIN color_types c ON s.color_type_id = c.id
ORDER BY
	ps.photograph_id,
	m.phase,
	m.first_name
;

-- name: UpdateProducerScene :exec
UPDATE producer_scenes SET have = ? WHERE id = ?;