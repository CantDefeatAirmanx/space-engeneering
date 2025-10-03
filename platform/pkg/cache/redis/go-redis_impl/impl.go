package platform_redis_redisgo

import (
	"fmt"

	"github.com/redis/go-redis/v9"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

var _ platform_redis.RedisCache = (*SingleNodeImpl)(nil)

type SingleNodeImpl struct {
	client      *redis.Client
	stringCache platform_redis.StringCache
	setCache    platform_redis.SetCache
	hashCache   platform_redis.HashCache

	addr    string
	options []OptionFunc
}

func NewSingleNodeClient(
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
	cfg := NewSingleNodeConfig(options...)

	redisOptions := &redis.Options{
		Addr:       addr,
		Username:   cfg.Username,
		Password:   cfg.Password,
		ClientName: cfg.ClientName,
		DB:         cfg.DB,

		PoolSize:    cfg.PoolSize,
		PoolTimeout: cfg.PoolTimeout,

		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,

		MaxRetries: cfg.MaxRetries,
		OnConnect:  cfg.OnConnect,

		ConnMaxIdleTime: cfg.ConnMaxIdleTime,
		ConnMaxLifetime: cfg.ConnMaxLifetime,

		Network: "tcp",
	}
	client := redis.NewClient(redisOptions)

	stringCache := NewStringCache(client)
	setCache := NewSetCache(client)
	hashCache := NewHashCache(client)

	return &SingleNodeImpl{
		client:      client,
		stringCache: stringCache,
		setCache:    setCache,
		hashCache:   hashCache,

		addr:    addr,
		options: options,
	}, nil
}

func (s *SingleNodeImpl) String() platform_redis.StringCache {
	return s.stringCache
}

func (s *SingleNodeImpl) Set() platform_redis.SetCache {
	return s.setCache
}

func (s *SingleNodeImpl) Hash() platform_redis.HashCache {
	return s.hashCache
}
