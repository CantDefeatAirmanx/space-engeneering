package service_ship_assembly

import (
	"context"

	order_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/order/v1"
)

func (s *ShipAssemblyServiceImpl) ProcessOrderPaidEvent(
	ctx context.Context,
	orderPaidEvent *order_events_v1.OrderPaidEvent,
) error {
	return s.createAutomaticAssembly(ctx, orderPaidEvent)
}
