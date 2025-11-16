-- name: SaveMessages :copyfrom
INSERT INTO message (id, chat_id, role, content)
VALUES (sqlc.arg('id'), sqlc.arg('chat_id'), sqlc.arg('role'), sqlc.arg('content'));
