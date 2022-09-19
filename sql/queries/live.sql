-- name: RegistLive :exec
INSERT INTO lives (name) VALUES (?);

-- name: GetLiveList :many
SELECT id, name FROM lives;