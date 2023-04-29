-- name: GetProducerScenesByGroupId :many
SELECT
    ps.producer_id,
    ps.photograph_id,
    ps.member_id,
    c.name AS color,
    p.name AS photograph,
    m.name AS member,
    s.ssr_plus,
    CASE
      WHEN ps.have IS NULL then false
      WHEN ps.have = 0 then false
      WHEN ps.have = 1 then true
    end as ps_have
FROM
    producer_scenes ps
    JOIN scenes s ON ps.photograph_id = s.photograph_id AND ps.member_id = s.member_id AND ps.ssr_plus = s.ssr_plus
    JOIN photograph p on ps.photograph_id = p.id
    JOIN members m on ps.member_id = m.id
    JOIN color_types c ON s.color_type_id = c.id
WHERE
    m.group_id = ?
ORDER BY
    p.id,
    m.phase,
    m.first_name
;

-- name: GetProducerScenes :many
SELECT
    ps.producer_id,
    ps.photograph_id,
    ps.member_id,
    c.name AS color,
    p.name AS photograph,
    m.name AS member,
    s.ssr_plus,
    ps.have
FROM
    producer_scenes ps
    JOIN scenes s ON ps.photograph_id = s.photograph_id AND ps.member_id = s.member_id AND ps.ssr_plus = s.ssr_plus
    JOIN photograph p on ps.photograph_id = p.id
    JOIN members m on ps.member_id = m.id
    JOIN color_types c ON s.color_type_id = c.id
WHERE
    p.name LIKE ?
    AND m.name LIKE ?
ORDER BY
    p.id,
    m.phase,
    m.first_name
;

-- name: RegistProducerScene :exec
INSERT OR REPLACE INTO producer_scenes (
	producer_id,
	photograph_id,
    member_id,
    ssr_plus,
    have
) VALUES (?, ?, ?, ?, ?)
;

-- name: UpdateProducerScene :exec
UPDATE
    producer_scenes
SET
    have = ?
WHERE
	producer_id = ?
	AND photograph_id = ?
    AND member_id = ?
;

-- name: InitProducerSceneAll :exec
UPDATE
    producer_scenes
SET
    have = 0
WHERE
	producer_id = ?
    AND member_id = ?
;