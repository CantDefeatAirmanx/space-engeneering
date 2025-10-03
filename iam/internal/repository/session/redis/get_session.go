package repository_session_redis

import (
	"context"
	"fmt"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
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
		return nil, err
	}

	return convertSessionToModel(sessionInfo)
}
