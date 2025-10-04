package repository_session_redis

import (
	"context"
	"errors"
	"fmt"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	platform_redis "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/cache/redis"
)

func (s *SessionRepositoryRedisImpl) DeleteSession(
	ctx context.Context,
	sessionUUID string,
) error {
	curSession, err := s.GetSession(ctx, sessionUUID)
	if err != nil {
		if errors.Is(err, platform_redis.ErrNotFound) {
			return model_session.ErrNotFound
		}
		return err
	}

	// ToDo: Сделать в platform/redis возможность выполнения транзакций
	_, err = s.redisCache.
		Set().
		SRem(
			ctx,
			fmt.Sprintf(userSessionsKeyV1, curSession.UserUUID),
			sessionUUID,
		)
	if err != nil {
		return err
	}
	_, err = s.redisCache.
		Hash().
		HDel(
			ctx,
			fmt.Sprintf(sessionDataKeyV1, curSession.UUID),
			sessionHashUUIDPropKey,
		)
	if err != nil {
		return err
	}
	return nil
}
