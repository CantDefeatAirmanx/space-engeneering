package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var (
	_ ConfigInterface       = (*ConfigType)(nil)
	_ MongoConfigInterface  = (*MongoConfigType)(nil)
	_ GRPCConfigInterface   = (*GRPCConfigType)(nil)
	_ LoggerConfigInterface = (*LoggerConfigType)(nil)
)

var Config = &ConfigType{}

type ConfigType struct {
	configData ConfigData
	grpc       GRPCConfigInterface
	mongo      MongoConfigInterface
	logger     LoggerConfigInterface
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,

		grpc: &GRPCConfigType{
			host: configData.GRPCConfig.Host,
			port: configData.GRPCConfig.Port,
		},

		mongo: &MongoConfigType{
			uri:        configData.MongoConfig.URI,
			authSource: configData.MongoConfig.AuthSource,
			dbName:     configData.MongoConfig.DBName,
			password:   configData.MongoConfig.Password,
			port:       configData.MongoConfig.Port,
			username:   configData.MongoConfig.Username,
			imageName:  configData.MongoConfig.ImageName,
		},

		logger: &LoggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},
	}
}

func (c *ConfigType) GRPC() GRPCConfigInterface {
	return c.grpc
}

func (c *ConfigType) IsDev() bool {
	return IsDev
}

func (c *ConfigType) Mongo() MongoConfigInterface {
	return c.mongo
}

func (c *ConfigType) Logger() LoggerConfigInterface {
	return c.logger
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

type MongoConfigType struct {
	uri        string
	authSource string
	dbName     string
	password   string
	port       int
	username   string
	imageName  string
}

func (m *MongoConfigType) ImageName() string {
	return m.imageName
}

func (m *MongoConfigType) AuthSource() string {
	return m.authSource
}

func (m *MongoConfigType) DBName() string {
	return m.dbName
}

func (m *MongoConfigType) Password() string {
	return m.password
}

func (m *MongoConfigType) Port() int {
	return m.port
}

func (m *MongoConfigType) Username() string {
	return m.username
}

func (m *MongoConfigType) URI() string {
	return m.uri
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
