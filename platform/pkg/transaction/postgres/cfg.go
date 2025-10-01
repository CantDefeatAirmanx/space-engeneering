package platform_transaction_postgres

import (
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

type Config struct {
	RollbackMaxAttempts int
	Logger              platform_transaction.Logger
}

func NewConfig(opts ...ConfigOption) *Config {
	cfg := &Config{RollbackMaxAttempts: 3, Logger: logger.DefaultInfoLogger()}

	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

type ConfigOption func(cfg *Config)

func WithLogger(logger platform_transaction.Logger) ConfigOption {
	return func(cfg *Config) {
		cfg.Logger = logger
	}
}

func WithRollbackMaxAttempts(rollbackMaxAttempts int) ConfigOption {
	return func(cfg *Config) {
		cfg.RollbackMaxAttempts = rollbackMaxAttempts
	}
}
