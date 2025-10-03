package service_user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func (u *UserServiceImpl) Register(
	ctx context.Context,
	user model_user.UserRegisterData,
) (*RegisterResult, error) {
	var result RegisterResult

	_, err := u.txManager.BeginTx(ctx, func(ctx context.Context) error {
		executor, err := u.txManager.ExtractExecutorFromCtx(ctx)
		if err != nil {
			return err
		}
		tx, err := u.txManager.ExtractTransactionFromCtx(ctx)
		if err != nil {
			return err
		}

		txUserRepo := u.userRepo.WithExecutor(executor)
		txNotificationMethodRepo := u.notificationMethodRepo.WithExecutor(executor)

		uuid := uuid.Must(uuid.NewV7())
		passwordHash, err := u.pwdHasher.Hash([]byte(user.Password))
		if err != nil {
			return err
		}

		_, err = txUserRepo.CreateUser(ctx, &model_user.UserInfoWithHashPwd{
			UUID:         uuid.String(),
			PasswordHash: string(passwordHash),

			UserShortInfo: model_user.UserShortInfo{
				Login: user.Login,
				Email: user.Email,
			},

			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			tx.RollbackWithRetry(ctx)
			return err
		}

		methodsWithUUIDs := createNotificationMethodsWithUUIDs(user.NotificationMethods)

		err = txNotificationMethodRepo.CreateNotificationMethods(
			ctx,
			methodsWithUUIDs,
		)
		if err != nil {
			tx.RollbackWithRetry(ctx)
			return err
		}

		err = txNotificationMethodRepo.SetUserNotificationMethods(
			ctx,
			uuid.String(),
			methodsWithUUIDs,
		)
		if err != nil {
			tx.RollbackWithRetry(ctx)
			return err
		}

		err = tx.Commit(ctx)
		if err != nil {
			return err
		}

		result = RegisterResult{
			UserUUID: uuid.String(),
		}
		return nil
	})

	return &result, err
}

func createNotificationMethodsWithUUIDs(
	notificationMethods []model_notification_method.NotificationMethod,
) []model_notification_method.NotificationMethod {
	return lo.Map(
		notificationMethods,
		func(method model_notification_method.NotificationMethod, _ int) model_notification_method.NotificationMethod {
			return model_notification_method.NotificationMethod{
				UUID:         uuid.Must(uuid.NewV7()).String(),
				ProviderName: method.ProviderName,
				Target:       method.Target,
			}
		},
	)
}
