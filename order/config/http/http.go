package config_http

type HttpServerConfigData struct {
	Host              string `env:"host,required"`
	Port              int    `env:"port,required"`
	Timeout           int    `env:"timeout,required"`
	ReadHeaderTimeout int    `env:"readHeaderTimeout" envDefault:"5000"`
	ShutdownTimeout   int    `env:"shutdownTimeout" envDefault:"10000"`
}
