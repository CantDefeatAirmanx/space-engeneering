package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	Kafka() KafkaConfigInterface
	Logger() LoggerConfigInterface
	IsDev() bool
	Postgres() PostgresConfigInterface
	GRPC() GRPCConfigInterface
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

type PostgresConfigInterface interface {
	Port() int
	User() string
	Password() string
	DbName() string
	MigrationsDir() string
	Uri() string
}

type GRPCConfigInterface interface {
	Host() string
	Port() int
}
