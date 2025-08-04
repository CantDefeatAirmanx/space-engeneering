package service_order

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (s *OrderServiceImpl) UpdateOrder(
	ctx context.Context,
	orderUUID string,
	update UpdateOrderFields,
) error {
	err := s.orderRepository.UpdateOrderFields(
		ctx,
		orderUUID,
		model_order.UpdateOrderFields(update),
	)

	return err
}
