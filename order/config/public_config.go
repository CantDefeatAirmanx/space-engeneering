package config

var (
	_ ConfigInterface                = (*ConfigImpl)(nil)
	_ HttpServerConfigInterface      = (*HttpServerConfigType)(nil)
	_ InventoryClientConfigInterface = (*InventoryClientConfigType)(nil)
	_ PaymentClientConfigInterface   = (*PaymentClientConfigType)(nil)
	_ PostgresConfigInterface        = (*PostgresConfigType)(nil)
)

var Config = &ConfigImpl{}

type ConfigImpl struct {
	configData      ConfigData
	httpServer      HttpServerConfigInterface
	inventoryClient InventoryClientConfigInterface
	paymentClient   PaymentClientConfigInterface
	postgres        PostgresConfigInterface
}

func NewConfig(configData ConfigData) *ConfigImpl {
	return &ConfigImpl{
		configData: configData,

		httpServer: &HttpServerConfigType{
			host: configData.HttpServer.Host,
			port: configData.HttpServer.Port,
		},

		inventoryClient: &InventoryClientConfigType{
			url: configData.InventoryClient.Url,
		},

		paymentClient: &PaymentClientConfigType{
			url: configData.PaymentClient.Url,
		},

		postgres: &PostgresConfigType{
			dbName:        configData.Postgres.DbName,
			uri:           configData.Postgres.Uri,
			migrationsDir: configData.Postgres.MigrationsDir,
			password:      configData.Postgres.Password,
			port:          configData.Postgres.Port,
			user:          configData.Postgres.User,
		},
	}
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

func (c *ConfigImpl) Postgres() PostgresConfigInterface {
	return c.postgres
}

type HttpServerConfigType struct {
	host string
	port int
}

func (h *HttpServerConfigType) Host() string {
	return h.host
}

func (h *HttpServerConfigType) Port() int {
	return h.port
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
