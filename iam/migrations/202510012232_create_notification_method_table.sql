-- +goose Up
CREATE TABLE IF NOT EXISTS notification_methods (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    provider_name VARCHAR(255) NOT NULL,
    target VARCHAR(255) NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_notification_methods_provider_name ON notification_methods USING hash (provider_name);

-- +goose Down
DROP INDEX IF EXISTS idx_notification_methods_provider_name;

DROP TABLE IF EXISTS notification_methods;
