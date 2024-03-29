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