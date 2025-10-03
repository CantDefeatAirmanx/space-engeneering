package config_redis

type RedisConfigData struct {
	Host         string `env:"host,required"`
	Password     string `env:"password,required"`
	ExternalPort int    `env:"externalPort,required"`
}
