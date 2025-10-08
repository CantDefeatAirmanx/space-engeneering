package repository_session_redis

import (
	"time"

	repository_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/session"
	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

var _ repository_session.SessionRepository = (*SessionRepositoryRedisImpl)(nil)

type SessionRepositoryRedisImpl struct {
	redisCache platform_redis.RedisCache
	ttl        time.Duration
}

func NewSessionRepositoryRedisImpl(
	redisCache platform_redis.RedisCache,
	ttl time.Duration,
) repository_session.SessionRepository {
	return &SessionRepositoryRedisImpl{
		redisCache: redisCache,
		ttl:        ttl,
	}
}
