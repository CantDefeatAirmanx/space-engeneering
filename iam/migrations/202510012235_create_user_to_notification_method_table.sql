-- +goose Up
CREATE TABLE IF NOT EXISTS user_to_notification_methods (
    user_uuid UUID NOT NULL,
    notification_method_uuid UUID NOT NULL,
    
    PRIMARY KEY (user_uuid, notification_method_uuid),
    
    FOREIGN KEY (user_uuid) REFERENCES users(uuid) ON DELETE CASCADE,
    FOREIGN KEY (notification_method_uuid) REFERENCES notification_methods(uuid) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS user_to_notification_methods;
