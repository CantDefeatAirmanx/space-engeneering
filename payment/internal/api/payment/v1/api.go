package api_payment_v1

import (
	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

var _ payment_v1.PaymentServiceServer = (*Api)(nil)

type Api struct {
	payment_v1.UnimplementedPaymentServiceServer
	payOrderService service_pay_order.PayOrderService
}

type NewApiParams struct {
	PayOrderService service_pay_order.PayOrderService
}

func NewApi(params NewApiParams) *Api {
	return &Api{
		payOrderService: params.PayOrderService,
	}
}
