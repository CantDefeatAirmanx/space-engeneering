package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (repo *OrderRepositoryMap) CreateOrder(
	ctx context.Context,
	order model_order.Order,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repoOrder := ToRepository(
		&order,
	)
	repo.orders[order.OrderUUID] = repoOrder

	return nil
}
