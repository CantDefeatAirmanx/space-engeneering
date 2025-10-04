package platform_redis

import "context"

type SetCache interface {
	// SAdd adds the values to the set.
	//
	// If the key does not exist, it will be created.
	//
	// Returns the number of added members.
	// Possible errors: [ErrWrongType], [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed], [ErrTooManyClients].
	SAdd(ctx context.Context, key string, values ...string) (int64, RedisError)

	// SRem removes the values from the set.
	//
	// If the key does not exist, it is a noop.
	//
	// Returns the number of removed members.
	// Possible errors: [ErrWrongType], [ErrReadOnly], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	SRem(ctx context.Context, key string, values ...string) (int64, RedisError)

	// SIsMember checks if the value is a member of the set.
	//
	// If the key does not exist, it is a noop.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	SIsMember(ctx context.Context, key, value string) (bool, RedisError)

	// SInter returns the intersection of the sets.
	//
	// If the key does not exist, it is a noop.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	SInter(ctx context.Context, keys ...string) ([]string, RedisError)

	// SCard returns the number of members in the set.
	//
	// If the key does not exist, it is a noop.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	SCard(ctx context.Context, key string) (int64, RedisError)

	// SScan returns the iterator over the set.
	//
	// If the key does not exist, it is a noop.
	//
	// Returns the keys and the next cursor position for pagination.
	// Possible errors: [ErrWrongType], [ErrTimeout], [ErrConnectionLost],
	// [ErrPoolTimeout], [ErrPoolExhausted], [ErrClientClosed].
	SScan(ctx context.Context, key string, cursor uint64, match string, count int64) (keys []string, nextCursor uint64, err RedisError)
}
