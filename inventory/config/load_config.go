package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_grpc "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/grpc"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/logger"
	config_mongo "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/mongo"
)

type ConfigData struct {
	MongoConfig  config_mongo.MongoConfigData   `envPrefix:"mongo__"`
	GRPCConfig   config_grpc.GRPCConfigData     `envPrefix:"grpc__"`
	LoggerConfig config_logger.LoggerConfigData `envPrefix:"logger__"`
}

var (
	configData ConfigData
	isDev      = os.Getenv("GO_ENV") == "dev"
)

func LoadConfig(opts ...LoadConfigOption) error {
	options := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, opt := range opts {
		opt(&options)
	}

	if isDev {
		if err := godotenv.Load(options.EnvPath); err != nil {
			return err
		}
	}

	if err := env.Parse(&configData); err != nil {
		return err
	}

	Config = NewConfig(configData)

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
