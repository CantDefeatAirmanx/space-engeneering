package model_notification_method

import (
	common_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/common/v1"
)

var (
	notificationMethodProviderNameToProto = map[string]common_v1.NotificationMethodProviderName{
		"telegram": common_v1.NotificationMethodProviderName_NOTIFICATION_METHOD_PROVIDER_NAME_TELEGRAM,
	}

	notificationMethodProviderNameToModel = map[common_v1.NotificationMethodProviderName]ProviderName{
		common_v1.NotificationMethodProviderName_NOTIFICATION_METHOD_PROVIDER_NAME_TELEGRAM: ProviderNameTelegram,
	}
)

func ConvertNotificationMethodToModel(proto *common_v1.NotificationMethod) *NotificationMethod {
	return &NotificationMethod{
		ProviderName: notificationMethodProviderNameToModel[proto.ProviderName],
		Target:       proto.Target,
	}
}

func ConvertNotificationMethodToProto(model *NotificationMethod) *common_v1.NotificationMethod {
	return &common_v1.NotificationMethod{
		ProviderName: notificationMethodProviderNameToProto[string(model.ProviderName)],
		Target:       model.Target,
	}
}
