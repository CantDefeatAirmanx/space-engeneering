package platform_redis_redisgo

import "time"

const (
	defaultClientName      = "Unknown Redis Client"
	defaultPoolSize        = 10
	defaultReadTimeout     = 10 * time.Second
	defaultWriteTimeout    = 10 * time.Second
	defaultMaxRetries      = 3
	defaultConnMaxIdleTime = 10 * time.Minute
	defaultConnMaxLifetime = 0
)

var defaultPoolTimeout = defaultReadTimeout + 1*time.Second
