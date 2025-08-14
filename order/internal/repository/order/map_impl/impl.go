package repository_order_map

import (
	"sync"

	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
)

var _ repository_order.OrderRepository = (*OrderRepositoryMap)(nil)

type OrderRepositoryMap struct {
	mu     sync.RWMutex
	orders map[string]Order
}

func NewOrderRepositoryMap() *OrderRepositoryMap {
	return &OrderRepositoryMap{
		orders: make(map[string]Order),
	}
}
