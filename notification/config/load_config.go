package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_kafka "github.com/CantDefeatAirmanx/space-engeneering/notification/config/kafka"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/notification/config/logger"
	config_telegram "github.com/CantDefeatAirmanx/space-engeneering/notification/config/telegram"
)

var (
	configData ConfigData
	isDev      = os.Getenv("GO_ENV") == "dev"
)

type ConfigData struct {
	LoggerConfig   config_logger.LoggerConfigData     `envPrefix:"logger__"`
	TelegramConfig config_telegram.TelegramConfigData `envPrefix:"telegram__"`
	KafkaConfig    config_kafka.KafkaConfigData       `envPrefix:"kafka__"`
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
