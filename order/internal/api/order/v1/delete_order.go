package api_order_v1

import (
	"context"
	"net/http"

	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *api) DeleteOrder(
	ctx context.Context,
	params order_v1.DeleteOrderParams,
) (order_v1.DeleteOrderRes, error) {
	err := api.orderService.DeleteOrder(ctx, params.OrderUUID)
	if err != nil {
		res, err := handleServiceError[order_v1.DeleteOrderRes](err)
		if err != nil {
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: internalServerErrorMessage,
			}, nil
		}
		return res, nil
	}

	return &order_v1.DeleteOrderOK{}, nil
}
