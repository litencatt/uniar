-- name: GetGroup :many
SELECT * FROM `groups`;

-- name: GetMusicList :many
SELECT
	l.`name` AS live,
	m.name AS music,
	c.name AS TYPE,
	m.`length`,
	m.music_bonus AS bonus,
	m.master
FROM
	music m
	JOIN lives l ON m.live_id = l.id
	JOIN color_types c ON m.color_type_id = c.id
ORDER BY
	l.id
;


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
	s.ssr_plus
FROM
	scenes s
	JOIN photograph p ON s.photograph_id = p.id
	JOIN color_types c ON s.color_type_id = c.id
	JOIN members m ON s.member_id = m.id
-- WHERE
-- 	p.id = 1
ORDER BY
	c.id, s.expected_value desc, total desc
;

-- name: GetCollections :many
SELECT
	p. `name` AS photograph,
	m. `name` AS member,
	c.name AS color,
	s.ssr_plus,
	s.expected_value AS expected_value
FROM
	producer_scenes ps
	JOIN photograph p ON ps.photograph_id = p.id
	JOIN members m ON ps.member_id = m.id
	LEFT OUTER JOIN scenes s ON p.id = s.photograph_id
		AND m.id = s.member_id
	JOIN color_types c ON s.color_type_id = c.id
WHERE
	have = 1
	AND expected_value IS NOT NULL
ORDER BY
	s.expected_value DESC
;

-- name: GetMembers :many
SELECT
	g. `name` AS `group`,
	m.id AS member_id,
	m. `name`,
	m.phase,
	m.graduated
FROM
	members m
	JOIN `groups` g ON m.group_id = g.id
;