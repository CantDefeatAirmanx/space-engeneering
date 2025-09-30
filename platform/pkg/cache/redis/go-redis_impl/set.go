package platform_redis_redisgo

import (
	"context"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
)

var _ platform_redis.SetCache = (*SetCache)(nil)

type SetCache struct {
	client *redis.Client
}

func NewSetCache(client *redis.Client) platform_redis.SetCache {
	return &SetCache{
		client: client,
	}
}

func (s *SetCache) SAdd(
	ctx context.Context,
	key string,
	values ...string,
) (int64, platform_redis.RedisError) {
	cmd := s.client.SAdd(
		ctx,
		key,
		lo.Map(values, func(value string, _ int) interface{} {
			return value
		}),
	)
	res, err := cmd.Result()
	if err != nil {
		return 0, convertRedisError(err)
	}
	return res, nil
}

func (s *SetCache) SCard(
	ctx context.Context,
	key string,
) (int64, platform_redis.RedisError) {
	cmd := s.client.SCard(ctx, key)
	res, err := cmd.Result()
	if err != nil {
		return 0, convertRedisError(err)
	}
	return res, nil
}

func (s *SetCache) SRem(
	ctx context.Context,
	key string,
	values ...string,
) (int64, platform_redis.RedisError) {
	cmd := s.client.SRem(
		ctx,
		key,
		lo.Map(values, func(value string, _ int) interface{} {
			return value
		}),
	)
	res, err := cmd.Result()
	if err != nil {
		return 0, convertRedisError(err)
	}
	return res, nil
}

func (s *SetCache) SIsMember(
	ctx context.Context,
	key string,
	value string,
) (bool, platform_redis.RedisError) {
	cmd := s.client.SIsMember(ctx, key, value)
	res, err := cmd.Result()
	if err != nil {
		return false, convertRedisError(err)
	}
	return res, nil
}

func (s *SetCache) SScan(
	ctx context.Context,
	key string,
	cursor uint64,
	match string,
	count int64,
) ([]string, uint64, platform_redis.RedisError) {
	keys, nextCursor, err := s.client.SScan(
		ctx,
		key,
		cursor,
		match,
		count,
	).Result()
	if err != nil {
		return nil, 0, convertRedisError(err)
	}
	return keys, nextCursor, nil
}

func (s *SetCache) SInter(
	ctx context.Context,
	keys ...string,
) ([]string, platform_redis.RedisError) {
	cmd := s.client.SInter(ctx, keys...)
	res, err := cmd.Result()
	if err != nil {
		return nil, convertRedisError(err)
	}
	return res, nil
}
