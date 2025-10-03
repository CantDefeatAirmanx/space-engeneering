package service_auth

import (
	"context"

	model_session "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/session"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func (a *AuthServiceImpl) WhoAmI(
	ctx context.Context,
	params WhoAmIParams,
) (*WhoAmIResult, error) {
	session, err := a.sessionRepository.GetSession(
		ctx,
		params.SessionUUID,
	)
	if err != nil {
		return nil, model_session.ErrUnauthorized
	}

	user, err := a.userRepository.GetUserShortInfoWithHashPwd(
		ctx,
		model_user.UserFilter{
			UUID: session.UUID,
		},
	)
	if err != nil {
		return nil, model_session.ErrUnauthorized
	}

	notificationMethods, err := a.notificationMethodRepository.
		GetUserNotificationMethods(ctx, user.UUID)
	if err != nil {
		return nil, err
	}

	userModel := model_user.User{
		UUID: user.UUID,
		Info: model_user.UserFullInfo{
			UserShortInfo:       user.UserShortInfo,
			NotificationMethods: notificationMethods,
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &WhoAmIResult{
		Session: session,
		User:    &userModel,
	}, nil
}
