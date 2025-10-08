package repository_user

import (
	"context"

	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model_user.UserInfoWithHashPwd) (*model_user.UserShortInfo, error)
	GetUserShortInfo(ctx context.Context, filter model_user.UserFilter) (*model_user.UserInfoWithHashPwd, error)

	platform_transaction.WithExecutor[UserRepository, platform_transaction.Executor]
}
