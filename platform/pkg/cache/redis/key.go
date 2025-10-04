package platform_redis

import (
	"context"
	"time"
)

// KeyCache contains methods for working with keys
type KeyCache interface {
	// Expire sets the expiration time for the key
	Expire(ctx context.Context, key string, ttl time.Duration) RedisError
}
