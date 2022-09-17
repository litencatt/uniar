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

-- name: GetCollectionsWithColor :many
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
	AND c.name = ?
ORDER BY
	s.expected_value DESC
;
