package repository_notification_method_postgres

const (
	tableNotificationMethods = "notification_methods"

	columnNotificationMethodUUID         = "uuid"
	columnNotificationMethodProviderName = "provider_name"

	columnNotificationMethodCreatedAt = "created_at"
	columnNotificationMethodUpdatedAt = "updated_at"
)

const (
	tableUserToNotificationMethods = "user_to_notification_methods"

	columnUserToNotificationMethodUserUUID               = "user_uuid"
	columnUserToNotificationMethodNotificationMethodUUID = "notification_method_uuid"
	columnUserToNotificationMethodTarget                 = "target"
)

var (
	_ = columnNotificationMethodCreatedAt
	_ = columnNotificationMethodUpdatedAt
)
