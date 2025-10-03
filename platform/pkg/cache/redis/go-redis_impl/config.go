package platform_redis_redisgo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CommonConfig struct {
	Username   string
	Password   string
	ClientName string
	DB         int

	PoolSize    int
	PoolTimeout time.Duration

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	MaxRetries int
	OnConnect  func(ctx context.Context, cn *redis.Conn) error

	ConnMaxIdleTime time.Duration
	ConnMaxLifetime time.Duration

	SingleNodeClient *redis.Client
}

type ClusterConfig struct {
	Addrs []string
	CommonConfig
}

type SingleConfig struct {
	Addr string
	CommonConfig
}

func NewSingleNodeConfig(options ...OptionFunc) *CommonConfig {
	cfg := &CommonConfig{
		Username:   "",
		Password:   "",
		ClientName: defaultClientName,
		DB:         0,

		PoolSize:    defaultPoolSize,
		PoolTimeout: defaultPoolTimeout,

		ReadTimeout:  defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,

		MaxRetries: defaultMaxRetries,
		OnConnect:  func(ctx context.Context, cn *redis.Conn) error { return nil },

		ConnMaxIdleTime: defaultConnMaxIdleTime,
		ConnMaxLifetime: defaultConnMaxLifetime,
	}
	for _, option := range options {
		option(cfg)
	}
	return cfg
}

type OptionFunc func(config *CommonConfig)

func WithUsername(username string) OptionFunc {
	return func(config *CommonConfig) {
		config.Username = username
	}
}

func WithPassword(password string) OptionFunc {
	return func(config *CommonConfig) {
		config.Password = password
	}
}

func WithClientName(clientName string) OptionFunc {
	return func(config *CommonConfig) {
		config.ClientName = clientName
	}
}

func WithDB(db int) OptionFunc {
	return func(config *CommonConfig) {
		config.DB = db
	}
}

func WithPoolSize(poolSize int) OptionFunc {
	return func(config *CommonConfig) {
		config.PoolSize = poolSize
	}
}

func WithPoolTimeout(poolTimeout time.Duration) OptionFunc {
	return func(config *CommonConfig) {
		config.PoolTimeout = poolTimeout
	}
}

func WithReadTimeout(readTimeout time.Duration) OptionFunc {
	return func(config *CommonConfig) {
		config.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) OptionFunc {
	return func(config *CommonConfig) {
		config.WriteTimeout = writeTimeout
	}
}

func WithMaxRetries(maxRetries int) OptionFunc {
	return func(config *CommonConfig) {
		config.MaxRetries = maxRetries
	}
}

func WithOnConnect(onConnect func(ctx context.Context, cn *redis.Conn) error) OptionFunc {
	return func(config *CommonConfig) {
		config.OnConnect = onConnect
	}
}

func WithConnMaxIdleTime(connMaxIdleTime time.Duration) OptionFunc {
	return func(config *CommonConfig) {
		config.ConnMaxIdleTime = connMaxIdleTime
	}
}

func WithConnMaxLifetime(connMaxLifetime time.Duration) OptionFunc {
	return func(config *CommonConfig) {
		config.ConnMaxLifetime = connMaxLifetime
	}
}

func WithSingleNodeClient(singleNodeClient *redis.Client) OptionFunc {
	return func(config *CommonConfig) {
		config.SingleNodeClient = singleNodeClient
	}
}
