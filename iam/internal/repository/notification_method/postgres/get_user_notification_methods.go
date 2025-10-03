package repository_notification_method_postgres

import (
	"context"

	"github.com/Masterminds/squirrel"

	model_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/notification_method"
)

func (n *NotificationMethodRepositoryPostgres) GetUserNotificationMethods(
	ctx context.Context,
	userUUID string,
) ([]model_notification_method.NotificationMethod, error) {
	query, args, err := squirrel.
		Select(
			"n."+columnNotificationMethodUUID,
			"n."+columnNotificationMethodProviderName,
			"n."+columnNotificationMethodTarget,
		).
		From(tableUserToNotificationMethods + " AS utn").
		Where(squirrel.Eq{columnUserToNotificationMethodUserUUID: userUUID}).
		Join(tableNotificationMethods + " AS n ON n.uuid = utn.notification_method_uuid").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := n.executor.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	repoMethods := make([]NotificationMethod, 0)
	for rows.Next() {
		var nm NotificationMethod
		err := rows.Scan(
			&nm.NotificationMethodUUID,
			&nm.NotificationMethodProviderName,
			&nm.NotificationMethodTarget,
		)
		if err != nil {
			return nil, err
		}
		repoMethods = append(repoMethods, nm)
	}

	return convertNotificationMethodsToModel(repoMethods), nil
}
