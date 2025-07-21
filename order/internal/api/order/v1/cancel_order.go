package api_order_v1

import (
	"context"
	"errors"
	"net/http"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *api) CancelOrder(
	ctx context.Context,
	params order_v1.CancelOrderParams,
) (order_v1.CancelOrderRes, error) {
	err := api.orderService.CancelOrder(ctx, params.OrderUUID)
	if err != nil {
		switch {
		case errors.Is(err, &model_order.ErrOrderNotFound{}):
			return &order_v1.NotFoundError{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			}, nil
		case errors.Is(err, &model_order.ErrOrderConflict{}):
			return &order_v1.ConflictError{
				Code:    http.StatusConflict,
				Message: err.Error(),
			}, nil
		default:
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}, nil
		}
	}

	return &order_v1.CancelOrderOK{}, nil
}
