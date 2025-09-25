package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var Config ConfigInterface

var (
	_ ConfigInterface = (*configType)(nil)

	_ LoggerConfigInterface   = (*loggerConfigType)(nil)
	_ TelegramConfigInterface = (*telegramConfigType)(nil)
	_ KafkaConfigInterface    = (*kafkaConfigType)(nil)
)

type configType struct {
	configData ConfigData
	logger     LoggerConfigInterface
	telegram   TelegramConfigInterface
	kafka      KafkaConfigInterface
}

func (c *configType) IsDev() bool {
	return isDev
}

func (c *configType) Kafka() KafkaConfigInterface {
	return c.kafka
}

func (c *configType) Logger() LoggerConfigInterface {
	return c.logger
}

func (c *configType) Telegram() TelegramConfigInterface {
	return c.telegram
}

func newConfig(configData ConfigData) *configType {
	return &configType{
		configData: configData,

		logger: &loggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},

		telegram: &telegramConfigType{
			botToken: configData.TelegramConfig.BotToken,

			ordersNotificationsChatId:   configData.TelegramConfig.OrdersNotificationsChatId,
			ordersNotificationsThreadId: configData.TelegramConfig.OrdersNotificationsThreadId,

			assembliesNotificationsChatId:   configData.TelegramConfig.AssembliesNotificationsChatId,
			assembliesNotificationsThreadId: configData.TelegramConfig.AssembliesNotificationsThreadId,
		},

		kafka: &kafkaConfigType{
			brokers:       configData.KafkaConfig.Brokers,
			orderTopic:    configData.KafkaConfig.OrderTopic,
			assemblyTopic: configData.KafkaConfig.AssemblyTopic,
		},
	}
}

type loggerConfigType struct {
	level   logger.Level
	encoder logger.EncoderType
}

func (l *loggerConfigType) Encoder() logger.EncoderType {
	return l.encoder
}

func (l *loggerConfigType) Level() logger.Level {
	return l.level
}

type telegramConfigType struct {
	botToken                    string
	ordersNotificationsChatId   int64
	ordersNotificationsThreadId int

	assembliesNotificationsChatId   int64
	assembliesNotificationsThreadId int
}

func (t *telegramConfigType) AssembliesNotificationsChatId() int {
	return int(t.assembliesNotificationsChatId)
}

func (t *telegramConfigType) AssembliesNotificationsThreadId() int {
	return t.assembliesNotificationsThreadId
}

func (t *telegramConfigType) BotToken() string {
	return t.botToken
}

func (t *telegramConfigType) OrdersNotificationsChatId() int {
	return int(t.ordersNotificationsChatId)
}

func (t *telegramConfigType) OrdersNotificationsThreadId() int {
	return t.ordersNotificationsThreadId
}

type kafkaConfigType struct {
	brokers       []string
	orderTopic    string
	assemblyTopic string
}

func (k *kafkaConfigType) Brokers() []string {
	return k.brokers
}

func (k *kafkaConfigType) AssemblyTopic() string {
	return k.assemblyTopic
}

func (k *kafkaConfigType) OrderTopic() string {
	return k.orderTopic
}
