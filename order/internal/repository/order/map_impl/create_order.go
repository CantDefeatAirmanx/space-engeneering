package repository_order_map

import (
	"context"

	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

func (repo *OrderRepositoryMap) CreateOrder(
	ctx context.Context,
	order repository_order_model.Order,
) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.orders[order.OrderUUID] = order

	return nil
}
