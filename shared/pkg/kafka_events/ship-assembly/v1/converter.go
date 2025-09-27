package kafka_events_ship_assembly

import (
	assembly_events_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/events/assembly/v1"
)

func ConvertShipAssembledModelToProto(event *ShipAssembledEvent) assembly_events_v1.ShipAssembledEvent {
	return assembly_events_v1.ShipAssembledEvent{
		EventUuid:    event.EventUUID,
		AssemblyUuid: event.AssemblyUUID,
		OrderUuid:    event.OrderUUID,
		UserUuid:     event.UserUUID,
		BuildTimeSec: int64(event.BuildTimeSec),
	}
}

func ConvertShipAssembledProtoToModel(event *assembly_events_v1.ShipAssembledEvent) ShipAssembledEvent {
	return ShipAssembledEvent{
		EventUUID:    event.EventUuid,
		AssemblyUUID: event.AssemblyUuid,
		OrderUUID:    event.OrderUuid,
		UserUUID:     event.UserUuid,
		BuildTimeSec: int(event.BuildTimeSec),
	}
}
