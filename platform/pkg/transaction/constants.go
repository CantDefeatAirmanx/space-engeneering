package platform_transaction

type (
	transactionCtxKey string
	executorCtxKey    string
)

const (
	TransactionCtxKey transactionCtxKey = "transaction"
	ExecutorCtxKey    executorCtxKey    = "executor"
)
