package service_order

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	client_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	kafka_events_order "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/order/v1"
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

	if order.Status != model_order.OrderStatusPendingPayment {
		return nil, fmt.Errorf("%w: %s", model_order.ErrOrderConflict, fmt.Sprintf("Order %s is not in pending payment status", params.OrderUUID))
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
		return nil, err
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
		return nil, err
	}

	paidEvent := kafka_events_order.OrderPaidEvent{
		EventUUID:     uuid.New().String(),
		OrderUUID:     order.OrderUUID,
		UserUUID:      order.UserUUID,
		PaymentMethod: kafka_events_order.PaymentMethod(params.PaymentMethod),
	}
	if err = s.orderProducer.ProduceOrderPaid(ctx, paidEvent); err != nil {
		return nil, err
	}

	return &PayOrderResult{
		TransactionUUID: payRes.TransactionUUID,
	}, nil
}
