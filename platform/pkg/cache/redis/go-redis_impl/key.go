package platform_redis_redisgo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

var _ platform_redis.KeyCache = (*KeyCache)(nil)

type KeyCache struct {
	client *redis.Client
}

func NewKeyCache(client *redis.Client) platform_redis.KeyCache {
	return &KeyCache{
		client: client,
	}
}

func (k *KeyCache) Expire(
	ctx context.Context,
	key string,
	ttl time.Duration,
) platform_redis.RedisError {
	cmd := k.client.Expire(ctx, key, ttl)
	err := cmd.Err()
	if err != nil {
		return convertRedisError(err)
	}
	return nil
}
