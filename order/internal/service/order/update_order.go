package service_order

import (
	"context"

	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

func (s *OrderServiceImpl) UpdateOrder(ctx context.Context, orderUUID string, update UpdateOrderFields) error {
	repositoryUpdate := toRepositoryUpdate(update)

	err := s.orderRepository.UpdateOrderFields(
		ctx,
		orderUUID,
		repositoryUpdate,
	)

	return err
}

func toRepositoryUpdate(modelUpdate UpdateOrderFields) repository_order.UpdateOrderFields {
	result := repository_order.UpdateOrderFields{
		TransactionUUID: modelUpdate.TransactionUUID,
	}

	if modelUpdate.Status != nil {
		repoStatus := repository_order_model.OrderStatus(
			*modelUpdate.Status,
		)
		result.Status = &repoStatus
	}

	if modelUpdate.PaymentMethod != nil {
		repoPaymentMethod := repository_order_model.PaymentMethod(
			*modelUpdate.PaymentMethod,
		)
		result.PaymentMethod = &repoPaymentMethod
	}

	return result
}
