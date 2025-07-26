package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
)

func (repo *OrderRepositoryMap) UpdateOrderFields(
	ctx context.Context,
	orderUUID string,
	update repository_order.UpdateOrderFields,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	order, ok := repo.orders[orderUUID]
	if !ok {
		return &model_order.ErrOrderNotFound{
			OrderUUID: orderUUID,
		}
	}

	if update.Status != nil {
		order.Status = *update.Status
	}

	if update.TransactionUUID != nil {
		order.TransactionUUID = update.TransactionUUID
	}

	if update.PaymentMethod != nil {
		order.PaymentMethod = update.PaymentMethod
	}

	repo.orders[orderUUID] = order

	return nil
}
