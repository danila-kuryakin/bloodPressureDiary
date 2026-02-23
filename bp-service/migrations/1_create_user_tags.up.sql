CREATE TABLE user_tags (
    id SERIAL PRIMARY KEY,
    user_id TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,
    is_active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE INDEX idx_user_tags_user_id
    ON user_tags(user_id);

CREATE INDEX idx_user_tags_user_active
    ON user_tags(user_id, is_active);