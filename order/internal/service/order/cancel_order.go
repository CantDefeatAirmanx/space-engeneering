package service_order

import (
	"context"
	"fmt"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (s *OrderServiceImpl) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.orderRepository.GetOrder(ctx, orderUUID)
	if err != nil {
		return err
	}

	isOrderAvailableForCancel, errMessage := getIsOrderAvailableForCancel(order)

	if !isOrderAvailableForCancel {
		return fmt.Errorf("%w: %s", model_order.ErrOrderConflict, errMessage)
	}

	canceledStatus := model_order.OrderStatusCancelled
	err = s.orderRepository.UpdateOrderFields(ctx, orderUUID, model_order.UpdateOrderFields{
		Status: &canceledStatus,
	})
	if err != nil {
		return model_order.ErrOrderInternal
	}

	return nil
}

func getIsOrderAvailableForCancel(order *model_order.Order) (isAvailable bool, errMessage string) {
	switch order.Status {
	case model_order.OrderStatusPaid:
		return false, "Order already paid"
	case model_order.OrderStatusCancelled:
		return false, "Order already cancelled"
	default:
		return true, ""
	}
}
