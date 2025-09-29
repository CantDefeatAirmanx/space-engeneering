package platform_redis_redisgo

import (
	"context"
	"errors"
	"strings"

	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
	"github.com/redis/go-redis/v9"
)

func convertRedisError(err error) platform_redis.RedisError {
	if err == nil {
		return nil
	}

	if isContextError(err) {
		return err
	}

	if errors.Is(err, redis.Nil) {
		return platform_redis.ErrNotFound
	}

	if errors.Is(err, redis.ErrPoolTimeout) {
		return platform_redis.ErrPoolTimeout
	}

	if errors.Is(err, redis.ErrClosed) {
		return platform_redis.ErrClientClosed
	}

	errString := err.Error()
	switch {
	case strings.Contains(errString, "NOAUTH"):
		return platform_redis.ErrAuthError
	case strings.Contains(errString, "connection refused"):
		return platform_redis.ErrNetworkError
	}

	return errors.Join(platform_redis.ErrUnknownError, err)
}

func isContextError(err error) bool {
	return errors.Is(
		err, context.Canceled) ||
		errors.Is(err, context.DeadlineExceeded) ||
		errors.Is(err, context.DeadlineExceeded)
}
