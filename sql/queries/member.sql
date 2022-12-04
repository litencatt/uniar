-- name: GetMembers :many
SELECT
	g.name AS group_name,
	m.name,
	m.phase,
	m.graduated
FROM
	members m
	JOIN groups g ON m.group_id = g.id
ORDER BY
	g.id, m.phase, m.first_name asc
;

-- name: GetMemberList :many
SELECT
	m.id,
	m.name
FROM
	members m
WHERE
	group_id = ?
ORDER BY
	m.phase, m.first_name asc
;

-- name: GetAllMembers :many
SELECT
	*
FROM
	members m
ORDER BY
	m.group_id, m.phase, m.first_name asc
;