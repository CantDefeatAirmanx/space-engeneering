package service_order

import (
	"context"

	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	kafka_events_order "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/order/v1"
)

type OrderService interface {
	// CreateOrder creates an order in the repository.
	//
	// Errors:
	//
	// - ([model_order.ErrOrderInternal]): if the order is not created
	// - ([model_order.ErrOrderInvalidArguments]): if payload is not valid
	CreateOrder(ctx context.Context, params CreateOrderParams) (*CreateOrderResult, error)
	// GetOrder returns an order from the repository.
	//
	// Errors:
	//
	// - ([model_order.ErrOrderNotFound]): if the order is not found
	//
	// - ([model_order.ErrOrderInternal]): if repository returns an other error
	GetOrder(ctx context.Context, orderUUID string) (*model_order.Order, error)
	// UpdateOrder updates an order in the repository.
	//
	// Errors:
	//
	// - ([model_order.ErrOrderNotFound]): if the order is not found
	//
	// - ([model_order.ErrOrderInternal]): if the order is not updated
	UpdateOrder(ctx context.Context, orderUUID string, update UpdateOrderFields) error
	// DeleteOrder deletes an order from the repository.
	//
	// Errors:
	//
	// - ([model_order.ErrOrderNotFound]): if the order is not found
	//
	// - ([model_order.ErrOrderInternal]): if the order is not deleted
	DeleteOrder(ctx context.Context, orderUUID string) error
	// CancelOrder cancels an order in the repository.
	//
	// Errors:
	//
	// - ([model_order.ErrOrderNotFound]): if the order is not found
	//
	// - ([model_order.ErrOrderConflict]): if the order is not cancelable
	//
	// - ([model_order.ErrOrderInternal]): if the order is not canceled
	CancelOrder(ctx context.Context, orderUUID string) error
	// PayOrder pays an order in the repository.
	//
	// Errors:
	//
	// - ([model_order.ErrOrderNotFound]): if the order is not found
	//
	// - ([model_order.ErrOrderConflict]): if the order is not available for payment
	//
	// - ([model_order.ErrOrderInternal]): if the order payment fails
	PayOrder(ctx context.Context, params PayOrderParams) (*PayOrderResult, error)
}

type OrderProducer interface {
	ProduceOrderPaid(ctx context.Context, order kafka_events_order.OrderPaidEvent) error
}

type CreateOrderParams struct {
	UserUUID  string
	PartUuids []string
}

type CreateOrderResult struct {
	OrderUUID  string
	TotalPrice float64
}

type PayOrderParams struct {
	OrderUUID     string
	PaymentMethod model_order.PaymentMethod
}

type PayOrderResult struct {
	TransactionUUID string
}

type UpdateOrderFields struct {
	Status          *model_order.OrderStatus
	TransactionUUID *string
	PaymentMethod   *model_order.PaymentMethod
}
