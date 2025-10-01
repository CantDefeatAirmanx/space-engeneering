package platform_redis_redisgo

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

var _ platform_redis.HashCache = (*HashCache)(nil)

type HashCache struct {
	client *redis.Client
}

func NewHashCache(client *redis.Client) platform_redis.HashCache {
	return &HashCache{
		client: client,
	}
}

func (h *HashCache) HSet(
	ctx context.Context,
	key string,
	values map[string]string,
) (int64, platform_redis.RedisError) {
	cmd := h.client.HSet(ctx, key, values)
	res, err := cmd.Result()
	if err != nil {
		return 0, convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HGet(
	ctx context.Context,
	key string,
	field string,
) (string, platform_redis.RedisError) {
	cmd := h.client.HGet(ctx, key, field)
	res, err := cmd.Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", platform_redis.ErrNotFound
		}
		return "", convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HGetAll(
	ctx context.Context,
	key string,
) (map[string]string, platform_redis.RedisError) {
	cmd := h.client.HGetAll(ctx, key)
	res, err := cmd.Result()
	if err != nil {
		return nil, convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HMGet(
	ctx context.Context,
	key string,
	fields ...string,
) ([]string, platform_redis.RedisError) {
	cmd := h.client.HMGet(ctx, key, fields...)
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

func (h *HashCache) HDel(
	ctx context.Context,
	key string,
	fields ...string,
) (int64, platform_redis.RedisError) {
	cmd := h.client.HDel(ctx, key, fields...)
	res, err := cmd.Result()
	if err != nil {
		return 0, convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HExists(
	ctx context.Context,
	key string,
	field string,
) (bool, platform_redis.RedisError) {
	cmd := h.client.HExists(ctx, key, field)
	res, err := cmd.Result()
	if err != nil {
		return false, convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HLen(
	ctx context.Context,
	key string,
) (int64, platform_redis.RedisError) {
	cmd := h.client.HLen(ctx, key)
	res, err := cmd.Result()
	if err != nil {
		return 0, convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HScan(
	ctx context.Context,
	key string,
	cursor uint64,
	match string,
	count int64,
) ([]string, uint64, platform_redis.RedisError) {
	fields, nextCursor, err := h.client.HScan(
		ctx,
		key,
		cursor,
		match,
		count,
	).Result()
	if err != nil {
		return nil, 0, convertRedisError(err)
	}

	return fields, nextCursor, nil
}

func (h *HashCache) HKeys(
	ctx context.Context,
	key string,
) ([]string, platform_redis.RedisError) {
	cmd := h.client.HKeys(ctx, key)
	res, err := cmd.Result()
	if err != nil {
		return nil, convertRedisError(err)
	}

	return res, nil
}

func (h *HashCache) HVals(
	ctx context.Context,
	key string,
) ([]string, platform_redis.RedisError) {
	cmd := h.client.HVals(ctx, key)
	res, err := cmd.Result()
	if err != nil {
		return nil, convertRedisError(err)
	}

	return res, nil
}
