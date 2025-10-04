package service_user

import (
	"context"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func (u *UserServiceImpl) GetUser(
	ctx context.Context,
	userUUID string,
) (*model_user.User, error) {
	user, err := u.userRepo.GetUserShortInfoWithHashPwd(
		ctx,
		model_user.UserFilter{UUID: userUUID},
	)
	if err != nil {
		return nil, err
	}
	notificationMethods, err := u.
		notificationMethodRepo.
		GetUserNotificationMethods(
			ctx,
			userUUID,
		)
	if err != nil {
		return nil, err
	}

	return &model_user.User{
		UUID: user.UUID,
		Info: mapInfoOnUserModel(
			user.UserShortInfo,
			notificationMethods,
		),
	}, nil
}

func mapInfoOnUserModel(
	user model_user.UserShortInfo,
	notificationMethods []model_notification_method.NotificationMethod,
) model_user.UserFullInfo {
	return model_user.UserFullInfo{
		UserShortInfo: model_user.UserShortInfo{
			Login: user.Login,
			Email: user.Email,
		},

		NotificationMethods: notificationMethods,
	}
}
