package service_order

import (
	client_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1"
	client_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1"
)

var _ OrderService = (*OrderServiceImpl)(nil)

type OrderServiceImpl struct {
	orderRepository OrderRepository
	inventoryClient client_inventory_v1.InventoryV1Client
	paymentClient   client_payment_v1.PaymentV1Client
}

type NewOrderServiceParams struct {
	OrderRepository OrderRepository
	InventoryClient client_inventory_v1.InventoryV1Client
	PaymentClient   client_payment_v1.PaymentV1Client
}

func NewOrderService(params NewOrderServiceParams) *OrderServiceImpl {
	return &OrderServiceImpl{
		orderRepository: params.OrderRepository,
		inventoryClient: params.InventoryClient,
		paymentClient:   params.PaymentClient,
	}
}
