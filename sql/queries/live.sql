-- name: RegistLive :exec
INSERT INTO lives (name, group_id) VALUES (?, ?);

-- name: GetLiveList :many
SELECT id, name FROM lives;