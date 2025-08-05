package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (repo *OrderRepositoryMap) UpdateOrderFields(
	ctx context.Context,
	orderUUID string,
	update model_order.UpdateOrderFields,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	order, ok := repo.orders[orderUUID]
	if !ok {
		return model_order.ErrOrderNotFound
	}

	if update.Status != nil {
		repoStatus := OrderStatus(
			*update.Status,
		)
		order.Status = repoStatus
	}

	if update.TransactionUUID != nil {
		order.TransactionUUID = update.TransactionUUID
	}

	if update.PaymentMethod != nil {
		repoPaymentMethod := PaymentMethod(
			*update.PaymentMethod,
		)
		order.PaymentMethod = &repoPaymentMethod
	}

	repo.orders[orderUUID] = order

	return nil
}
