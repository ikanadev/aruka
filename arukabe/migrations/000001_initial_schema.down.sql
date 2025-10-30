DROP INDEX IF EXISTS idx_message_content;
DROP INDEX IF EXISTS idx_message_chat_created;
DROP TABLE IF EXISTS message;
DROP TYPE IF EXISTS message_role;
DROP TYPE IF EXISTS message_type;

DROP INDEX IF EXISTS idx_chat_title;
DROP INDEX IF EXISTS idx_chat_created_at;
DROP TABLE IF EXISTS chat;

DROP TABLE IF EXISTS model;
DROP TYPE IF EXISTS model_status;

DROP TABLE IF EXISTS provider;
