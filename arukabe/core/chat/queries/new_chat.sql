-- name: CreateChat :one
WITH inserted_chat AS (
  INSERT INTO chat (id, title, prompt, model_id)
  VALUES (sqlc.arg('id'), sqlc.arg('title'), sqlc.arg('prompt'), sqlc.arg('model_id'))
  RETURNING *
)
SELECT ic.*, sqlc.embed(m), sqlc.embed(p)
FROM inserted_chat ic
JOIN model m
ON ic.model_id = m.id
JOIN provider p
ON m.provider_id = p.id;
