package repository_session_redis

import (
	"time"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
)

func convertSessionToModel(
	session SessionInfoRedisType,
) (*model_session.Session, error) {
	res := model_session.Session{
		UUID: session[sessionHashUUIDPropKey],
	}

	dateKeys := []string{
		sessionHashCreatedAtPropKey,
		sessionHashUpdatedAtPropKey,
		sessionHashExpiresAtPropKey,
	}
	for _, dateKey := range dateKeys {
		err := parseTimeWithUpdateResult(&res, session[dateKey])
		if err != nil {
			return nil, err
		}
	}

	return &res, nil
}

func parseTimeWithUpdateResult(result *model_session.Session, dateString string) error {
	date, err := time.Parse(
		dateFormat,
		dateString,
	)
	if err != nil {
		return err
	}
	result.CreatedAt = date
	return nil
}
