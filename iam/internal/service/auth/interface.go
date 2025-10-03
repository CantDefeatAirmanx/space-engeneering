package service_auth

import (
	"context"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

type AuthService interface {
	Login(
		ctx context.Context,
		loginWithPasswordData *model_session.LoginWithPasswordData,
	) (*LoginWithPasswordResult, error)

	WhoAmI(
		ctx context.Context,
		params WhoAmIParams,
	) (*WhoAmIResult, error)
}

type LoginWithPasswordResult struct {
	SessionUUID string
}

type WhoAmIParams struct {
	SessionUUID string
}

type WhoAmIResult struct {
	Session *model_session.Session
	User    *model_user.User
}
