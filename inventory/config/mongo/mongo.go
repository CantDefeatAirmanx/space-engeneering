package config_mongo

type MongoConfigData struct {
	URI        string `env:"mongo__uri,required"`
	Username   string `env:"mongo__username,required"`
	Password   string `env:"mongo__password,required"`
	DBName     string `env:"mongo__dbName,required"`
	Port       int    `env:"mongo__port,required"`
	AuthSource string `env:"mongo__authSource,required"`
}
