package api_payment_v1

import (
	"context"

	"go.uber.org/zap"

	model_payment_method_converter "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/model/payment_method/converter"
	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interceptor"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

func (api Api) PayOrder(
	ctx context.Context,
	request *payment_v1.PayOrderRequest,
) (*payment_v1.PayOrderResponse, error) {
	ctxWithLogParams := context.WithValue(ctx, interceptor.LogParamsKey, []zap.Field{
		zap.String(orderUUIDLogKey, request.OrderUuid),
		zap.String(userUUIDLogKey, request.UserUuid),
	})

	paymentData, err := api.payOrderService.PayOrder(
		ctxWithLogParams,
		service_pay_order.PayOrderMethodParams{
			OrderUUID:     request.OrderUuid,
			UserUUID:      request.UserUuid,
			PaymentMethod: model_payment_method_converter.ToModel(request.PaymentMethod),
		},
	)
	if err != nil {
		return nil, err
	}

	return &payment_v1.PayOrderResponse{
		TransactionUuid: paymentData.TransactionUUID,
	}, nil
}
