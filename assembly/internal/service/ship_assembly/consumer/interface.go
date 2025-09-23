package service_ship_assembly_consumer

import (
	"context"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
)

type ShipAssemblyConsumer interface {
	WatchOrderPaidEvent(ctx context.Context)
	interfaces.WithClose
}
