package platform_redis

import "errors"

type RedisError error

var (
	ErrNotFound     = errors.New("item not found")
	ErrConfigError  = errors.New("redis config error")
	ErrPoolTimeout  = errors.New("redis pool timeout")
	ErrClientClosed = errors.New("redis client closed")
	ErrNetworkError = errors.New("redis network error")
	ErrAuthError    = errors.New("redis auth error")
	ErrUnknownError = errors.New("redis unknown error")
)
