package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	Mongo() MongoConfigInterface
	GRPC() GRPCConfigInterface
	Logger() LoggerConfigInterface
	IsDev() bool
}

type MongoConfigInterface interface {
	URI() string
	Username() string
	Password() string
	DBName() string
	Port() int
	AuthSource() string
}

type GRPCConfigInterface interface {
	Host() string
	Port() int
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}
