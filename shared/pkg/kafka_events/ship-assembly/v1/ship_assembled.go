package kafka_events_ship_assembly

type ShipAssembledEvent struct {
	AssemblyUUID string
	EventUUID    string
	OrderUUID    string
	UserUUID     string
	BuildTimeSec int
}
