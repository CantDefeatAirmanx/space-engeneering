package api_order_v1

import (
	"context"
	"net/http"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

var paymentMethodMap = map[order_v1.PaymentMethod]model_order.PaymentMethod{
	order_v1.PaymentMethodCARD:          model_order.PaymentMethodCard,
	order_v1.PaymentMethodSBP:           model_order.PaymentMethodSBP,
	order_v1.PaymentMethodCREDITCARD:    model_order.PaymentMethodCreditCard,
	order_v1.PaymentMethodINVESTORMONEY: model_order.PaymentMethodInvestorMoney,
	order_v1.PaymentMethodUNKNOWN:       model_order.PaymentMethodUnknown,
}

func (api *api) PayOrder(
	ctx context.Context,
	req *order_v1.PayOrderRequestBody,
	params order_v1.PayOrderParams,
) (order_v1.PayOrderRes, error) {
	payRes, err := api.orderService.PayOrder(ctx, service_order.PayOrderParams{
		OrderUUID:     params.OrderUUID,
		PaymentMethod: paymentMethodMap[req.PaymentMethod],
	})
	if err != nil {
		res, err := handleServiceError[order_v1.PayOrderRes](err)
		if err != nil {
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: internalServerErrorMessage,
			}, nil
		}
		return res, nil
	}

	return &order_v1.PayOrderResponseBody{
		TransactionUUID: payRes.TransactionUUID,
	}, nil
}
