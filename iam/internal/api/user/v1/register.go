package api_user_v1

import (
	"context"

	"github.com/samber/lo"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
	user_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/user/v1"
)

func (a *Api) Register(
	ctx context.Context, req *user_v1.RegisterRequest,
) (*user_v1.RegisterResponse, error) {
	user, err := a.userService.Register(ctx, mapReqToRegisterData(req))
	if err != nil {
		return nil, err
	}

	return &user_v1.RegisterResponse{
		UserUuid: user.UserUUID,
	}, nil
}

func mapReqToRegisterData(req *user_v1.RegisterRequest) model_user.UserRegisterData {
	return model_user.UserRegisterData{
		Login:    req.Info.Login,
		Email:    req.Info.Email,
		Password: req.Password,
		NotificationMethods: lo.Map(
			req.Info.NotificationMethods,
			func(method *common_v1.NotificationMethod, _ int) model_notification_method.NotificationMethod {
				return *model_notification_method.ConvertNotificationMethodToModel(method)
			},
		),
	}
}
