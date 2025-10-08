package repository_user_postgres

const (
	tableUsers = "users"

	columnUserUUID         = "uuid"
	columnUserLogin        = "login"
	columnUserEmail        = "email"
	columnUserPasswordHash = "password_hash"

	columnUserCreatedAt = "created_at"
	columnUserUpdatedAt = "updated_at"

	uniqueConstraintUsersLogin = "users_login_key"
	uniqueConstraintUsersEmail = "users_email_key"
)

const (
	tableNotificationMethods = "notification_methods"

	columnNotificationMethodUUID         = "notification_method_uuid"
	columnNotificationMethodProviderName = "provider_name"
	columnNotificationMethodTarget       = "target"
)

const (
	tableUserToNotificationMethods = "user_to_notification_methods"

	columnUserToNotificationMethodUserUUID               = "user_uuid"
	columnUserToNotificationMethodNotificationMethodUUID = "notification_method_uuid"
)

var (
	_ = tableNotificationMethods
	_ = columnNotificationMethodUUID
	_ = columnNotificationMethodProviderName
	_ = columnNotificationMethodTarget

	_ = tableUserToNotificationMethods
	_ = columnUserToNotificationMethodUserUUID
	_ = columnUserToNotificationMethodNotificationMethodUUID
)
