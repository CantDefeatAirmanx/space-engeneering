package repository_session_redis

import "time"

type SessionRedisTypeKey = string

const (
	namespace = "auth:v1"

	userSessionsKeyV1 = namespace + ":user:{%s}:sessions"
	sessionDataKeyV1  = namespace + ":session:%s"

	sessionHashUUIDPropKey      SessionRedisTypeKey = "uuid"
	sessionHashCreatedAtPropKey SessionRedisTypeKey = "created_at"
	sessionHashUpdatedAtPropKey SessionRedisTypeKey = "updated_at"
	sessionHashExpiresAtPropKey SessionRedisTypeKey = "expires_at"
)

var dateFormat = time.RFC3339

type SessionInfoRedisType = map[SessionRedisTypeKey]string
