package platform_redis

import (
	"context"
	"time"
)

type StringCache interface {
	// Get returns the string value of the key.
	//
	// If the key is not found, item is empty string and error is [ErrNotFound].
	// Possible errors: [ErrNotFound], [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	Get(ctx context.Context, key string) (string, RedisError)

	// MultiGet returns a slice of strings, the same length as the keys slice.
	//
	// If the key is not found or value is not a string, item is empty string.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	MultiGet(ctx context.Context, keys []string) ([]string, RedisError)

	// Set sets the string value of the key.
	//
	// If the key already exists, it will be overwritten.
	// Possible errors: [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed], [ErrTooManyClients].
	Set(ctx context.Context, key string, value string, ttl time.Duration) RedisError

	// MultiSet sets the string values of the keys.
	//
	// If the key already exists, it will be overwritten.
	// Possible errors: [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed], [ErrTooManyClients].
	MultiSet(ctx context.Context, valuesMap map[string]string) RedisError

	// Delete deletes the key.
	//
	// If the key does not exist, it is a noop.
	// Possible errors: [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	Delete(ctx context.Context, key string) RedisError
}

type StringOptions struct {
	Expiration time.Duration
}
