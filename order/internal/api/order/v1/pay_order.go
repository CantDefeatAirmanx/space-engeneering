package api_order_v1

import (
	"context"
	"errors"
	"fmt"
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
		case errors.Is(err, &model_order.ErrOrderInternal{}):
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}, nil
		default:
			return &order_v1.InternalServerError{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("%s: %s", internalServerErrorMessage, err.Error()),
			}, nil
		}
	}

	return &order_v1.PayOrderResponseBody{
		TransactionUUID: payRes.TransactionUUID,
	}, nil
}
