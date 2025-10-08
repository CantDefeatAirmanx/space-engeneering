package repository_notification_method_postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/samber/lo"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
)

var (
	notificationMethodProviderNameToModel = map[NotificationMethodProviderName]model_notification_method.ProviderName{
		"telegram": model_notification_method.ProviderNameTelegram,
	}

	notificationMethodProviderNameToRepo = map[model_notification_method.ProviderName]NotificationMethodProviderName{
		model_notification_method.ProviderNameTelegram: "telegram",
	}
)

func convertNotificationMethodsToModel(
	notificationMethods []NotificationMethod,
) []model_notification_method.NotificationMethod {
	return lo.Map(
		notificationMethods,
		func(notificationMethod NotificationMethod, _ int) model_notification_method.NotificationMethod {
			return *convertNotificationMethodToModel(notificationMethod)
		},
	)
}

func convertNotificationMethodToModel(
	notificationMethod NotificationMethod,
) *model_notification_method.NotificationMethod {
	return &model_notification_method.NotificationMethod{
		UUID:         notificationMethod.NotificationMethodUUID.String(),
		ProviderName: notificationMethodProviderNameToModel[notificationMethod.NotificationMethodProviderName],
		Target:       notificationMethod.NotificationMethodTarget,
	}
}

func convertNotificationMethodsToRepo(
	notificationMethods []model_notification_method.NotificationMethod,
) ([]NotificationMethod, error) {
	res := make([]NotificationMethod, 0, len(notificationMethods))
	for _, notificationMethod := range notificationMethods {
		repo, err := convertNotificationMethodToRepo(&notificationMethod)
		if err != nil {
			return nil, err
		}
		res = append(res, *repo)
	}
	return res, nil
}

func convertNotificationMethodToRepo(
	notificationMethod *model_notification_method.NotificationMethod,
) (*NotificationMethod, error) {
	result := NotificationMethod{
		NotificationMethodProviderName: notificationMethodProviderNameToRepo[notificationMethod.ProviderName],
		NotificationMethodTarget:       notificationMethod.Target,
	}

	var uuid pgtype.UUID
	if err := uuid.Scan(notificationMethod.UUID); err != nil {
		return nil, err
	}
	result.NotificationMethodUUID = uuid

	return &result, nil
}
