-- name: SelectSceneList :many
SELECT
	p.name AS scene,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.peformance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.peformance_max,
	s.skill_name as 期待値,
	s.ssr_plus
FROM
	scenes s
	JOIN photograph p ON s.photograph_id = p.id
	JOIN color_types c ON s.color_type_id = c.id
	JOIN members m ON s.member_id = m.id
ORDER BY
	c.id, s.skill_name desc, total desc;