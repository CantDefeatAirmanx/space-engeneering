package config_telegram

type TelegramConfigData struct {
	BotToken string `env:"botToken,required"`

	OrdersNotificationsChatId   int64 `env:"ordersNotificationsChatId,required"`
	OrdersNotificationsThreadId int   `env:"ordersNotificationsThreadId,required"`

	AssembliesNotificationsChatId   int64 `env:"assembliesNotificationsChatId,required"`
	AssembliesNotificationsThreadId int   `env:"assembliesNotificationsThreadId,required"`
}
