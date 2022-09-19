-- name: GetGroup :many
SELECT id, name FROM groups;

-- name: GetGroupNameById :one
SELECT name FROM groups WHERE id = ?;