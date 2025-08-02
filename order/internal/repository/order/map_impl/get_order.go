package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/converter"
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

	modelOrder := repository_order_converter.ToModel(&order)

	return &modelOrder, nil
}
