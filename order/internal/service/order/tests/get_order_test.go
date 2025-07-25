package service_order_tests

import (
	"errors"

	"github.com/brianvoe/gofakeit/v7"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

func (s *TestingSuite) TestGetOrderSuccess() {
	orderUUID := gofakeit.UUID()
	repoOrder := repository_order_model.Order{
		OrderUUID: orderUUID,
	}
	modelOrder := repository_order_converter.ToModel(&repoOrder)

	s.repoMock.EXPECT().GetOrder(s.ctx, orderUUID).Return(
		&repoOrder,
		nil,
	)

	result, err := s.service.GetOrder(s.ctx, orderUUID)

	s.NoError(err)
	s.Equal(modelOrder, *result)
}

func (s *TestingSuite) TestGetOrderNotFound() {
	orderUUID := gofakeit.UUID()
	s.repoMock.EXPECT().GetOrder(s.ctx, orderUUID).Return(
		nil,
		&repository_order.ErrOrderNotFound{OrderID: orderUUID},
	)

	result, err := s.service.GetOrder(s.ctx, orderUUID)

	s.Error(err)
	s.ErrorIs(err, &model_order.ErrOrderNotFound{})
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
	s.ErrorIs(err, &model_order.ErrOrderInternal{})
	s.Nil(result)
}
