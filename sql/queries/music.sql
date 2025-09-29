-- name: GetMusicList :many
SELECT
	l.name AS live,
	m.name AS music,
	c.name AS TYPE,
	m.length,
	m.music_bonus AS bonus,
	m.master
FROM
	music m
	JOIN lives l ON m.live_id = l.id
	JOIN color_types c ON m.color_type_id = c.id
ORDER BY
	l.id
;

-- name: GetMusicListWithColor :many
SELECT
	l.name AS live,
	m.name AS music,
	c.name AS TYPE,
	m.length,
	m.music_bonus AS bonus,
	m.master
FROM
	music m
	JOIN lives l ON m.live_id = l.id
	JOIN color_types c ON m.color_type_id = c.id
WHERE
	c.name = ?
ORDER BY
	l.id
;

-- name: RegistMusic :exec
INSERT INTO music (
	name,
	normal,
	pro,
	master,
	length,
	color_type_id,
	live_id
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetMusicById :one
SELECT * FROM music WHERE id = ?;

-- name: GetMusicListAll :many
SELECT
	m.*,
	l.name AS live_name,
	c.name AS color_name
FROM music m
JOIN lives l ON m.live_id = l.id
JOIN color_types c ON m.color_type_id = c.id
ORDER BY m.id DESC;

-- name: UpdateMusic :exec
UPDATE music
SET name = ?, normal = ?, pro = ?, master = ?,
    length = ?, color_type_id = ?, live_id = ?, music_bonus = ?
WHERE id = ?;

-- name: DeleteMusic :exec
DELETE FROM music WHERE id = ?;

-- name: SearchMusicList :many
SELECT
	m.*,
	l.name AS live_name,
	c.name AS color_name
FROM music m
JOIN lives l ON m.live_id = l.id
JOIN color_types c ON m.color_type_id = c.id
WHERE
	(CASE WHEN ?1 != '' THEN m.name LIKE '%' || ?1 || '%' ELSE 1 END) AND
	(CASE WHEN ?2 != 0 THEN m.live_id = ?2 ELSE 1 END) AND
	(CASE WHEN ?3 != 0 THEN m.color_type_id = ?3 ELSE 1 END)
ORDER BY m.id DESC;