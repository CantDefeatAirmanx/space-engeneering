package repository_session_redis

import (
	"time"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
)

func convertSessionToModel(
	session SessionInfoRedisType,
) (*model_session.Session, error) {
	res := model_session.Session{
		UUID:     session[sessionHashUUIDPropKey],
		UserUUID: session[sessionHashUserUUIDPropKey],
	}

	createdAt, err := parseTime(session[sessionHashCreatedAtPropKey])
	if err != nil {
		return nil, err
	}
	res.CreatedAt = createdAt

	updatedAt, err := parseTime(session[sessionHashUpdatedAtPropKey])
	if err != nil {
		return nil, err
	}
	res.UpdatedAt = updatedAt

	expiresAt, err := parseTime(session[sessionHashExpiresAtPropKey])
	if err != nil {
		return nil, err
	}
	res.ExpiresAt = expiresAt

	return &res, nil
}

func parseTime(dateString string) (time.Time, error) {
	date, err := time.Parse(
		dateFormat,
		dateString,
	)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
