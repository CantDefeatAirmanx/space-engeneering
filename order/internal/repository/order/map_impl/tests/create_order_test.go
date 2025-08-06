package repository_order_map_tests

import (
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/order/internal/shared/lib/helpers/test_data"
)

func (s *TestingSuite) TestCreateOrderSuccess() {
	order := helpers_test_data.GenerateRandomOrder()
	createdOrder, err := s.repo.CreateOrder(s.ctx, *order)

	s.NoError(err)
	s.NotEmpty(createdOrder.OrderUUID)

	res, err := s.repo.GetOrder(s.ctx, createdOrder.OrderUUID)

	s.NoError(err)
	s.Equal(*createdOrder, *res)
}
