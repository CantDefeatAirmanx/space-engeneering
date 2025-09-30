package platform_transaction_postgres

import (
	platform_transaction "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/transaction"
)

type TransactionManagerPostgresInterface interface {
	platform_transaction.TxManager[platform_transaction.Executor]
}
