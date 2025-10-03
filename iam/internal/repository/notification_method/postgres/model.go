package repository_notification_method_postgres

import "github.com/jackc/pgx/v5/pgtype"

type NotificationMethod struct {
	NotificationMethodUUID         pgtype.UUID                    `db:"notification_method_uuid"`
	NotificationMethodProviderName NotificationMethodProviderName `db:"provider_name"`
	NotificationMethodTarget       string                         `db:"target"`

	CreatedAt pgtype.Timestamp `db:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at"`
}

type NotificationMethodProviderName string

const (
	NotificationMethodProviderNameTelegram NotificationMethodProviderName = "telegram"
)
