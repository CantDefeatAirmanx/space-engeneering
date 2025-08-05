package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (repo *OrderRepositoryMap) GetOrder(
	ctx context.Context,
	orderUUID string,
) (*model_order.Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	order, ok := repo.orders[orderUUID]
	if !ok {
		return nil, model_order.ErrOrderNotFound
	}

	modelOrder := ToModel(&order)

	return &modelOrder, nil
}
