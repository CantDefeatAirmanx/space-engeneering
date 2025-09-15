package api_order_v1

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	model_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order/converter"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *Api) GetOrder(
	ctx context.Context,
	params order_v1.GetOrderParams,
) (order_v1.GetOrderRes, error) {
	contexts.GetLogParamsSetterFunc(ctx)(
		[]zap.Field{
			zap.String(orderUUIDLogKey, params.OrderUUID),
		},
	)

	order, err := api.orderService.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		res, err := handleServiceError[order_v1.GetOrderRes](err)
		if err != nil {
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: internalServerErrorMessage,
			}, nil
		}
		return res, nil
	}

	apiOrder := model_order_converter.ToApi(order)

	return &apiOrder, nil
}
