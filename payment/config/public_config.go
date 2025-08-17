package config

var (
	_ ConfigInterface = (*ConfigType)(nil)
	_ LoggerConfig    = (*LoggerConfigType)(nil)
	_ GRPCConfig      = (*GRPCConfigType)(nil)
)

var Config = &ConfigType{}

type ConfigType struct {
	configData ConfigData
	logger     LoggerConfig
	grpc       GRPCConfig
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,

		logger: &LoggerConfigType{
			level:  configData.LoggerConfig.Level,
			asJSON: configData.LoggerConfig.AsJSON,
		},

		grpc: &GRPCConfigType{
			host: configData.GRPCConfig.Host,
			port: configData.GRPCConfig.Port,
		},
	}
}

func (c *ConfigType) GRPC() GRPCConfig {
	return c.grpc
}

func (c *ConfigType) Logger() LoggerConfig {
	return c.logger
}

type LoggerConfigType struct {
	level  string
	asJSON bool
}

func (l *LoggerConfigType) AsJSON() bool {
	return l.asJSON
}

func (l *LoggerConfigType) Level() string {
	return l.level
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
