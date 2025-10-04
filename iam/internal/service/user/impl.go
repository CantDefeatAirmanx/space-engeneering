package service_user

import (
	repository_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/notification_method"
	repository_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/user"
	platform_pwdhasher "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/pwd-hasher"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var _ UserService = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	userRepo               repository_user.UserRepository
	notificationMethodRepo repository_notification_method.NotificationMethodRepository
	pwdHasher              platform_pwdhasher.PwdHasher
	txManager              platform_transaction.TxManager[platform_transaction.Executor]
}

func NewUserServiceImpl(
	userRepo repository_user.UserRepository,
	notificationMethodRepo repository_notification_method.NotificationMethodRepository,
	txManager platform_transaction.TxManager[platform_transaction.Executor],
	pwdHasher platform_pwdhasher.PwdHasher,
) UserService {
	return &UserServiceImpl{
		userRepo:               userRepo,
		notificationMethodRepo: notificationMethodRepo,
		txManager:              txManager,
		pwdHasher:              pwdHasher,
	}
}
