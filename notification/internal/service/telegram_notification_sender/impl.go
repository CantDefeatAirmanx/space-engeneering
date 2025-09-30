package telegram_notification_sender

import (
	"context"

	model_notification_sender "github.com/CantDefeatAirmanx/space-engeneering/notification/internal/model/notification_sender"
	platform_telegram "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/telegram"
)

var _ model_notification_sender.NotificationSender = (*TelegramNotificationSenderImpl)(nil)

type TelegramNotificationSenderImpl struct {
	telegramClient platform_telegram.TelegramClient
	chatID         int
	messageOptions platform_telegram.MessageOptions
}

func NewTelegramNotificationSender(
	telegramClient platform_telegram.TelegramClient,
	chatID int,
	opts ...platform_telegram.MessageOption,
) *TelegramNotificationSenderImpl {
	messageOptions := platform_telegram.MessageOptions{}
	for _, opt := range opts {
		opt(&messageOptions)
	}

	return &TelegramNotificationSenderImpl{
		telegramClient: telegramClient,
		chatID:         chatID,
		messageOptions: messageOptions,
	}
}

func (t *TelegramNotificationSenderImpl) SendNotification(ctx context.Context, notification string) error {
	_, err := t.telegramClient.SendMessage(
		ctx,
		notification,
		t.chatID,
		platform_telegram.WithThreadId(t.messageOptions.ThreadId),
	)
	return err
}
