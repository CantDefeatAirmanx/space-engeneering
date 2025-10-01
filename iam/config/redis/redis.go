package config_redis

type RedisConfigData struct {
	Password     string `env:"password,required"`
	ExternalPort int    `env:"externalPort,required"`
}
