package config

import (
	"os"

	config_grpc "github.com/CantDefeatAirmanx/space-engeneering/payment/config/grpc"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/payment/config/logger"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ConfigData struct {
	LoggerConfig config_logger.LoggerConfig `envPrefix:"logger__"`
	GRPCConfig   config_grpc.GRPCConfig     `envPrefix:"grpc__"`
}

var config ConfigData
var IS_DEV = os.Getenv("GO_ENV") == "dev"

func LoadConfig(opts ...LoadConfigOption) error {
	options := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, opt := range opts {
		opt(&options)
	}

	if IS_DEV {
		godotenv.Load(options.EnvPath)
	}

	if err := env.Parse(&config); err != nil {
		return err
	}

	Config = NewConfig(config)

	return nil
}

type LoadConfigOptions struct {
	EnvPath string
}
type LoadConfigOption func(o *LoadConfigOptions)

func WithEnvPath(path string) LoadConfigOption {
	return func(o *LoadConfigOptions) {
		o.EnvPath = path
	}
}
