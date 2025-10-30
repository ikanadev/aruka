CREATE TABLE provider (
  id UUID PRIMARY KEY,
  name VARCHAR(255) NOT NULL UNIQUE
);


CREATE TYPE model_status AS ENUM ('ACTIVE', 'INACTIVE', 'DEPRECATED');

CREATE TABLE model (
  id UUID PRIMARY KEY,
  model_identifier VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  provider_id UUID NOT NULL REFERENCES provider(id) ON DELETE CASCADE,
  status model_status NOT NULL DEFAULT 'ACTIVE',
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE chat (
  id UUID PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  prompt TEXT NOT NULL,
  pinned BOOLEAN NOT NULL DEFAULT FALSE,
  model_id UUID NOT NULL REFERENCES model(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  archived_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);
CREATE INDEX idx_chat_title ON chat (title);
CREATE INDEX idx_chat_created_at ON chat (created_at DESC);


CREATE TYPE message_type AS ENUM ('TEXT');
CREATE TYPE message_role AS ENUM ('USER', 'ASSISTANT');

CREATE TABLE message (
  id UUID PRIMARY KEY,
  chat_id UUID NOT NULL REFERENCES chat(id) ON DELETE CASCADE,
  type message_type NOT NULL DEFAULT 'TEXT',
  role message_role NOT NULL DEFAULT 'USER',
  content JSONB NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_message_content ON message USING GIN (content);
CREATE INDEX idx_message_chat_created ON message (chat_id, created_at);
