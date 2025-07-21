package api_payment_v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	model_payment_method_converter "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/model/payment_method/converter"
	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

func (api Api) PayOrder(
	ctx context.Context,
	request *payment_v1.PayOrderRequest,
) (*payment_v1.PayOrderResponse, error) {
	paymentData, err := api.payOrderService.PayOrder(
		ctx,
		service_pay_order.PayOrderMethodParams{
			OrderUUID:     request.OrderUuid,
			UserUUID:      request.UserUuid,
			PaymentMethod: model_payment_method_converter.ToModel(request.PaymentMethod),
		},
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error. %v", err)
	}

	return &payment_v1.PayOrderResponse{
		TransactionUuid: paymentData.TransactionUUID,
	}, nil
}
