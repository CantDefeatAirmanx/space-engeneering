package config

import (
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

type ConfigInterface interface {
	Kafka() KafkaConfigInterface
	Logger() LoggerConfigInterface
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}

type KafkaConfigInterface interface {
	Brokers() []string
	ExternalPort() int
	InternalPort() int
	ControllerPort() int
	UiPort() int
}
