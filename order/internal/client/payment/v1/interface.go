package client_payment_v1

import (
	"context"
)

type PaymentV1Client interface {
	// PayOrder pays an order
	//
	// Errors:
	// - [client_payment_v1.ErrInternalServerError]: if the server is not available
	// - [client_payment_v1.ErrInvalidArguments]: if the arguments are invalid
	PayOrder(
		ctx context.Context,
		params PayOrderParams,
	) (*PayOrderResult, error)
}

type PayOrderParams struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod PaymentMethod
}

type PayOrderResult struct {
	TransactionUUID string
}

type PaymentMethod int

const (
	PaymentMethodUnknown       PaymentMethod = 0
	PaymentMethodCard          PaymentMethod = 1
	PaymentMethodSPB           PaymentMethod = 2
	PaymentMethodCreditCard    PaymentMethod = 3
	PaymentMethodInvestorMoney PaymentMethod = 4
)
