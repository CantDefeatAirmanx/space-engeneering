package platform_redis_redisgo

import (
	"context"
	"fmt"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
	"github.com/redis/go-redis/v9"
)

var _ platform_redis.RedisCache = (*SingleNodeImpl)(nil)

type SingleNodeImpl struct {
	client *redis.Client
}

func NewSignleNodeClient(
	addr string,
	options ...OptionFunc,
) (impl *SingleNodeImpl, err error) {
	defer func() {
		if r := recover(); r != nil {
			impl = nil
			err = fmt.Errorf("failed to create redis client: %v", r)
		}
	}()

	if addr == "" {
		return nil, fmt.Errorf("%w: addr is required", platform_redis.ErrConfigError)
	}

	client := redis.NewClient(&redis.Options{
		Addr:       addr,
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

		Network: "tcp",
	})

	return &SingleNodeImpl{
		client,
	}, nil
}
