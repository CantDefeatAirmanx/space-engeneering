package repository_session_redis

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
)

func (s *SessionRepositoryRedisImpl) CreateUserSession(
	ctx context.Context,
	params model_session.CreateUserSessionParams,
) (*model_session.Session, error) {
	sessionUUID := uuid.Must(uuid.NewV7())

	session := model_session.Session{
		UUID:      sessionUUID.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: params.ExpiresAt,
	}

	// ToDo: Сделать в platform/redis возможность выполнения транзакций
	userSessionsKey := fmt.Sprintf(userSessionsKeyV1, params.UserUUID)
	_, err := s.redisCache.
		Set().
		SAdd(
			ctx,
			userSessionsKey,
			sessionUUID.String(),
		)
	if err != nil {
		return nil, err
	}
	_, err = s.redisCache.
		Hash().
		HSet(
			ctx,
			fmt.Sprintf(sessionDataKeyV1, sessionUUID.String()),
			map[string]string{
				sessionHashUUIDPropKey:      session.UUID,
				sessionHashCreatedAtPropKey: session.CreatedAt.Format(dateFormat),
				sessionHashUpdatedAtPropKey: session.UpdatedAt.Format(dateFormat),
				sessionHashExpiresAtPropKey: session.ExpiresAt.Format(dateFormat),
			},
		)
	if err != nil {
		return nil, err
	}

	return &session, nil
}
