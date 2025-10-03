package repository_session_redis

import (
	repository_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/session"
	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

var _ repository_session.SessionRepository = (*SessionRepositoryRedisImpl)(nil)

type SessionRepositoryRedisImpl struct {
	redisCache platform_redis.RedisCache
}

func NewSessionRepositoryRedisImpl(
	redisCache platform_redis.RedisCache,
) repository_session.SessionRepository {
	return &SessionRepositoryRedisImpl{
		redisCache: redisCache,
	}
}
