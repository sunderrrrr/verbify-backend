CREATE TABLE chats (
                       id          SERIAL PRIMARY KEY,  -- UUID
                       user_id     INTEGER NOT NULL,         -- ID пользователя
                       task_id     INTEGER NOT NULL,         -- ID задания
                       CONSTRAINT unique_user_task_chat UNIQUE (user_id, task_id)
);
CREATE TABLE messages (
                          id          SERIAL PRIMARY KEY,
                          chat_id     VARCHAR(36) NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
                          role        VARCHAR(20) NOT NULL,  -- "user" или "assistant"
                          content     TEXT NOT NULL,
);
CREATE INDEX idx_chat_id ON messages (chat_id);
CREATE INDEX idx_user_task ON chats (user_id, task_id);