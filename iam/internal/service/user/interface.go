package service_user

import (
	"context"

	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

type UserService interface {
	GetUser(
		ctx context.Context,
		userUUID string,
	) (*model_user.User, error)

	Register(
		ctx context.Context,
		user model_user.UserRegisterData,
	) (*RegisterResult, error)
}

type RegisterResult struct {
	UserUUID string
}
