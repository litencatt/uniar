-- name: GetScenes :many
SELECT
	p.name AS photograph,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.peformance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.peformance_max,
	s.expected_value,
	s.ssr_plus,
	pm.bond_level_curent AS bonds,
	pm.discography_disc_total AS discography
FROM
	scenes s
	JOIN photograph p ON s.photograph_id = p.id
	JOIN color_types c ON s.color_type_id = c.id
	JOIN members m ON s.member_id = m.id
	JOIN producer_members pm ON s.member_id = pm.member_id
ORDER BY
	s.expected_value desc, total desc
;

-- name: GetScenesWithColor :many
SELECT
	p.name AS photograph,
	m.name AS member,
	c.name AS color,
	s.vocal_max + s.dance_max + s.peformance_max + 430 AS total,
	s.vocal_max,
	s.dance_max,
	s.peformance_max,
	s.expected_value,
	s.ssr_plus,
	pm.bond_level_curent AS bonds,
	pm.discography_disc_total AS discography
FROM
	scenes s
	JOIN photograph p ON s.photograph_id = p.id
	JOIN color_types c ON s.color_type_id = c.id
	JOIN members m ON s.member_id = m.id
	JOIN producer_members pm ON s.member_id = pm.member_id
WHERE
	c.name = ?
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
	peformance_max,
	center_skill_name,
	expected_value,
	ssr_plus
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
;