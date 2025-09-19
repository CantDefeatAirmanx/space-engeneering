package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	Kafka() KafkaConfigInterface
	Logger() LoggerConfigInterface
	IsDev() bool
}

type KafkaConfigInterface interface {
	Brokers() []string
	AssemblyTopic() string
	OrderTopic() string
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}
