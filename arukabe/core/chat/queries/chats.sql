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
WHERE c.deleted_at IS NULL
ORDER BY c.created_at DESC
LIMIT sqlc.arg('limit')
OFFSET sqlc.arg('offset');

-- name: CountChats :one
SELECT COUNT(*) FROM chat WHERE deleted_at IS NULL;

-- name: UpdateChat :exec
UPDATE chat
SET
  title = COALESCE(sqlc.narg('title'), title),
  prompt = COALESCE(sqlc.narg('prompt'), prompt),
  pinned = COALESCE(sqlc.narg('pinned'), pinned),
  updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('chat_id');

-- name: DeleteChat :exec
UPDATE chat SET deleted_at = NOW() WHERE id = $1;