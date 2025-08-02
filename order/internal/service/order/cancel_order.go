package service_order

import (
	"context"
	"fmt"
	"slices"

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

type conflictStatus struct {
	Value         model_order.OrderStatus
	GetErrMessage func(orderUUID string) string
}

var conflictStatuses = []conflictStatus{
	{
		Value: model_order.OrderStatusPaid,
		GetErrMessage: func(orderUUID string) string {
			return fmt.Sprintf("Order %s already paid", orderUUID)
		},
	},
	{
		Value: model_order.OrderStatusCancelled,
		GetErrMessage: func(orderUUID string) string {
			return fmt.Sprintf("Order %s already cancelled", orderUUID)
		},
	},
}

func getIsOrderAvailableForCancel(order *model_order.Order) (isAvailable bool, errMessage string) {
	conflictIdx := slices.IndexFunc(conflictStatuses, func(c conflictStatus) bool {
		return c.Value == order.Status
	})

	if conflictIdx != -1 {
		conflictObj := conflictStatuses[conflictIdx]
		errMessage := conflictObj.GetErrMessage(order.OrderUUID)

		return false, errMessage
	}

	return true, ""
}
