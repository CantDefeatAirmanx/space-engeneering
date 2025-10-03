package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var Config *configType

var (
	_ ConfigInterface = (*configType)(nil)

	_ LoggerConfigInterface   = (*loggerConfigType)(nil)
	_ GRPCConfigInterface     = (*grpcConfigType)(nil)
	_ PostgresConfigInterface = (*postgresConfigType)(nil)
	_ RedisConfigInterface    = (*redisConfigType)(nil)
)

type configType struct {
	logger   LoggerConfigInterface
	grpc     GRPCConfigInterface
	postgres PostgresConfigInterface
	redis    RedisConfigInterface
}

func newConfig(configData configData) *configType {
	return &configType{
		logger: &loggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},

		grpc: &grpcConfigType{
			host: configData.GRPCConfig.Host,
			port: configData.GRPCConfig.Port,
		},

		postgres: &postgresConfigType{
			port:     configData.PostgresConfig.Port,
			user:     configData.PostgresConfig.User,
			password: configData.PostgresConfig.Password,
			dbName:   configData.PostgresConfig.DBName,
			uri:      configData.PostgresConfig.Uri,
		},

		redis: &redisConfigType{
			host:         configData.RedisConfig.Host,
			password:     configData.RedisConfig.Password,
			externalPort: configData.RedisConfig.ExternalPort,
		},
	}
}

func (c *configType) IsDev() bool {
	return isDev
}

func (c *configType) GRPC() GRPCConfigInterface {
	return c.grpc
}

func (c *configType) Logger() LoggerConfigInterface {
	return c.logger
}

func (c *configType) Postgres() PostgresConfigInterface {
	return c.postgres
}

func (c *configType) Redis() RedisConfigInterface {
	return c.redis
}

type grpcConfigType struct {
	host string
	port int
}

func (g *grpcConfigType) Host() string {
	return g.host
}

func (g *grpcConfigType) Port() int {
	return g.port
}

type postgresConfigType struct {
	dbName   string
	port     int
	user     string
	password string
	uri      string
}

func (p *postgresConfigType) DbName() string {
	return p.dbName
}

func (p *postgresConfigType) Password() string {
	return p.password
}

func (p *postgresConfigType) Port() int {
	return p.port
}

func (p *postgresConfigType) Uri() string {
	return p.uri
}

func (p *postgresConfigType) User() string {
	return p.user
}

type redisConfigType struct {
	host         string
	password     string
	externalPort int
}

func (r *redisConfigType) Host() string {
	return r.host
}

func (r *redisConfigType) ExternalPort() int {
	return r.externalPort
}

func (r *redisConfigType) Password() string {
	return r.password
}

type loggerConfigType struct {
	level   logger.Level
	encoder logger.EncoderType
}

func (l *loggerConfigType) Encoder() logger.EncoderType {
	return l.encoder
}

func (l *loggerConfigType) Level() logger.Level {
	return l.level
}
