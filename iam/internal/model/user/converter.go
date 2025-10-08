package model_user

import (
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
)

func ConverUserToModel(proto *common_v1.User) *User {
	result := &User{
		UUID: proto.Uuid,
	}

	if proto.Info != nil {
		result.Info = ConvertUserInfoToModel(proto.Info)
	}

	if proto.UpdatedAt != nil {
		result.UpdatedAt = proto.UpdatedAt.AsTime()
	}

	if proto.CreatedAt != nil {
		result.CreatedAt = proto.CreatedAt.AsTime()
	}

	return result
}

func ConvertUserToProto(model *User) *common_v1.User {
	result := &common_v1.User{
		Uuid: model.UUID,

		UpdatedAt: timestamppb.New(model.UpdatedAt),
		CreatedAt: timestamppb.New(model.CreatedAt),
	}
	userInfo := ConvertUserInfoToProto(&model.Info)
	result.Info = &userInfo

	return result
}

func ConvertUserInfoToModel(proto *common_v1.UserInfo) UserFullInfo {
	return UserFullInfo{
		UserShortInfo: UserShortInfo{
			Login: proto.Login,
			Email: proto.Email,
		},

		NotificationMethods: lo.Map(
			proto.NotificationMethods,
			func(item *common_v1.NotificationMethod, _ int) model_notification_method.NotificationMethod {
				return *model_notification_method.ConvertNotificationMethodToModel(item)
			},
		),
	}
}

func ConvertUserInfoToProto(model *UserFullInfo) common_v1.UserInfo {
	return common_v1.UserInfo{
		Login: model.Login,
		Email: model.Email,

		NotificationMethods: lo.Map(
			model.NotificationMethods,
			func(item model_notification_method.NotificationMethod, _ int) *common_v1.NotificationMethod {
				return model_notification_method.ConvertNotificationMethodToProto(&item)
			},
		),
	}
}
