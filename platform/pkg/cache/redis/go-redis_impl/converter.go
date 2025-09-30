package platform_redis_redisgo

import (
	"context"
	"errors"
	"io"
	"net"
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

	if errors.Is(err, redis.ErrPoolExhausted) {
		return platform_redis.ErrPoolExhausted
	}

	if errors.Is(err, redis.ErrClosed) {
		return platform_redis.ErrClientClosed
	}

	if isIOError(err) {
		return platform_redis.ErrConnectionLost
	}

	if isNetworkTimeout(err) {
		return platform_redis.ErrTimeout
	}

	errString := err.Error()
	switch {
	case strings.HasPrefix(errString, "WRONGTYPE"):
		return errors.Join(platform_redis.ErrWrongType, err)
	case strings.HasPrefix(errString, "READONLY"):
		return platform_redis.ErrReadOnly
	case strings.HasPrefix(errString, "CLUSTERDOWN"):
		return platform_redis.ErrClusterDown
	case strings.Contains(errString, "max number of clients reached"):
		return platform_redis.ErrTooManyClients
	case strings.Contains(errString, "NOAUTH"):
		return platform_redis.ErrAuthError
	case strings.Contains(errString, "connection refused"):
		return platform_redis.ErrNetworkError
	}

	return errors.Join(platform_redis.ErrUnknownError, err)
}

func isContextError(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

func isIOError(err error) bool {
	return errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF)
}

func isNetworkTimeout(err error) bool {
	netErr, ok := err.(net.Error)
	return ok && netErr.Timeout()
}
