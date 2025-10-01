package config_grpc

type GRPCConfigData struct {
	Host string `env:"host,required"`
	Port int    `env:"port,required"`
}
