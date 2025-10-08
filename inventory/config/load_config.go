package config

import (
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"

	config_grpc "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/grpc"
	config_logger "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/logger"
	config_mongo "github.com/CantDefeatAirmanx/space-engeneering/inventory/config/mongo"
	config_auth_client "github.com/CantDefeatAirmanx/space-engeneering/order/config/auth_client"
)

type ConfigData struct {
	AuthClient   config_auth_client.AuthClientConfigData `envPrefix:"authClient__"`
	MongoConfig  config_mongo.MongoConfigData            `envPrefix:"mongo__"`
	GRPCConfig   config_grpc.GRPCConfigData              `envPrefix:"grpc__"`
	LoggerConfig config_logger.LoggerConfigData          `envPrefix:"logger__"`
}

const (
	goEnvKey = "GO_ENV"
)

var (
	configData ConfigData
	IsDev      = os.Getenv(goEnvKey) == "dev"
	IsTest     = os.Getenv(goEnvKey) == "test"
	IsProd     = os.Getenv(goEnvKey) == "prod" || os.Getenv(goEnvKey) == ""
)

func LoadConfig(opts ...LoadConfigOption) error {
	options := LoadConfigOptions{
		EnvPath: ".env",
		IsTest:  false,
	}
	for _, opt := range opts {
		opt(&options)
	}

	if IsDev || IsTest || options.IsTest {
		if err := godotenv.Load(options.EnvPath); err != nil {
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
	IsTest  bool
}
type LoadConfigOption func(o *LoadConfigOptions)

func WithEnvPath(path string) LoadConfigOption {
	return func(o *LoadConfigOptions) {
		o.EnvPath = path
	}
}

func WithIsTest(isTest bool) LoadConfigOption {
	return func(o *LoadConfigOptions) {
		o.IsTest = isTest
	}
}
