-- name: GetAllScenes :many
SELECT
	s.photograph_id,
	s.member_id,
	s.ssr_plus
FROM
	scenes s
;

-- name: GetScenesWithColor :many
SELECT
	s.id,
	p.name AS photograph,
	p.abbreviation,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.performance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.performance_max,
	s.expected_value,
	s.ssr_plus,
	pm.bond_level_curent AS bonds,
	pm.discography_disc_total AS discography,
	case
		when ps.have = 1 then true
		when ps.have != 1 then false
		when ps.have is NULL then false
	end as ps_have
FROM
	scenes s
	JOIN photograph p ON s.photograph_id = p.id
	JOIN color_types c ON s.color_type_id = c.id
	JOIN members m ON s.member_id = m.id
	LEFT OUTER JOIN producer_members pm ON s.member_id = pm.member_id
	LEFT OUTER JOIN producer_scenes ps
		ON s.photograph_id = ps.photograph_id AND s.member_id = ps.member_id AND s.ssr_plus = ps.ssr_plus
WHERE
	c.name LIKE ?
	AND m.name LIKE ?
	AND p.name LIKE ?
ORDER BY
	s.expected_value desc, total desc
;

-- name: RegistScene :exec
INSERT INTO scenes (
	photograph_id,
	member_id,
	color_type_id,
	vocal_max,
	dance_max,
	performance_max,
	center_skill,
	expected_value,
	ssr_plus
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
;