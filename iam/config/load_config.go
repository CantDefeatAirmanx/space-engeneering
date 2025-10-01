package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_grpc "github.com/CantDefeatAirmanx/space-engeneering/iam/config/grpc"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/iam/config/logger"
	config_postgres "github.com/CantDefeatAirmanx/space-engeneering/iam/config/postgres"
	config_redis "github.com/CantDefeatAirmanx/space-engeneering/iam/config/redis"
)

type configData struct {
	LoggerConfig   config_logger.LoggerConfigData     `envPrefix:"logger__"`
	GRPCConfig     config_grpc.GRPCConfigData         `envPrefix:"grpc__"`
	PostgresConfig config_postgres.PostgresConfigData `envPrefix:"postgres__"`
	RedisConfig    config_redis.RedisConfigData       `envPrefix:"redis__"`
}

var isDev = os.Getenv("GO_ENV") == "dev"

func LoadConfig(opts ...LoadConfigOption) error {
	cfg := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	configData := configData{}
	if isDev {
		if err := godotenv.Load(cfg.EnvPath); err != nil {
			return err
		}
	}

	if err := env.Parse(&configData); err != nil {
		return err
	}

	Config = newConfig(configData)

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
