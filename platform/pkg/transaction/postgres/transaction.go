package platform_transaction_postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"

	platform_postgres "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/db/postgres"
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

var (
	_ platform_transaction.Transaction = (*transaction)(nil)
	_ platform_postgres.Executor       = (*transaction)(nil)
)

type transaction struct {
	tx                  pgx.Tx
	isRollBackCompleted bool

	maxAttempts int
	logger      platform_transaction.Logger
}

func newTransaction(
	tx pgx.Tx,
	maxAttempts int,
	logger platform_transaction.Logger,
) *transaction {
	return &transaction{tx: tx, maxAttempts: maxAttempts, logger: logger}
}

func (t *transaction) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	return t.tx.Exec(ctx, query, args...)
}

func (t *transaction) Query(ctx context.Context, query string, args ...any) (pgx.Rows, error) {
	return t.tx.Query(ctx, query, args...)
}

func (t *transaction) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	return t.tx.QueryRow(ctx, query, args...)
}

func (t *transaction) Commit(ctx context.Context) error {
	return t.tx.Commit(ctx)
}

func (t *transaction) Rollback(ctx context.Context) error {
	if t.isRollBackCompleted {
		return nil
	}

	err := t.tx.Rollback(ctx)
	if err != nil {
		return err
	}
	t.isRollBackCompleted = true
	return nil
}

func (t *transaction) RollbackWithRetry(ctx context.Context) {
	for attempt := range t.maxAttempts {
		if err := t.tx.Rollback(ctx); err != nil {
			if attempt == t.maxAttempts-1 {
				t.logger.Error("Failed to rollback transaction", zap.Error(err))
				return
			}

			ticker := time.NewTicker(10 * time.Second)
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				ticker.Stop()
				continue
			}
		} else {
			return
		}
	}
}
