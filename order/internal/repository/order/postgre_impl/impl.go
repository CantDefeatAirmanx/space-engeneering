package repository_order_postgre

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository_order.OrderRepository = (*OrderRepositoryPostgre)(nil)

type OrderRepositoryPostgre struct {
	db *pgxpool.Pool
}

func NewOrderRepositoryPostgre(
	db *pgxpool.Pool,
) *OrderRepositoryPostgre {
	return &OrderRepositoryPostgre{db: db}
}

func (o *OrderRepositoryPostgre) CreateOrder(
	ctx context.Context,
	order model_order.Order,
) error {
	panic("unimplemented")
}

func (o *OrderRepositoryPostgre) DeleteOrder(
	ctx context.Context,
	orderUUID string,
) error {
	panic("unimplemented")
}

func (o *OrderRepositoryPostgre) GetOrder(
	ctx context.Context,
	orderUUID string,
) (*model_order.Order, error) {
	panic("unimplemented")
}

func (o *OrderRepositoryPostgre) UpdateOrderFields(
	ctx context.Context,
	orderUUID string,
	update model_order.UpdateOrderFields,
) error {
	panic("unimplemented")
}
