package service_order

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
)

type OrderRepository interface {
	// CreateOrder creates an order in the repository.
	// Errors:
	//
	// - [model_order.ErrOrderInternal]: if the order is not created
	CreateOrder(ctx context.Context, order model_order.Order) error
	// GetOrder returns an order from the repository.
	//
	// Errors:
	//
	// - [model_order.ErrOrderNotFound]: if the order is not found
	GetOrder(ctx context.Context, orderUUID string) (*model_order.Order, error)
	// DeleteOrder deletes an order from the repository.
	//
	// Errors:
	//
	// - [model_order.ErrOrderNotFound]: if the order is not found
	DeleteOrder(ctx context.Context, orderUUID string) error
	// UpdateOrderFields updates the fields of an order in the repository.
	//
	// Errors:
	//
	// - [model_order.ErrOrderNotFound]: if the order is not found
	UpdateOrderFields(ctx context.Context, orderUUID string, update model_order.UpdateOrderFields) error
}
