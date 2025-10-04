package service_auth

import (
	repository_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/notification_method"
	repository_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/session"
	repository_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/user"
	platform_pwdhasher "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/pwd-hasher"
)

var _ AuthService = (*AuthServiceImpl)(nil)

type AuthServiceImpl struct {
	userRepository               repository_user.UserRepository
	notificationMethodRepository repository_notification_method.NotificationMethodRepository
	sessionRepository            repository_session.SessionRepository
	pwdHasher                    platform_pwdhasher.PwdHasher
}

func NewAuthService(
	userRepository repository_user.UserRepository,
	notificationMethodRepository repository_notification_method.NotificationMethodRepository,
	sessionRepository repository_session.SessionRepository,
	pwdHasher platform_pwdhasher.PwdHasher,
) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository:               userRepository,
		notificationMethodRepository: notificationMethodRepository,
		sessionRepository:            sessionRepository,

		pwdHasher: pwdHasher,
	}
}
