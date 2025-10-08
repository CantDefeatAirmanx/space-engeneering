package service_auth

import (
	"context"
	"time"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func (a *AuthServiceImpl) Login(
	ctx context.Context,
	loginWithPasswordData *model_session.LoginWithPasswordData,
) (*LoginWithPasswordResult, error) {
	user, err := a.userRepository.GetUserShortInfo(
		ctx,
		model_user.UserFilter{
			Login: loginWithPasswordData.Login,
			Email: loginWithPasswordData.Email,
		},
	)
	if err != nil {
		return nil, model_session.ErrInvalidCredentials
	}

	if !a.pwdHasher.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(loginWithPasswordData.Password),
	) {
		return nil, model_session.ErrInvalidCredentials
	}

	session, err := a.sessionRepository.CreateUserSession(
		ctx,
		model_session.CreateUserSessionParams{
			UserUUID:  user.UUID,
			ExpiresAt: time.Now().Add(time.Hour * 24),
		},
	)
	if err != nil {
		return nil, err
	}

	return &LoginWithPasswordResult{
		SessionUUID: session.UUID,
	}, nil
}
