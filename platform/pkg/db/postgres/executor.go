package platform_postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	_ Executor = (*pgxpool.Pool)(nil)
	_ Executor = (*pgx.Conn)(nil)
	_ Executor = (pgx.Tx)(nil)
)

type Executor interface {
	Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...any) pgx.Row
}

type WithExecutor[T any] interface {
	WithExecutor(executor Executor) T
}
