-- name: GetProducerMember :many
SELECT
    m.name,
    pm.bond_level_curent,
    pm.discography_disc_total
FROM
    producer_members pm
    JOIN members m ON pm.member_id = m.id
ORDER BY
    m.group_id, m.phase, m.first_name
;

-- name: RegistProducerMember :exec
INSERT OR IGNORE INTO producer_members (
    producer_id,
    member_id,
    bond_level_curent,
    bond_level_collection_max,
    bond_level_scene_max,
    discography_disc_total,
    discography_disc_total_max
)
VALUES (?, ?, 0 ,0 ,0 ,0 ,0);


-- name: UpdateProducerMember :exec
UPDATE
    producer_members
SET
    bond_level_curent = ?,
    discography_disc_total = ?
WHERE
    producer_id = ? AND member_id = ?
;
