package service_order

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
)

func (s *OrderServiceImpl) GetOrder(ctx context.Context, orderUUID string) (*model_order.Order, error) {
	order, err := s.orderRepository.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, err
	}

	orderModel := repository_order_converter.ToModel(order)

	return &orderModel, nil
}
