package config_logger

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type LoggerConfigData struct {
	Level   logger.Level       `env:"level" envDefault:"info"`
	Encoder logger.EncoderType `env:"encoder" envDefault:"json"`
}
