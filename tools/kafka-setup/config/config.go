package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_kafka "github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/config/kafka"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/tools/kafka-setup/config/logger"
)

type ConfigData struct {
	Kafka  config_kafka.KafkaConfig   `envPrefix:"kafka__"`
	Logger config_logger.LoggerConfig `envPrefix:"logger__"`
}

func LoadConfig(options ...LoadConfigOption) error {
	opts := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, opt := range options {
		opt(&opts)
	}
	cfg := &ConfigData{}

	if err := godotenv.Load(opts.EnvPath); err != nil {
		return err
	}

	if err := env.Parse(cfg); err != nil {
		return err
	}

	Config = NewConfig(*cfg)
	return nil
}
