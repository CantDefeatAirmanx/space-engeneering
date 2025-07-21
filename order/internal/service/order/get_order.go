package service_order

import (
	"context"
	"errors"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
)

func (s *OrderServiceImpl) GetOrder(ctx context.Context, orderUUID string) (*model_order.Order, error) {
	order, err := s.orderRepository.GetOrder(ctx, orderUUID)
	if err != nil {
		if errors.Is(err, &repository_order.ErrOrderNotFound{}) {
			return nil, &model_order.ErrOrderNotFound{
				OrderUUID: orderUUID,
				Err:       err,
			}
		}

		return nil, &model_order.ErrOrderInternal{
			OrderUUID: orderUUID,
			Err:       err,
		}
	}

	orderModel := repository_order_converter.ToModel(order)

	return &orderModel, nil
}
