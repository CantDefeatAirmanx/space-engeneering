package repository_session_redis

import (
	"context"
	"fmt"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

func (s *SessionRepositoryRedisImpl) GetSession(
	ctx context.Context,
	sessionUUID string,
) (*model_session.Session, error) {
	sessionInfo, err := s.redisCache.
		Hash().
		HGetAll(
			ctx,
			fmt.Sprintf(sessionDataKeyV1, sessionUUID),
		)
	if err != nil {
		if err == platform_redis.ErrNotFound {
			return nil, model_session.ErrSessionExpired
		}
		return nil, err
	}

	return convertSessionToModel(sessionInfo)
}
