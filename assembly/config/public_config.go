package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var (
	_ ConfigInterface       = (*ConfigType)(nil)
	_ LoggerConfigInterface = (*LoggerConfigType)(nil)
	_ KafkaConfigInterface  = (*KafkaConfigType)(nil)
)

var Config *ConfigType

type ConfigType struct {
	configData ConfigData
	logger     LoggerConfigInterface
	kafka      KafkaConfigInterface
}

func (c *ConfigType) IsDev() bool {
	return isDev
}

func (c *ConfigType) Kafka() KafkaConfigInterface {
	return c.kafka
}

func (c *ConfigType) Logger() LoggerConfigInterface {
	return c.logger
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,

		logger: &LoggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},

		kafka: &KafkaConfigType{
			brokers:       configData.KafkaConfig.Brokers,
			orderTopic:    configData.KafkaConfig.OrderTopic,
			assemblyTopic: configData.KafkaConfig.AssemblyTopic,
		},
	}
}

type LoggerConfigType struct {
	level   logger.Level
	encoder logger.EncoderType
}

func (l *LoggerConfigType) Encoder() logger.EncoderType {
	return l.encoder
}

func (l *LoggerConfigType) Level() logger.Level {
	return l.level
}

type KafkaConfigType struct {
	brokers       []string
	orderTopic    string
	assemblyTopic string
}

func (k *KafkaConfigType) Brokers() []string {
	return k.brokers
}

func (k *KafkaConfigType) AssemblyTopic() string {
	return k.assemblyTopic
}

func (k *KafkaConfigType) OrderTopic() string {
	return k.orderTopic
}
