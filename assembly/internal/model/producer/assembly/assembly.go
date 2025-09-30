package model_producer_assembly

import (
	"context"

	kafka_events_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/kafka_events/ship-assembly/v1"
)

type ShipAssemblyProducer interface {
	ProduceAssemblyCompleted(
		ctx context.Context,
		assemblyCompletedEvent kafka_events_ship_assembly.ShipAssembledEvent,
	) error
}
