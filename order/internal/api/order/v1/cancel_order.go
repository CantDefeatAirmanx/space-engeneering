package api_order_v1

import (
	"context"
	"net/http"

	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *api) CancelOrder(
	ctx context.Context,
	params order_v1.CancelOrderParams,
) (order_v1.CancelOrderRes, error) {
	err := api.orderService.CancelOrder(ctx, params.OrderUUID)
	if err != nil {
		res, err := handleServiceError[order_v1.CancelOrderRes](err)
		if err != nil {
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: internalServerErrorMessage,
			}, nil
		}
		return res, nil
	}

	return &order_v1.CancelOrderOK{}, nil
}
