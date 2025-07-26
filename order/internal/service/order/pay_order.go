package service_order

import (
	"context"
	"fmt"
	"time"

	client_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
)

const (
	payOrderTimeout = 3 * time.Second
)

var paymentMethodMap = map[model_order.PaymentMethod]client_payment_v1.PaymentMethod{
	model_order.PaymentMethodCard:          client_payment_v1.PaymentMethodCard,
	model_order.PaymentMethodSBP:           client_payment_v1.PaymentMethodSPB,
	model_order.PaymentMethodCreditCard:    client_payment_v1.PaymentMethodCreditCard,
	model_order.PaymentMethodInvestorMoney: client_payment_v1.PaymentMethodInvestorMoney,
	model_order.PaymentMethodUnknown:       client_payment_v1.PaymentMethodUnknown,
}

func (s *OrderServiceImpl) PayOrder(ctx context.Context, params PayOrderParams) (*PayOrderResult, error) {
	order, err := s.orderRepository.GetOrder(ctx, params.OrderUUID)
	if err != nil {
		return nil, err
	}

	orderModel := repository_order_converter.ToModel(order)

	if orderModel.Status != model_order.OrderStatusPendingPayment {
		return nil, &model_order.ErrOrderConflict{
			OrderUUID:  params.OrderUUID,
			ErrMessage: fmt.Sprintf("Order %s is not in pending payment status", params.OrderUUID),
		}
	}

	payDeadline, cancel := context.WithTimeout(
		ctx,
		payOrderTimeout,
	)
	defer cancel()

	payRes, err := s.paymentClient.PayOrder(
		payDeadline,
		client_payment_v1.PayOrderParams{
			OrderUUID:     params.OrderUUID,
			PaymentMethod: paymentMethodMap[params.PaymentMethod],
		},
	)
	if err != nil {
		return nil, &model_order.ErrOrderInternal{
			OrderUUID: params.OrderUUID,
			Err:       err,
		}
	}

	newStatus := model_order.OrderStatusPaid
	err = s.UpdateOrder(
		payDeadline,
		params.OrderUUID,
		UpdateOrderFields{
			Status:          &newStatus,
			TransactionUUID: &payRes.TransactionUUID,
		},
	)
	if err != nil {
		return nil, &model_order.ErrOrderInternal{
			OrderUUID: params.OrderUUID,
			Err:       err,
		}
	}

	return &PayOrderResult{
		TransactionUUID: payRes.TransactionUUID,
	}, nil
}
