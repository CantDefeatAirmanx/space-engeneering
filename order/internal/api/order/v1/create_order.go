package api_order_v1

import (
	"context"
	"net/http"

	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *api) CreateOrder(
	ctx context.Context,
	req *order_v1.CreateOrderRequestBody,
) (order_v1.CreateOrderRes, error) {
	result, err := api.orderService.CreateOrder(
		ctx,
		service_order.CreateOrderParams{
			UserUUID:  req.UserUUID,
			PartUuids: req.PartUuids,
		},
	)
	if err != nil {
		return &order_v1.InternalServerError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	return &order_v1.CreateOrderResponseBody{
		OrderUUID:  result.OrderUUID,
		TotalPrice: result.TotalPrice,
	}, nil
}
