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

-- name: GetMemberList :many
SELECT id, name FROM members WHERE group_id = ?;