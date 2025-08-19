package service_order_tests

import (
	"errors"

	"github.com/brianvoe/gofakeit/v7"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (s *TestingSuite) TestDeleteOrderSuccess() {
	orderUUID := gofakeit.UUID()
	s.repoMock.EXPECT().DeleteOrder(s.ctx, orderUUID).Return(nil)

	err := s.service.DeleteOrder(s.ctx, orderUUID)

	s.NoError(err)
}

func (s *TestingSuite) TestDeleteOrderNotFound() {
	orderUUID := gofakeit.UUID()
	s.repoMock.EXPECT().DeleteOrder(s.ctx, orderUUID).Return(
		model_order.ErrOrderNotFound,
	)

	err := s.service.DeleteOrder(s.ctx, orderUUID)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderNotFound)
}

func (s *TestingSuite) TestDeleteOrderRepositoryUnknownError() {
	orderUUID := gofakeit.UUID()
	s.repoMock.EXPECT().DeleteOrder(s.ctx, orderUUID).Return(
		errors.New("unknown error"),
	)

	err := s.service.DeleteOrder(s.ctx, orderUUID)

	s.Error(err)
}
