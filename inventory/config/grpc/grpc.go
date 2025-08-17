package config_grpc

type GRPCConfigData struct {
	Host string `env:"grpc__host,required"`
	Port int    `env:"grpc__port,required"`
}
