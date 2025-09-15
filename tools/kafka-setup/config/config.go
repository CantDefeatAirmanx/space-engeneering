package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

var Config *config

type config struct {
	Kafka KafkaConfig `envPrefix:"kafka__"`
}

func LoadConfig(options ...LoadConfigOption) error {
	opts := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, opt := range options {
		opt(&opts)
	}
	cfg := &config{}

	if err := godotenv.Load(opts.EnvPath); err != nil {
		return err
	}

	if err := env.Parse(cfg); err != nil {
		return err
	}

	Config = cfg
	return nil
}
