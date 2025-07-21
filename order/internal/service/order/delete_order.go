package service_order

import (
	"context"
	"errors"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
)

func (s *OrderServiceImpl) DeleteOrder(ctx context.Context, orderUUID string) error {
	err := s.orderRepository.DeleteOrder(ctx, orderUUID)
	if err != nil {
		if errors.Is(err, &repository_order.ErrOrderNotFound{}) {
			return &model_order.ErrOrderNotFound{
				OrderUUID: orderUUID,
				Err:       err,
			}
		}

		return &model_order.ErrOrderInternal{
			OrderUUID: orderUUID,
			Err:       err,
		}
	}

	return nil
}
