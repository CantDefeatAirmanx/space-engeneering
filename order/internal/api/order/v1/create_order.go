package api_order_v1

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

func (api *Api) CreateOrder(
	ctx context.Context,
	req *order_v1.CreateOrderRequestBody,
) (order_v1.CreateOrderRes, error) {
	contexts.GetLogParamsSetterFunc(ctx)(
		[]zap.Field{
			zap.String(userUUIDLogKey, req.UserUUID),
			zap.Strings(partUUIDsLogKey, req.PartUuids),
		},
	)

	result, err := api.orderService.CreateOrder(
		ctx,
		service_order.CreateOrderParams{
			UserUUID:  req.UserUUID,
			PartUuids: req.PartUuids,
		},
	)
	if err != nil {
		res, err := handleServiceError[order_v1.CreateOrderRes](err)
		if err != nil {
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: internalServerErrorMessage,
			}, nil
		}
		return res, nil
	}

	return &order_v1.CreateOrderResponseBody{
		OrderUUID:  result.OrderUUID,
		TotalPrice: result.TotalPrice,
	}, nil
}
