-- name: GetMembers :many
SELECT
	g. `name` AS `group`,
	m. `name`,
	m.phase,
	m.graduated
FROM
	members m
	JOIN `groups` g ON m.group_id = g.id
ORDER BY
	g.id, m.phase, m.first_name asc
;

-- name: GetMemberList :many
SELECT id, name FROM members WHERE group_id = ?;