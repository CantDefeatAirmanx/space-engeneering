package service_order

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (s *OrderServiceImpl) GetOrder(ctx context.Context, orderUUID string) (*model_order.Order, error) {
	order, err := s.orderRepository.GetOrder(ctx, orderUUID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
