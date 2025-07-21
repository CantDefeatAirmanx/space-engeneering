package repository_order_map

import (
	"context"

	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
)

func (repo *OrderRepositoryMap) DeleteOrder(
	ctx context.Context,
	orderUUID string,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	_, ok := repo.orders[orderUUID]
	if !ok {
		return &repository_order.ErrOrderNotFound{
			OrderID: orderUUID,
			Err:     repository_order.ErrOrder,
		}
	}

	delete(repo.orders, orderUUID)

	return nil
}
