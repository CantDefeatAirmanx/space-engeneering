package repository_user_postgres

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	UUID         pgtype.UUID
	Login        string
	Email        string
	PasswordHash string

	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
