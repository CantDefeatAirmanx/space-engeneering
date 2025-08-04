package repository_order_map_tests

import (
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/order/internal/shared/lib/helpers/test_data"
)

func (s *TestingSuite) TestDeleteOrderSuccess() {
	order := helpers_test_data.GenerateRandomOrder()

	s.repo.CreateOrder(s.ctx, *order)
	s.repo.DeleteOrder(s.ctx, order.OrderUUID)

	_, err := s.repo.GetOrder(s.ctx, order.OrderUUID)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderNotFound)
}
