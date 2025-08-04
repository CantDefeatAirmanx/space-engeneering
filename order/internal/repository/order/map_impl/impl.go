package repository_order_map

import (
	"sync"

	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

var _ repository_order.OrderRepository = (*OrderRepositoryMap)(nil)

type OrderRepositoryMap struct {
	mu     sync.RWMutex
	orders map[string]repository_order_model.Order
}

func NewOrderRepositoryMap() *OrderRepositoryMap {
	return &OrderRepositoryMap{
		orders: make(map[string]repository_order_model.Order),
	}
}
