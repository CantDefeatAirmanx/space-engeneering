package model_ship_assembly

import "time"

type ShipAssembly struct {
	AssemblyUUID string
	OrderUUID    string

	Status ShipAssemblyStatus

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SelectShipAssemblyParams struct {
	AssemblyUUID string
	OrderUUID    string
}

type UpdateShipAssemblyFields struct {
	Status ShipAssemblyStatus
}

type ShipAssemblyStatus string

const (
	ShipAssemblyStatusUnspecified ShipAssemblyStatus = "UNSPECIFIED"
	ShipAssemblyStatusNotStarted  ShipAssemblyStatus = "NOT_STARTED"
	ShipAssemblyStatusPending     ShipAssemblyStatus = "PENDING"
	ShipAssemblyStatusCompleted   ShipAssemblyStatus = "COMPLETED"
)
