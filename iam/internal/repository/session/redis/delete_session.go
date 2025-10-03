package repository_session_redis

import (
	"context"
	"fmt"
)

func (s *SessionRepositoryRedisImpl) DeleteSession(
	ctx context.Context,
	sessionUUID string,
) error {
	// ToDo: Сделать в platform/redis возможность выполнения транзакций
	_, err := s.redisCache.
		Set().
		SRem(
			ctx,
			fmt.Sprintf(userSessionsKeyV1, sessionUUID),
			sessionUUID,
		)
	if err != nil {
		return err
	}
	_, err = s.redisCache.
		Hash().
		HDel(
			ctx,
			fmt.Sprintf(sessionDataKeyV1, sessionUUID),
			sessionHashUUIDPropKey,
		)
	if err != nil {
		return err
	}
	return nil
}
