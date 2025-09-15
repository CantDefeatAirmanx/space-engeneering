package config_mongo

type MongoConfigData struct {
	Username   string `env:"username,required"`
	Host       string `env:"host,required"`
	Password   string `env:"password,required"`
	DBName     string `env:"dbName,required"`
	Port       int    `env:"port,required"`
	AuthSource string `env:"authSource,required"`
	ImageName  string `env:"imageName" envDefault:"mongo:7.0.5"`
}
