package repository_notification_method

import (
	"context"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

type NotificationMethodRepository interface {
	GetUserNotificationMethods(
		ctx context.Context,
		userUUID string,
	) ([]model_notification_method.NotificationMethod, error)

	SetUserNotificationMethods(
		ctx context.Context,
		userUUID string,
		notificationMethods []model_notification_method.NotificationMethod,
	) error

	CreateNotificationMethods(
		ctx context.Context,
		notificationMethods []model_notification_method.NotificationMethod,
	) error

	platform_transaction.WithExecutor[NotificationMethodRepository, platform_transaction.Executor]
}
