package repository_order_map_tests

import (
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/order/internal/shared/lib/helpers/test_data"
)

func (s *TestingSuite) TestDeleteOrderSuccess() {
	order := helpers_test_data.GenerateRandomOrder()
	repoOrder := repository_order_converter.ToRepository(order)

	s.repo.CreateOrder(s.ctx, repoOrder)
	s.repo.DeleteOrder(s.ctx, repoOrder.OrderUUID)

	_, err := s.repo.GetOrder(s.ctx, repoOrder.OrderUUID)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderNotFound)
}
