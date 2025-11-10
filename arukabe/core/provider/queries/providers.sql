-- name: ListProviders :many
SELECT
  sqlc.embed(p),
  sqlc.embed(m)
FROM provider p
INNER JOIN model m ON p.id = m.provider_id
WHERE m.status = sqlc.arg('status')
ORDER BY m.created_at DESC;

-- name: AllProviders :many
SELECT id, name FROM provider;

-- name: AllProviderModels :many
SELECT
  sqlc.embed(m),
  sqlc.embed(p)
FROM provider p
INNER JOIN model m ON p.id = m.provider_id;

-- name: SaveProviders :copyfrom
INSERT INTO provider (id, name)
VALUES (sqlc.arg('id'), sqlc.arg('name'));

-- name: SaveModels :copyfrom
INSERT INTO model (id, provider_id, model_identifier, name, created_at)
VALUES (sqlc.arg('id'), sqlc.arg('provider_id'), sqlc.arg('model_identifier'), sqlc.arg('name'), sqlc.arg('created_at'));
