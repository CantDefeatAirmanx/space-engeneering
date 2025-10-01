package repository_ship_assembly_postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type ShipAssembly struct {
	AssemblyUUID pgtype.UUID        `db:"assembly_uuid"`
	OrderUUID    pgtype.UUID        `db:"order_uuid"`
	Status       ShipAssemblyStatus `db:"status"`

	CreatedAt pgtype.Timestamp `db:"created_at"`
	UpdatedAt pgtype.Timestamp `db:"updated_at"`
}

type ShipAssemblyUpdateFields struct {
	Status    ShipAssemblyStatus `db:"status"`
	StartTime *pgtype.Timestamp  `db:"assembly_start_time"`
}

type ShipAssemblyStatus string

const (
	ShipAssemblyStatusUnspecified ShipAssemblyStatus = "UNSPECIFIED"
	ShipAssemblyStatusNotStarted  ShipAssemblyStatus = "NOT_STARTED"
	ShipAssemblyStatusPending     ShipAssemblyStatus = "PENDING"
	ShipAssemblyStatusCompleted   ShipAssemblyStatus = "COMPLETED"
)
