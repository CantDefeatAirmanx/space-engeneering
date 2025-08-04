package service_order

import (
	"context"
)

func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, orderUUID string) error {
	err := s.orderRepository.DeleteOrder(ctx, orderUUID)
	if err != nil {
		return err
	}

	return nil
}
