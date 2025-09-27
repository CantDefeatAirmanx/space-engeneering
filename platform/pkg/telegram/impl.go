package platform_telegram

import (
	"github.com/go-telegram/bot"
)

type TelegramClientImpl struct {
	tgBot *bot.Bot
}

func NewTelegramClient(
	botToken string,
	options ...bot.Option,
) (TelegramClient, error) {
	tgBot, err := bot.New(botToken, options...)
	if err != nil {
		return nil, err
	}

	return &TelegramClientImpl{
		tgBot: tgBot,
	}, nil
}
