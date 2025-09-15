package service_order_tests

import (
	"errors"

	"github.com/brianvoe/gofakeit/v7"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (s *TestingSuite) TestGetOrderSuccess() {
	orderUUID := gofakeit.UUID()
	modelOrder := &model_order.Order{
		OrderUUID: orderUUID,
	}

	s.repoMock.EXPECT().GetOrder(s.ctx, orderUUID).Return(
		modelOrder,
		nil,
	)

	result, err := s.service.GetOrder(s.ctx, orderUUID)

	s.NoError(err)
	s.Equal(modelOrder, result)
}

func (s *TestingSuite) TestGetOrderNotFound() {
	orderUUID := gofakeit.UUID()
	s.repoMock.EXPECT().GetOrder(s.ctx, orderUUID).Return(
		nil,
		model_order.ErrOrderNotFound,
	)

	result, err := s.service.GetOrder(s.ctx, orderUUID)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderNotFound)
	s.Nil(result)
}

func (s *TestingSuite) TestGetOrderRepositoryUnknownError() {
	orderUUID := gofakeit.UUID()
	s.repoMock.EXPECT().GetOrder(s.ctx, orderUUID).Return(
		nil,
		errors.New("unknown error"),
	)

	result, err := s.service.GetOrder(s.ctx, orderUUID)

	s.Error(err)
	s.Nil(result)
}
