package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var (
	_ ConfigInterface       = (*ConfigType)(nil)
	_ LoggerConfigInterface = (*LoggerConfigType)(nil)
	_ GRPCConfigInterface   = (*GRPCConfigType)(nil)
)

var Config = &ConfigType{}

type ConfigType struct {
	configData ConfigData
	logger     LoggerConfigInterface
	grpc       GRPCConfigInterface
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,

		logger: &LoggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},

		grpc: &GRPCConfigType{
			host: configData.GRPCConfig.Host,
			port: configData.GRPCConfig.Port,
		},
	}
}

func (c *ConfigType) GRPC() GRPCConfigInterface {
	return c.grpc
}

func (c *ConfigType) Logger() LoggerConfigInterface {
	return c.logger
}

type LoggerConfigType struct {
	level   logger.Level
	encoder logger.EncoderType
}

func (l *LoggerConfigType) Level() logger.Level {
	return l.level
}

func (l *LoggerConfigType) Encoder() logger.EncoderType {
	return l.encoder
}

type GRPCConfigType struct {
	host string
	port int
}

func (g *GRPCConfigType) Host() string {
	return g.host
}

func (g *GRPCConfigType) Port() int {
	return g.port
}
