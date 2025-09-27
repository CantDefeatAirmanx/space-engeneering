package config_grpc

type GRPCConfig struct {
	Host string `env:"host,required"`
	Port int    `env:"port,required"`
}
