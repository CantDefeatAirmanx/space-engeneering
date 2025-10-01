package model_consumer_order

import (
	"context"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

type OrderConsumer interface {
	WatchOrderPaidEvent(ctx context.Context)
	SetOrderPaidProcessor(orderPaidProcessor WithProcessOrderPaidEvent)
	interfaces.WithClose
}

type WithProcessOrderPaidEvent interface {
	ProcessOrderPaidEvent(ctx context.Context, orderPaidEvent *order_events_v1.OrderPaidEvent) error
}
