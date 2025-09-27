package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_kafka "github.com/CantDefeatAirmanx/space-engeneering/assembly/config/kafka"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/assembly/config/logger"
	config_postgres "github.com/CantDefeatAirmanx/space-engeneering/assembly/config/postgres"
	config_grpc "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/grpc"
)

var (
	config ConfigData
	isDev  = os.Getenv("GO_ENV") == "dev"
)

type ConfigData struct {
	LoggerConfig   config_logger.LoggerConfigData     `envPrefix:"logger__"`
	KafkaConfig    config_kafka.KafkaConfigData       `envPrefix:"kafka__"`
	PostgresConfig config_postgres.PostgresConfigData `envPrefix:"postgres__"`
	GRPCConfig     config_grpc.GRPCConfigData         `envPrefix:"grpc__"`
}

func LoadConfig(opts ...LoadConfigOption) error {
	cfg := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	if isDev {
		if err := godotenv.Load(cfg.EnvPath); err != nil {
			return err
		}
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
