package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (repo *OrderRepositoryMap) DeleteOrder(
	ctx context.Context,
	orderUUID string,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.orders[orderUUID]
	if !ok {
		return &model_order.ErrOrderNotFound{
			OrderUUID: orderUUID,
		}
	}

	delete(repo.orders, orderUUID)

	return nil
}
