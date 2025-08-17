package config

type ConfigInterface interface {
	Logger() LoggerConfig
	GRPC() GRPCConfig
}

type LoggerConfig interface {
	Level() string
	AsJSON() bool
}

type GRPCConfig interface {
	Host() string
	Port() int
}
