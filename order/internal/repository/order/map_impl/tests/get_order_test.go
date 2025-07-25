package repository_order_map_tests

import (
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/order/internal/shared/lib/helpers/test_data"
)

func (s *TestingSuite) TestGetOrderSuccess() {
	order := helpers_test_data.GenerateRandomOrder()
	repoOrder := repository_order_converter.ToRepository(order)

	s.repo.CreateOrder(s.ctx, repoOrder)

	res, err := s.repo.GetOrder(s.ctx, repoOrder.OrderUUID)

	s.NoError(err)
	s.Equal(order.OrderUUID, res.OrderUUID)
}

func (s *TestingSuite) TestGetOrderNotFound() {
	order := helpers_test_data.GenerateRandomOrder()
	repoOrder := repository_order_converter.ToRepository(order)

	s.repo.CreateOrder(s.ctx, repoOrder)

	_, err := s.repo.GetOrder(s.ctx, "random-uuid")

	s.Error(err)
	s.ErrorIs(err, &repository_order.ErrOrderNotFound{})
}
