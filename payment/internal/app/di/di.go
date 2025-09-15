package di

import (
	"context"

	api_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/api/payment/v1"
	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
)

type DiContainer struct {
	paymentV1Api    *api_payment_v1.Api
	payOrderService service_pay_order.PayOrderService
}

func NewDiContainer() *DiContainer {
	return &DiContainer{}
}

func (d *DiContainer) GetPaymentV1Api(
	ctx context.Context,
) *api_payment_v1.Api {
	if d.paymentV1Api != nil {
		return d.paymentV1Api
	}

	d.paymentV1Api = api_payment_v1.NewApi(
		api_payment_v1.NewApiParams{
			PayOrderService: d.GetPayOrderService(ctx),
		},
	)

	return d.paymentV1Api
}

func (d *DiContainer) GetPayOrderService(
	ctx context.Context,
) service_pay_order.PayOrderService {
	if d.payOrderService != nil {
		return d.payOrderService
	}

	payOrderService := service_pay_order.NewPayOrderServiceImpl()
	d.payOrderService = payOrderService

	return payOrderService
}
