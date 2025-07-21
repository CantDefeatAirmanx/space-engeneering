package service_order

import (
	"context"
	"errors"
	"fmt"
	"slices"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

func (s *OrderServiceImpl) CancelOrder(ctx context.Context, orderUUID string) error {
	order, err := s.orderRepository.GetOrder(ctx, orderUUID)
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
	orderModel := repository_order_converter.ToModel(order)

	isOrderAvailableForCancel, errMessage := getIsOrderAvailableForCancel(orderModel)
	if !isOrderAvailableForCancel {
		return &model_order.ErrOrderConflict{
			OrderUUID:  orderUUID,
			ErrMessage: errMessage,
		}
	}

	repoStatus := repository_order_model.OrderStatus(
		model_order.OrderStatusCancelled,
	)
	err = s.orderRepository.UpdateOrderFields(ctx, orderUUID, repository_order.UpdateOrderFields{
		Status: &repoStatus,
	})
	if err != nil {
		return &model_order.ErrOrderInternal{
			OrderUUID: orderUUID,
			Err:       err,
		}
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

func getIsOrderAvailableForCancel(order model_order.Order) (isAvailable bool, errMessage string) {
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
