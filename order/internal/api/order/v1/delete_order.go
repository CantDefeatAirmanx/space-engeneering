package api_order_v1

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *Api) DeleteOrder(
	ctx context.Context,
	params order_v1.DeleteOrderParams,
) (order_v1.DeleteOrderRes, error) {
	contexts.GetLogParamsSetterFunc(ctx)(
		[]zap.Field{
			zap.String(orderUUIDLogKey, params.OrderUUID),
		},
	)

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
