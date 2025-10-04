package repository_user_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"

	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func (u *UserRepositoryPostgres) CreateUser(
	ctx context.Context,
	user *model_user.UserInfoWithHashPwd,
) (*model_user.UserShortInfo, error) {
	query, args, err := squirrel.Insert(tableUsers).
		Columns(columnUserUUID, columnUserLogin, columnUserEmail, columnUserPasswordHash).
		Values(user.UUID, user.Login, user.Email, user.PasswordHash).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	_, err = u.executor.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			return nil, handlePgError(pgErr)
		}

		return nil, err
	}

	return &model_user.UserShortInfo{
		Login: user.Login,
		Email: user.Email,
	}, nil
}

func handlePgError(pgErr *pgconn.PgError) error {
	if pgErr.Code == "23505" {
		switch pgErr.ConstraintName {
		case uniqueConstraintUsersLogin:
			return fmt.Errorf("%w: %v", model_user.ErrUserAlreadyExists, "login already exists")
		case uniqueConstraintUsersEmail:
			return fmt.Errorf("%w: %v", model_user.ErrUserAlreadyExists, "email already exists")
		}
	}
	return pgErr
}
