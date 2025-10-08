package repository_notification_method_postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
)

func (n *NotificationMethodRepositoryPostgres) SetUserNotificationMethods(
	ctx context.Context,
	userUUID string,
	notificationMethods []model_notification_method.NotificationMethod,
) error {
	if len(notificationMethods) == 0 {
		return fmt.Errorf("%w: %v", model_notification_method.ErrInvalidArguments, "notification methods are empty")
	}

	repoMethods, err := convertNotificationMethodsToRepo(notificationMethods)
	if err != nil {
		return fmt.Errorf("%w: %v", model_notification_method.ErrInvalidArguments, err)
	}

	queryBuilder := squirrel.
		Insert(tableUserToNotificationMethods).
		Columns(
			columnUserToNotificationMethodUserUUID,
			columnUserToNotificationMethodNotificationMethodUUID,
			columnUserToNotificationMethodTarget,
		).
		PlaceholderFormat(squirrel.Dollar)
	for _, method := range repoMethods {
		queryBuilder = queryBuilder.Values(userUUID, method.NotificationMethodUUID, method.NotificationMethodTarget)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = n.executor.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
