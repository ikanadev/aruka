-- name: ListProviders :many
SELECT
  p.id as id,
  p.name as name,
  m.id as model_id,
  m.model_identifier as model_identifier,
  m.name as model_name,
  m.status as model_status,
  m.created_at as model_created_at
FROM provider p
INNER JOIN model m ON p.id = m.provider_id
WHERE m.status = sqlc.arg('status')
ORDER BY m.created_at DESC;

-- name: GetProviders :many
SELECT id, name FROM provider;

-- name: GetProviderModels :many
SELECT p.id as id, p.name as name, m.name as model_name, m.model_identifier as model_identifier
FROM provider p
INNER JOIN model m ON p.id = m.provider_id;

-- name: SaveProviders :copyfrom
INSERT INTO provider (id, name)
VALUES (sqlc.arg('id'), sqlc.arg('name'));

-- name: SaveModels :copyfrom
INSERT INTO model (id, provider_id, model_identifier, name, created_at)
VALUES (sqlc.arg('id'), sqlc.arg('provider_id'), sqlc.arg('model_identifier'), sqlc.arg('name'), sqlc.arg('created_at'));
