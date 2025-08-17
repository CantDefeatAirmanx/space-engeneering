package config

type ConfigInterface interface {
	IsDev() bool
	Postgres() PostgresConfigInterface
	HttpServer() HttpServerConfigInterface
	InventoryClient() InventoryClientConfigInterface
	PaymentClient() PaymentClientConfigInterface
}

type HttpServerConfigInterface interface {
	Host() string
	Port() int
}

type PostgresConfigInterface interface {
	Port() int
	User() string
	Password() string
	DbName() string
	Uri() string
	MigrationsDir() string
}

type InventoryClientConfigInterface interface {
	Url() string
}

type PaymentClientConfigInterface interface {
	Url() string
}
