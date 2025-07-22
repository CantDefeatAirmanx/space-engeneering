package api_order_v1

import (
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

var _ order_v1.Handler = (*api)(nil)

const (
	internalServerErrorMessage = "internal server error"
)

type api struct {
	orderService service_order.OrderService
}

type NewApiParams struct {
	OrderService service_order.OrderService
}

func NewApi(params NewApiParams) *api {
	return &api{
		orderService: params.OrderService,
	}
}
