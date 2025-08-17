package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	Logger() LoggerConfigInterface
	GRPC() GRPCConfigInterface
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}

type GRPCConfigInterface interface {
	Host() string
	Port() int
}
