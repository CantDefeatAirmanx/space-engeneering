package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var (
	_ ConfigInterface         = (*ConfigType)(nil)
	_ LoggerConfigInterface   = (*LoggerConfigType)(nil)
	_ KafkaConfigInterface    = (*KafkaConfigType)(nil)
	_ PostgresConfigInterface = (*PostgresConfigType)(nil)
	_ GRPCConfigInterface     = (*GRPCConfigType)(nil)
)

var Config *ConfigType

type ConfigType struct {
	configData ConfigData
	logger     LoggerConfigInterface
	kafka      KafkaConfigInterface
	postgres   PostgresConfigInterface
	grpc       GRPCConfigInterface
}

func (c *ConfigType) IsDev() bool {
	return isDev
}

func (c *ConfigType) GRPC() GRPCConfigInterface {
	return c.grpc
}

func (c *ConfigType) Postgres() PostgresConfigInterface {
	return c.postgres
}

func (c *ConfigType) Kafka() KafkaConfigInterface {
	return c.kafka
}

func (c *ConfigType) Logger() LoggerConfigInterface {
	return c.logger
}

func NewConfig(configData ConfigData) *ConfigType {
	return &ConfigType{
		configData: configData,

		logger: &LoggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},

		kafka: &KafkaConfigType{
			brokers:       configData.KafkaConfig.Brokers,
			orderTopic:    configData.KafkaConfig.OrderTopic,
			assemblyTopic: configData.KafkaConfig.AssemblyTopic,
		},

		postgres: &PostgresConfigType{
			dbName:        configData.PostgresConfig.DbName,
			uri:           configData.PostgresConfig.Uri,
			migrationsDir: configData.PostgresConfig.MigrationsDir,
			password:      configData.PostgresConfig.Password,
			port:          configData.PostgresConfig.Port,
			user:          configData.PostgresConfig.User,
		},

		grpc: &GRPCConfigType{
			host: configData.GRPCConfig.Host,
			port: configData.GRPCConfig.Port,
		},
	}
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

type LoggerConfigType struct {
	level   logger.Level
	encoder logger.EncoderType
}

func (l *LoggerConfigType) Encoder() logger.EncoderType {
	return l.encoder
}

func (l *LoggerConfigType) Level() logger.Level {
	return l.level
}

type KafkaConfigType struct {
	brokers       []string
	orderTopic    string
	assemblyTopic string
}

func (k *KafkaConfigType) Brokers() []string {
	return k.brokers
}

func (k *KafkaConfigType) AssemblyTopic() string {
	return k.assemblyTopic
}

func (k *KafkaConfigType) OrderTopic() string {
	return k.orderTopic
}

type PostgresConfigType struct {
	dbName        string
	uri           string
	migrationsDir string
	password      string
	port          int
	user          string
}

func (p *PostgresConfigType) DbName() string {
	return p.dbName
}

func (p *PostgresConfigType) MigrationsDir() string {
	return p.migrationsDir
}

func (p *PostgresConfigType) Password() string {
	return p.password
}

func (p *PostgresConfigType) Port() int {
	return p.port
}

func (p *PostgresConfigType) Uri() string {
	return p.uri
}

func (p *PostgresConfigType) User() string {
	return p.user
}
