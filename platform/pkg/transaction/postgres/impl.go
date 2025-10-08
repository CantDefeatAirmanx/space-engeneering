package platform_transaction_postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var _ platform_transaction.TxManager[platform_transaction.Executor] = (*TransactionManagerPostgres)(nil)

type TransactionManagerPostgres struct {
	pool *pgxpool.Pool
	cfg  *Config
}

func NewTransactionManagerPostgres(
	pool *pgxpool.Pool,
	opts ...ConfigOption,
) TransactionManagerPostgresInterface {
	cfg := NewConfig(opts...)

	return &TransactionManagerPostgres{pool: pool, cfg: cfg}
}

func (t *TransactionManagerPostgres) BeginTx(
	ctx context.Context,
	operationsFn func(ctx context.Context, tx platform_transaction.Transaction) error,
) error {
	pgxTx, err := t.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     pgx.ReadWrite,
		DeferrableMode: pgx.NotDeferrable,
	})
	if err != nil {
		return err
	}

	transaction := newTransaction(
		pgxTx,
		t.cfg.RollbackMaxAttempts,
		t.cfg.Logger,
	)
	withTransaction := t.createCtxWithTx(
		ctx,
		transaction,
	)
	withExecutor := t.createCtxWithExecutor(
		withTransaction,
		transaction,
	)
	if err := operationsFn(withExecutor, transaction); err != nil {
		transaction.RollbackWithRetry(ctx)
		return err
	}

	err = transaction.Commit(ctx)
	if err != nil {
		transaction.RollbackWithRetry(ctx)
		return err
	}

	return nil
}

func (t *TransactionManagerPostgres) createCtxWithTx(
	ctx context.Context,
	tx platform_transaction.Transaction,
) context.Context {
	return context.WithValue(
		ctx,
		platform_transaction.TransactionCtxKey,
		tx,
	)
}

func (t *TransactionManagerPostgres) createCtxWithExecutor(
	ctx context.Context,
	executor platform_transaction.Executor,
) context.Context {
	return context.WithValue(
		ctx,
		platform_transaction.ExecutorCtxKey,
		executor,
	)
}

func (t *TransactionManagerPostgres) ExtractExecutorFromCtx(
	ctx context.Context,
) (platform_transaction.Executor, error) {
	executor, ok := ctx.Value(
		platform_transaction.ExecutorCtxKey,
	).(platform_transaction.Executor)
	if !ok {
		return nil, errors.New("executor not found in context")
	}
	return executor, nil
}
