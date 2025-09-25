package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	IsDev() bool
	Logger() LoggerConfigInterface
	Telegram() TelegramConfigInterface
	Kafka() KafkaConfigInterface
}

type KafkaConfigInterface interface {
	Brokers() []string
	OrderTopic() string
	AssemblyTopic() string
}

type TelegramConfigInterface interface {
	BotToken() string

	OrdersNotificationsChatId() int
	OrdersNotificationsThreadId() int

	AssembliesNotificationsChatId() int
	AssembliesNotificationsThreadId() int
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}
