package repository_order_map

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

func (repo *OrderRepositoryMap) GetOrder(
	ctx context.Context,
	orderUUID string,
) (*repository_order_model.Order, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	order, ok := repo.orders[orderUUID]
	if !ok {
		return nil, &model_order.ErrOrderNotFound{
			OrderUUID: orderUUID,
		}
	}

	return &order, nil
}
