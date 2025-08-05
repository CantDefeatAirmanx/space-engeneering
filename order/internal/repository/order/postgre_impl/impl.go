package repository_order_postgre

import (
	"github.com/jackc/pgx/v5/pgxpool"

	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
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
