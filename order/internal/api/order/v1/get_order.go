package api_order_v1

import (
	"context"
	"errors"
	"net/http"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	model_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order/converter"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *api) GetOrder(
	ctx context.Context,
	params order_v1.GetOrderParams,
) (order_v1.GetOrderRes, error) {
	order, err := api.orderService.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		switch {
		case errors.Is(err, &model_order.ErrOrderNotFound{}):
			return &order_v1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}, nil
		default:
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}, nil
		}
	}

	apiOrder := model_order_converter.ToApi(order)

	return &apiOrder, nil
}
