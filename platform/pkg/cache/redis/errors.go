package platform_redis

import "errors"

type RedisError error

var (
	ErrNotFound       = errors.New("item not found")
	ErrConfigError    = errors.New("redis config error")
	ErrPoolTimeout    = errors.New("redis pool timeout")
	ErrPoolExhausted  = errors.New("redis pool exhausted")
	ErrClientClosed   = errors.New("redis client closed")
	ErrNetworkError   = errors.New("redis network error")
	ErrConnectionLost = errors.New("redis connection lost")
	ErrTimeout        = errors.New("redis operation timeout")
	ErrAuthError      = errors.New("redis auth error")
	ErrWrongType      = errors.New("redis wrong data type")
	ErrReadOnly       = errors.New("redis read-only mode")
	ErrClusterDown    = errors.New("redis cluster down")
	ErrTooManyClients = errors.New("redis too many clients")
	ErrUnknownError   = errors.New("redis unknown error")
)
