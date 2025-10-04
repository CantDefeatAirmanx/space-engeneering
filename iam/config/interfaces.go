package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	IsDev() bool
	Logger() LoggerConfigInterface
	GRPC() GRPCConfigInterface
	Postgres() PostgresConfigInterface
	Redis() RedisConfigInterface
	Auth() AuthConfigInterface
}

type GRPCConfigInterface interface {
	Host() string
	Port() int
}

type PostgresConfigInterface interface {
	Port() int
	User() string
	Password() string
	DbName() string
	Uri() string
}

type RedisConfigInterface interface {
	Host() string
	Password() string
	ExternalPort() int
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}

type AuthConfigInterface interface {
	SessionTTLHours() int
}
