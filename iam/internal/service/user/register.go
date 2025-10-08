package service_user

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
	repository_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/notification_method"
	repository_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/user"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

func (u *UserServiceImpl) Register(
	ctx context.Context,
	user model_user.UserRegisterData,
) (*RegisterResult, error) {
	var result RegisterResult

	err := u.txManager.BeginTx(ctx, func(ctx context.Context, tx platform_transaction.Transaction) error {
		txUserRepo, txNotificationMethodRepo,
			err := getTransactionalRepos(ctx, u)
		if err != nil {
			return err
		}

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
			return err
		}
		result = RegisterResult{
			UserUUID: uuid.String(),
		}

		methodsWithUUIDs := createNotificationMethodsWithUUIDs(
			user.NotificationMethods,
		)
		if len(methodsWithUUIDs) > 0 {
			return createUserNotificationMethods(
				ctx,
				txNotificationMethodRepo,
				uuid,
				methodsWithUUIDs,
			)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &result, err
}

func getTransactionalRepos(
	ctx context.Context,
	u *UserServiceImpl,
) (repository_user.UserRepository, repository_notification_method.NotificationMethodRepository, error) {
	executor, err := u.txManager.ExtractExecutorFromCtx(ctx)
	if err != nil {
		return nil, nil, err
	}
	if err != nil {
		return nil, nil, err
	}

	txUserRepo := u.userRepo.WithExecutor(executor)
	txNotificationMethodRepo := u.notificationMethodRepo.WithExecutor(executor)

	return txUserRepo, txNotificationMethodRepo, nil
}

func createUserNotificationMethods(
	ctx context.Context,
	txNotificationMethodRepo repository_notification_method.NotificationMethodRepository,
	userUuid uuid.UUID,
	methodsWithUUIDs []model_notification_method.NotificationMethod,
) error {
	err := txNotificationMethodRepo.CreateNotificationMethods(
		ctx,
		methodsWithUUIDs,
	)
	if err != nil {
		return err
	}

	err = txNotificationMethodRepo.SetUserNotificationMethods(
		ctx,
		userUuid.String(),
		methodsWithUUIDs,
	)
	if err != nil {
		return err
	}

	return nil
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
