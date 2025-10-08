package config

import "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"

type ConfigInterface interface {
	IsDev() bool
	Postgres() PostgresConfigInterface
	HttpServer() HttpServerConfigInterface
	InventoryClient() InventoryClientConfigInterface
	PaymentClient() PaymentClientConfigInterface
	AuthClient() AuthClientConfigInterface
	Logger() LoggerConfigInterface
	Kafka() KafkaConfigInterface
}

type HttpServerConfigInterface interface {
	Host() string
	Port() int
	Timeout() int
	ReadHeaderTimeout() int
	ShutdownTimeout() int
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

type KafkaConfigInterface interface {
	Brokers() []string
	OrderTopic() string
}

type LoggerConfigInterface interface {
	Level() logger.Level
	Encoder() logger.EncoderType
}

type AuthClientConfigInterface interface {
	Url() string
}
