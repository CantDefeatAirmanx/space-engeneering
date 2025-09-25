package platform_telegram

import (
	"context"

	"github.com/go-telegram/bot"
)

func (c *TelegramClientImpl) SendMessage(
	ctx context.Context,
	message string,
	chatID int,
	opts ...MessageOption,
) (*MessageInfo, error) {
	messOptions := MessageOptions{}
	for _, opt := range opts {
		opt(&messOptions)
	}

	params := bot.SendMessageParams{
		ChatID: chatID,
		Text:   message,
	}
	if messOptions.ThreadId != 0 {
		params.MessageThreadID = messOptions.ThreadId
	}

	result, err := c.tgBot.SendMessage(ctx, &params)
	if err != nil {
		return nil, err
	}

	return &MessageInfo{
		MessageID: result.ID,
		ChatID:    int(result.Chat.ID),
		ThreadID:  result.MessageThreadID,
	}, nil
}
