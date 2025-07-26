package repository_order

import (
	"context"

	repository_order_model "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/model"
)

type OrderRepository interface {
	// CreateOrder creates an order in the repository.
	// Errors:
	//
	// - [model_order.ErrOrderInternal]: if the order is not created
	CreateOrder(ctx context.Context, order repository_order_model.Order) error
	// GetOrder returns an order from the repository.
	//
	// Errors:
	//
	// - [model_order.ErrOrderNotFound]: if the order is not found
	GetOrder(ctx context.Context, orderUUID string) (*repository_order_model.Order, error)
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
	UpdateOrderFields(ctx context.Context, orderUUID string, update UpdateOrderFields) error
}

type UpdateOrderFields struct {
	Status          *repository_order_model.OrderStatus
	TransactionUUID *string
	PaymentMethod   *repository_order_model.PaymentMethod
}
