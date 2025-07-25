package service_order_tests

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/mock"

	client_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
)

func (s *TestingSuite) TestPayOrderSuccess() {
	orderUUID := gofakeit.UUID()
	transactionUUID := gofakeit.UUID()

	repoOrder := repository_order_model.Order{
		OrderUUID: orderUUID,
		Status:    repository_order_model.OrderStatusPendingPayment,
	}

	s.repoMock.EXPECT().GetOrder(s.ctx, orderUUID).Return(
		&repoOrder,
		nil,
	)

	s.paymentClientMock.EXPECT().PayOrder(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, _ = ctx.Deadline()
			return true
		}),
		client_payment_v1.PayOrderParams{
			OrderUUID:     orderUUID,
			PaymentMethod: client_payment_v1.PaymentMethodCard,
		},
	).Return(
		&client_payment_v1.PayOrderResult{
			TransactionUUID: transactionUUID,
		},
		nil,
	)

	s.repoMock.EXPECT().UpdateOrderFields(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, _ = ctx.Deadline()
			return true
		}),
		orderUUID,
		mock.MatchedBy(func(update repository_order.UpdateOrderFields) bool {
			return true
		}),
	).Return(nil)

	result, err := s.service.PayOrder(
		s.ctx,
		service_order.PayOrderParams{
			OrderUUID:     orderUUID,
			PaymentMethod: model_order.PaymentMethodCard,
		},
	)

	s.NoError(err)
	s.Equal(transactionUUID, result.TransactionUUID)
}
