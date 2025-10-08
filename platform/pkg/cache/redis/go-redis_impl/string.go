package platform_redis_redisgo

import (
	"context"
	"errors"
	"time"

	redis "github.com/redis/go-redis/v9"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

var _ platform_redis.StringCache = (*StringCache)(nil)

type StringCache struct {
	client *redis.Client
}

func NewStringCache(client *redis.Client) platform_redis.StringCache {
	return &StringCache{
		client: client,
	}
}

func (s *StringCache) Get(
	ctx context.Context,
	key string,
) (string, platform_redis.RedisError) {
	cmd := s.client.Get(ctx, key)
	res, err := cmd.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", platform_redis.ErrNotFound
		}
		return res, convertRedisError(err)
	}

	return res, nil
}

func (s *StringCache) MultiGet(
	ctx context.Context,
	keys []string,
) ([]string, platform_redis.RedisError) {
	cmd := s.client.MGet(ctx, keys...)
	res, err := cmd.Result()
	if err != nil {
		return nil, convertRedisError(err)
	}

	stringsResult := make([]string, len(res))
	for idx := range res {
		if res[idx] == nil {
			stringsResult[idx] = ""
			continue
		}
		stringsResult[idx] = res[idx].(string)
	}

	return stringsResult, nil
}

func (s *StringCache) Set(
	ctx context.Context,
	key string,
	value string,
	ttl time.Duration,
) platform_redis.RedisError {
	cmd := s.client.Set(ctx, key, value, ttl)

	err := cmd.Err()
	if err != nil {
		return convertRedisError(err)
	}

	return nil
}

func (s *StringCache) MultiSet(
	ctx context.Context, valuesMap map[string]string,
) platform_redis.RedisError {
	cmd := s.client.MSet(ctx, valuesMap)

	err := cmd.Err()
	if err != nil {
		return convertRedisError(err)
	}

	return nil
}

func (s *StringCache) Delete(ctx context.Context, key string) platform_redis.RedisError {
	cmd := s.client.Del(ctx, key)

	err := cmd.Err()
	if err != nil {
		return convertRedisError(err)
	}

	return nil
}
