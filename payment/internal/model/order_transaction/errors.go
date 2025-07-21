package model_order_transaction

import "fmt"

type ErrOrderTransactionInternal struct {
	OrderUUID       string
	TransactionUUID string
	Err             error
}

func (e ErrOrderTransactionInternal) Error() string {
	return fmt.Sprintf(
		"Order transaction internal error. OrderUUID: %s. TransactionUUID: %s. %v",
		e.OrderUUID,
		e.TransactionUUID,
		e.Err,
	)
}

func (e ErrOrderTransactionInternal) Unwrap() error {
	return e.Err
}

func (e ErrOrderTransactionInternal) Is(target error) bool {
	_, ok := target.(*ErrOrderTransactionInternal)
	return ok
}
