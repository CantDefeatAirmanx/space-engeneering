package config_postgres

type PostgresConfigData struct {
	User          string `env:"user,required"`
	Password      string `env:"password,required"`
	DbName        string `env:"dbName,required"`
	Port          int    `env:"port,required"`
	Uri           string `env:"uri,required"`
	MigrationsDir string `env:"migrationsDir,required"`
}
