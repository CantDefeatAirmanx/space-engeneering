package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

var (
	_ ConfigInterface                = (*ConfigImpl)(nil)
	_ HttpServerConfigInterface      = (*HttpServerConfigType)(nil)
	_ InventoryClientConfigInterface = (*InventoryClientConfigType)(nil)
	_ PaymentClientConfigInterface   = (*PaymentClientConfigType)(nil)
	_ PostgresConfigInterface        = (*PostgresConfigType)(nil)
	_ LoggerConfigInterface          = (*LoggerConfigType)(nil)
	_ KafkaConfigInterface           = (*KafkaConfigType)(nil)
	_ AuthClientConfigInterface      = (*AuthClientConfigType)(nil)
)

var Config = &ConfigImpl{}

type ConfigImpl struct {
	configData      ConfigData
	httpServer      HttpServerConfigInterface
	inventoryClient InventoryClientConfigInterface
	paymentClient   PaymentClientConfigInterface
	postgres        PostgresConfigInterface
	logger          LoggerConfigInterface
	kafka           KafkaConfigInterface
	authClient      AuthClientConfigInterface
}

func NewConfig(configData ConfigData) *ConfigImpl {
	return &ConfigImpl{
		configData: configData,

		httpServer: &HttpServerConfigType{
			host:              configData.HttpServer.Host,
			port:              configData.HttpServer.Port,
			timeout:           configData.HttpServer.Timeout,
			readHeaderTimeout: configData.HttpServer.ReadHeaderTimeout,
			shutdownTimeout:   configData.HttpServer.ShutdownTimeout,
		},

		inventoryClient: &InventoryClientConfigType{
			url: configData.InventoryClient.Url,
		},

		paymentClient: &PaymentClientConfigType{
			url: configData.PaymentClient.Url,
		},

		authClient: &AuthClientConfigType{
			url: configData.AuthClient.Url,
		},

		postgres: &PostgresConfigType{
			dbName:        configData.Postgres.DbName,
			uri:           configData.Postgres.Uri,
			migrationsDir: configData.Postgres.MigrationsDir,
			password:      configData.Postgres.Password,
			port:          configData.Postgres.Port,
			user:          configData.Postgres.User,
		},

		logger: &LoggerConfigType{
			level:   configData.LoggerConfig.Level,
			encoder: configData.LoggerConfig.Encoder,
		},

		kafka: &KafkaConfigType{
			brokers:    configData.KafkaConfig.Brokers,
			orderTopic: configData.KafkaConfig.OrderTopic,
		},
	}
}

func (c *ConfigImpl) Logger() LoggerConfigInterface {
	return c.logger
}

func (c *ConfigImpl) IsDev() bool {
	return isDev
}

func (c *ConfigImpl) HttpServer() HttpServerConfigInterface {
	return c.httpServer
}

func (c *ConfigImpl) InventoryClient() InventoryClientConfigInterface {
	return c.inventoryClient
}

func (c *ConfigImpl) PaymentClient() PaymentClientConfigInterface {
	return c.paymentClient
}

func (c *ConfigImpl) AuthClient() AuthClientConfigInterface {
	return c.authClient
}

func (c *ConfigImpl) Postgres() PostgresConfigInterface {
	return c.postgres
}

func (c *ConfigImpl) Kafka() KafkaConfigInterface {
	return c.kafka
}

type HttpServerConfigType struct {
	host              string
	port              int
	timeout           int
	readHeaderTimeout int
	shutdownTimeout   int
}

func (h *HttpServerConfigType) Host() string {
	return h.host
}

func (h *HttpServerConfigType) Port() int {
	return h.port
}

func (h *HttpServerConfigType) ReadHeaderTimeout() int {
	return h.readHeaderTimeout
}

func (h *HttpServerConfigType) ShutdownTimeout() int {
	return h.shutdownTimeout
}

func (h *HttpServerConfigType) Timeout() int {
	return h.timeout
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

type InventoryClientConfigType struct {
	url string
}

func (i *InventoryClientConfigType) Url() string {
	return i.url
}

type PaymentClientConfigType struct {
	url string
}

func (p *PaymentClientConfigType) Url() string {
	return p.url
}

type AuthClientConfigType struct {
	url string
}

func (a *AuthClientConfigType) Url() string {
	return a.url
}

type KafkaConfigType struct {
	brokers    []string
	orderTopic string
}

func (k *KafkaConfigType) Brokers() []string {
	return k.brokers
}

func (k *KafkaConfigType) OrderTopic() string {
	return k.orderTopic
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
