package config

var _ ConfigInterface = (*ConfigType)(nil)
var _ MongoConfigInterface = (*MongoConfigType)(nil)
var _ GRPCConfigInterface = (*GRPCConfigType)(nil)

var Config = &ConfigType{}

type ConfigType struct {
	configData ConfigData
	grpc       GRPCConfigInterface
	mongo      MongoConfigInterface
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
		},
	}
}

func (c *ConfigType) GRPC() GRPCConfigInterface {
	return c.grpc
}

func (c *ConfigType) IsDev() bool {
	return isDev
}

func (c *ConfigType) Mongo() MongoConfigInterface {
	return c.mongo
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
