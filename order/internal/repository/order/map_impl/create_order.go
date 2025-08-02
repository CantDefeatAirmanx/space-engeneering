package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
)

func (repo *OrderRepositoryMap) CreateOrder(
	ctx context.Context,
	order model_order.Order,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repoOrder := repository_order_converter.ToRepository(
		&order,
	)
	repo.orders[order.OrderUUID] = repoOrder

	return nil
}
