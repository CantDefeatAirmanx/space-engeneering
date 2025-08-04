package service_pay_order

import (
	"context"

	model_payment_method "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/model/payment_method"
)

type PayOrderService interface {
	PayOrder(ctx context.Context, params PayOrderMethodParams) (*PayOrderMethodReturn, error)
}

type PayOrderMethodParams struct {
	OrderUUID     string
	UserUUID      string
	PaymentMethod model_payment_method.PaymentMethod
}

type PayOrderMethodReturn struct {
	TransactionUUID string
}
