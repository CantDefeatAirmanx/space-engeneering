package repository_order_map

import (
	"context"

	"github.com/google/uuid"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

func (repo *OrderRepositoryMap) CreateOrder(
	ctx context.Context,
	order model_order.Order,
) (*model_order.Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	orderUUID, err := uuid.NewV7()
	if err != nil {
		return nil, model_order.ErrOrderInternal
	}

	order.OrderUUID = orderUUID.String()

	repoOrder := ToRepository(
		&order,
	)
	repo.orders[order.OrderUUID] = repoOrder

	return &order, nil
}
