-- +goose Up
CREATE TABLE IF NOT EXISTS user_to_notification_methods (
    user_uuid UUID NOT NULL,
    notification_method_uuid UUID NOT NULL,
    target JSONB NOT NULL,

    PRIMARY KEY (user_uuid, notification_method_uuid),
 
    FOREIGN KEY (user_uuid) REFERENCES users(uuid) ON DELETE CASCADE,
    FOREIGN KEY (notification_method_uuid) REFERENCES notification_methods(uuid) ON DELETE CASCADE
);

COMMENT ON COLUMN user_to_notification_methods.target IS 
'JSONB структура зависит от provider. Telegram: {"chat_id": number, "thread_id"?: number}';

-- +goose Down
DROP TABLE IF EXISTS user_to_notification_methods;
