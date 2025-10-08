package repository_session

import (
	"context"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
)

type SessionRepository interface {
	CreateUserSession(
		ctx context.Context,
		params model_session.CreateUserSessionParams,
	) (*model_session.Session, error)

	GetSession(
		ctx context.Context,
		sessionUUID string,
	) (*model_session.Session, error)

	DeleteSession(
		ctx context.Context,
		sessionUUID string,
	) error
}
