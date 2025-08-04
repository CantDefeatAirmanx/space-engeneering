package repository_order_map_tests

import (
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/order/internal/shared/lib/helpers/test_data"
)

func (s *TestingSuite) TestGetOrderSuccess() {
	order := helpers_test_data.GenerateRandomOrder()

	s.repo.CreateOrder(s.ctx, *order)

	res, err := s.repo.GetOrder(s.ctx, order.OrderUUID)

	s.NoError(err)
	s.Equal(order.OrderUUID, res.OrderUUID)
}

func (s *TestingSuite) TestGetOrderNotFound() {
	order := helpers_test_data.GenerateRandomOrder()

	s.repo.CreateOrder(s.ctx, *order)

	_, err := s.repo.GetOrder(s.ctx, "random-uuid")

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderNotFound)
}
