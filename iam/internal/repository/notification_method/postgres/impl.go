package repository_notification_method_postgres

import (
	repository_notification_method "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/notification_method"
	platform_postgres "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/db/postgres"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var _ repository_notification_method.NotificationMethodRepository = (*NotificationMethodRepositoryPostgres)(nil)

type NotificationMethodRepositoryPostgres struct {
	executor platform_postgres.Executor
}

func NewNotificationMethodRepositoryPostgres(
	executor platform_postgres.Executor,
) repository_notification_method.NotificationMethodRepository {
	return &NotificationMethodRepositoryPostgres{
		executor: executor,
	}
}

func (n *NotificationMethodRepositoryPostgres) WithExecutor(
	executor platform_transaction.Executor,
) repository_notification_method.NotificationMethodRepository {
	return NewNotificationMethodRepositoryPostgres(executor)
}
