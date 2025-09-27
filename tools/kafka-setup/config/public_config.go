package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var (
	_ ConfigInterface       = (*ConfigType)(nil)
	_ LoggerConfigInterface = (*LoggerConfigType)(nil)
	_ KafkaConfigInterface  = (*KafkaConfigType)(nil)
)

var Config = &ConfigType{}

type ConfigType struct {
	configData ConfigData
	logger     LoggerConfigInterface
	kafka      KafkaConfigInterface
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,
		logger: &LoggerConfigType{
			level:   configData.Logger.Level,
			encoder: configData.Logger.Encoder,
		},
		kafka: &KafkaConfigType{
			brokers:        configData.Kafka.Brokers,
			externalPort:   configData.Kafka.ExternalPort,
			internalPort:   configData.Kafka.InternalPort,
			controllerPort: configData.Kafka.ControllerPort,
			uiPort:         configData.Kafka.UiPort,
		},
	}
}

func (c *ConfigType) Logger() LoggerConfigInterface {
	return c.logger
}

type LoggerConfigType struct {
	level   logger.Level
	encoder logger.EncoderType
}

func (l *LoggerConfigType) Level() logger.Level {
	return l.level
}

func (l *LoggerConfigType) Encoder() logger.EncoderType {
	return l.encoder
}

func (c *ConfigType) Kafka() KafkaConfigInterface {
	return c.kafka
}

type KafkaConfigType struct {
	brokers        []string
	externalPort   int
	internalPort   int
	controllerPort int
	uiPort         int
}

func (k *KafkaConfigType) Brokers() []string {
	return k.brokers
}

func (k *KafkaConfigType) ExternalPort() int {
	return k.externalPort
}

func (k *KafkaConfigType) InternalPort() int {
	return k.internalPort
}

func (k *KafkaConfigType) ControllerPort() int {
	return k.controllerPort
}

func (k *KafkaConfigType) UiPort() int {
	return k.uiPort
}
