-- name: RegistPhotograph :exec
INSERT INTO photograph (name, group_id, photo_type) VALUES (?, ?, ?);

-- name: GetPhotographList :many
SELECT
	id, name
FROM
	photograph
WHERE group_id = ? AND photo_type = ?
ORDER BY group_id, id ASC
;