package config_http

type HttpServerConfigData struct {
	Host string `env:"host,required"`
	Port int    `env:"port,required"`
}
