package config

import (
	"context"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_http "github.com/CantDefeatAirmanx/space-engeneering/order/config/http"
	config_inventory_client "github.com/CantDefeatAirmanx/space-engeneering/order/config/inventory_client"
	config_payment_client "github.com/CantDefeatAirmanx/space-engeneering/order/config/payment_client"
	config_postgres "github.com/CantDefeatAirmanx/space-engeneering/order/config/postgres"
)

type ConfigData struct {
	Postgres        config_postgres.PostgresConfigData                `envPrefix:"postgres__"`
	HttpServer      config_http.HttpServerConfigData                  `envPrefix:"httpServer__"`
	InventoryClient config_inventory_client.InventoryClientConfigData `envPrefix:"inventoryClient__"`
	PaymentClient   config_payment_client.PaymentClientConfigData     `envPrefix:"paymentClient__"`
	IsDev           bool
}

var (
	configData ConfigData
	isDev      = os.Getenv("GO_ENV") == "dev"
)

func LoadConfig(ctx context.Context, options ...LoadConfigOption) error {
	opts := LoadConfigOptions{
		EnvPath: ".env",
	}
	for _, option := range options {
		option(&opts)
	}

	if isDev {
		if err := godotenv.Load(opts.EnvPath); err != nil {
			return err
		}
	}

	if err := env.Parse(&configData); err != nil {
		return err
	}

	Config = NewConfig(configData)

	return nil
}

type LoadConfigOptions struct {
	EnvPath string
}

type LoadConfigOption func(*LoadConfigOptions)

func WithEnvPath(envPath string) LoadConfigOption {
	return func(options *LoadConfigOptions) {
		options.EnvPath = envPath
	}
}
