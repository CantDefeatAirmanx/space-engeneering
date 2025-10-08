package platform_transaction

import (
	"context"

	"go.uber.org/zap"

	platform_postgres "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/db/postgres"
)

// Executor - интерфейс для работы с конкретной реализацией базы данных.
//
// На текущий момент используется только postgres.
type Executor interface {
	platform_postgres.Executor
}

type WithExecutor[GRepository any, GExecutor Executor] interface {
	WithExecutor(executor Executor) GRepository
}

type TxManager[GExecutor Executor] interface {
	// Начинает новую транзакцию.
	//
	// Добавляет транзакцию в контекст.
	BeginTx(
		ctx context.Context,
		operationsFn func(ctx context.Context, tx Transaction) error,
	) error

	// Извлекает конкретную реализацию базы данных из контекста.
	//
	// Используется внутри operationsFn для получения конкретной реализации базы данных.
	ExtractExecutorFromCtx(ctx context.Context) (GExecutor, error)
}

type Transaction interface {
	// Коммитит транзакцию.
	//
	// Если транзакция уже закомичена, то это noop.
	Commit(ctx context.Context) error

	// Откатывает транзакцию, если она не закомичена.
	//
	// Если транзакция уже закомичена, то это noop.
	Rollback(ctx context.Context) error

	// Откатывает транзакцию, если она не закомичена.
	//
	// Если транзакция уже закомичена, то это noop.
	//
	// Повторяет откат до тех пор, пока не произойдет успешный откат или не будет достигнуто максимальное количество попыток.
	// Максимальное количество попыток задается в конфигурации.
	RollbackWithRetry(ctx context.Context)
}

type Logger interface {
	Error(msg string, fields ...zap.Field)
}
