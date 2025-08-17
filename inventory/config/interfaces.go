package config

type ConfigInterface interface {
	Mongo() MongoConfigInterface
	GRPC() GRPCConfigInterface
	IsDev() bool
}

type MongoConfigInterface interface {
	URI() string
	Username() string
	Password() string
	DBName() string
	Port() int
	AuthSource() string
}

type GRPCConfigInterface interface {
	Host() string
	Port() int
}
