package repository_user_postgres

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	model_user "github.com/CantDefeatAirmanx/space-engeneering/iam/internal/model/user"
)

func (u *UserRepositoryPostgres) GetUserShortInfoWithHashPwd(
	ctx context.Context,
	filter model_user.UserFilter,
) (*model_user.UserInfoWithHashPwd, error) {
	query, args, err := squirrel.
		Select(
			columnUserUUID,
			columnUserLogin,
			columnUserEmail,
			columnUserPasswordHash,

			columnUserCreatedAt,
			columnUserUpdatedAt,
		).
		From(tableUsers).
		Where(getFilterParams(filter)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user User
	err = u.executor.QueryRow(ctx, query, args...).Scan(
		&user.UUID,
		&user.Login,
		&user.Email,
		&user.PasswordHash,

		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model_user.ErrUserNotFound
		}
		return nil, err
	}

	return convertRepoUserToModelUserInfoWithHashPwd(&user), nil
}

func getFilterParams(filter model_user.UserFilter) squirrel.Eq {
	res := squirrel.Eq{}

	if filter.UUID != "" {
		res[columnUserUUID] = filter.UUID
	}

	if filter.Login != "" {
		res[columnUserLogin] = filter.Login
	}

	if filter.Email != "" {
		res[columnUserEmail] = filter.Email
	}

	return res
}
