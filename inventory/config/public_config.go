package config

import (
	"fmt"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
)

var (
	_ ConfigInterface           = (*ConfigType)(nil)
	_ MongoConfigInterface      = (*MongoConfigType)(nil)
	_ GRPCConfigInterface       = (*GRPCConfigType)(nil)
	_ LoggerConfigInterface     = (*LoggerConfigType)(nil)
	_ AuthClientConfigInterface = (*AuthClientConfigType)(nil)
)

var Config = &ConfigType{}

type ConfigType struct {
	configData ConfigData
	grpc       GRPCConfigInterface
	mongo      MongoConfigInterface
	logger     LoggerConfigInterface
	authClient AuthClientConfigInterface
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,

		grpc: &GRPCConfigType{
			host: configData.GRPCConfig.Host,
			port: configData.GRPCConfig.Port,
		},

		mongo: &MongoConfigType{
			host:       configData.MongoConfig.Host,
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

		authClient: &AuthClientConfigType{
			url: configData.AuthClient.Url,
		},
	}
}

func (c *ConfigType) AuthClient() AuthClientConfigInterface {
	return c.authClient
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
	authSource string
	dbName     string
	password   string
	host       string
	port       int
	username   string
	imageName  string
}

func (m *MongoConfigType) Host() string {
	return m.host
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
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=%s",
		m.username,
		m.password,
		m.host,
		m.port,
		m.dbName,
		m.authSource,
	)
}

type AuthClientConfigType struct {
	url string
}

func (a *AuthClientConfigType) Url() string {
	return a.url
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
