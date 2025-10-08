package repository_notification_method_postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
)

func (n *NotificationMethodRepositoryPostgres) CreateNotificationMethods(
	ctx context.Context,
	notificationMethods []model_notification_method.NotificationMethod,
) error {
	repoMethods, err := convertNotificationMethodsToRepo(notificationMethods)
	if err != nil {
		return fmt.Errorf("%w: %v", model_notification_method.ErrInvalidArguments, err)
	}

	queryBuilder := squirrel.
		Insert(tableNotificationMethods).
		Columns(
			columnNotificationMethodUUID,
			columnNotificationMethodProviderName,
		).
		PlaceholderFormat(squirrel.Dollar)
	for _, method := range repoMethods {
		queryBuilder = queryBuilder.Values(
			method.NotificationMethodUUID,
			method.NotificationMethodProviderName,
		)
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
