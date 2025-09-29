package platform_redis

import (
	"context"
	"time"
)

type stringsCacheOperations interface {
	// Get returns the string value of the key.
	//
	// If the key is not found, item is empty string and error is [ErrNotFound].
	Get(ctx context.Context, key string) (string, RedisError)

	// MultiGet returns a slice of strings, the same length as the keys slice.
	// If the key is not found or value is not a string, item is empty string.
	MultiGet(ctx context.Context, keys []string) ([]string, RedisError)

	Set(ctx context.Context, key string, value string, ttl time.Duration) RedisError

	MultiSet(ctx context.Context, valuesMap map[string]string) RedisError

	Delete(ctx context.Context, key string) RedisError
}

type StringOptions struct {
	Expiration time.Duration
}
