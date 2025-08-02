package configs_order

import "time"

const (
	Port                = 8080
	Timeout             = 10 * time.Second
	ReadHeaderTimeout   = 5 * time.Second
	ShutdownTimeout     = 10 * time.Second
	PaymentServiceURL   = "localhost:50050"
	InventoryServiceURL = "localhost:50051"

	// ToDo: Add working with .env abstactions
	EnvPostgresDbURI = "DB_URI"
	EnvMigrationsDir = "MIGRATIONS_DIR"
)
