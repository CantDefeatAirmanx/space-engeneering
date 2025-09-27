package platform_telegram

import "context"

type TelegramClient interface {
	SendMessage(
		ctx context.Context,
		message string,
		chatID int,
		opts ...MessageOption,
	) (*MessageInfo, error)
}
