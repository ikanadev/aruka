-- name: GetChatMessages :many
SELECT
  sqlc.embed(m)
FROM message m
WHERE m.chat_id = sqlc.arg('chat_id')
ORDER BY m.created_at ASC;

-- name: GetChatProviderModel :one
SELECT
  sqlc.embed(c),
  sqlc.embed(m),
  sqlc.embed(p)
FROM chat c
INNER JOIN model m ON c.model_id = m.id
INNER JOIN provider p ON m.provider_id = p.id
WHERE c.id = sqlc.arg('chat_id');

-- name: GetAllChats :many
SELECT
  sqlc.embed(c),
  sqlc.embed(mdl),
  sqlc.embed(p)
FROM chat c
INNER JOIN model mdl ON c.model_id = mdl.id
INNER JOIN provider p ON mdl.provider_id = p.id
ORDER BY c.created_at DESC;