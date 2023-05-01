-- name: RegistPhotograph :exec
INSERT INTO photograph (name, group_id, photo_type) VALUES (?, ?, ?);

-- name: GetPhotographListAll :many
SELECT
    name,
    photo_type,
    released_at
FROM
    photograph
ORDER BY released_at ASC
;

-- name: GetPhotographList :many
SELECT
    id,
    name
FROM
    photograph
WHERE
    group_id = ?
    AND photo_type = ?
ORDER BY
    group_id, id ASC
;

-- name: GetPhotographByGroupId :many
SELECT
    id,
    name
FROM
    photograph
WHERE
    group_id = ?
ORDER BY
	name_for_order ASC
;

-- name: GetPhotographListByPhotoType :many
SELECT
    id,
    name
FROM
    photograph
WHERE
    photo_type = ?
ORDER BY
    group_id, id ASC
;

-- name: GetSsrPlusReleasedPhotographList :many
SELECT
	photograph_id
FROM
	scenes
WHERE
	ssr_plus = 1
GROUP BY
	photograph_id
;