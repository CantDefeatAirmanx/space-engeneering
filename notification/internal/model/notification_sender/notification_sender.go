package model_notification_sender

import (
	"context"
)

type NotificationSender interface {
	SendNotification(
		ctx context.Context,
		notification string,
	) error
}
