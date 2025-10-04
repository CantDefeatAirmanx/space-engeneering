package repository_user_postgres

import (
	repository_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/repository/user"
	platform_postgres "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/db/postgres"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var _ repository_user.UserRepository = (*UserRepositoryPostgres)(nil)

type UserRepositoryPostgres struct {
	executor platform_postgres.Executor
}

func NewUserRepositoryPostgres(
	executor platform_postgres.Executor,
) repository_user.UserRepository {
	return &UserRepositoryPostgres{executor: executor}
}

func (u *UserRepositoryPostgres) WithExecutor(
	executor platform_transaction.Executor,
) repository_user.UserRepository {
	return NewUserRepositoryPostgres(executor)
}
