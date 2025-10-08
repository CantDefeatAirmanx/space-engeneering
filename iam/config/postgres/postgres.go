package config_postgres

type PostgresConfigData struct {
	Port     int    `env:"port,required"`
	User     string `env:"user,required"`
	Password string `env:"password,required"`
	DBName   string `env:"dbName,required"`
	Uri      string `env:"uri,required"`
}
