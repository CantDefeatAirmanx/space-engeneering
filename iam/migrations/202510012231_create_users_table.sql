-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_login ON users USING hash (login);
CREATE INDEX idx_users_email ON users USING hash (email);
CREATE INDEX idx_users_password_hash ON users USING hash (password_hash);

-- +goose Down
DROP INDEX IF EXISTS idx_users_login;
DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_password_hash;

DROP TABLE IF EXISTS users;
