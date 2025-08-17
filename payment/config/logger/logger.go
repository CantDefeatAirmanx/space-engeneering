package config_logger

type LoggerConfig struct {
	Level  string `env:"level,required"`
	AsJSON bool   `env:"asJson,required"`
}
