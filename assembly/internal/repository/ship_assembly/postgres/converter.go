package repository_ship_assembly_postgres

import (
	"github.com/jackc/pgx/v5/pgtype"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func AssemblyToModel(shipAssembly ShipAssembly) *model_ship_assembly.ShipAssembly {
	res := &model_ship_assembly.ShipAssembly{}

	res.AssemblyUUID = shipAssembly.AssemblyUUID.String()
	res.OrderUUID = shipAssembly.OrderUUID.String()
	res.Status = model_ship_assembly.ShipAssemblyStatus(shipAssembly.Status)
	res.CreatedAt = shipAssembly.CreatedAt.Time
	res.UpdatedAt = shipAssembly.UpdatedAt.Time

	return res
}

func AssemblyToRepo(model *model_ship_assembly.ShipAssembly) (*ShipAssembly, error) {
	res := &ShipAssembly{}

	var orderUUID pgtype.UUID
	if err := orderUUID.Scan(model.OrderUUID); err != nil {
		return nil, err
	}
	res.OrderUUID = orderUUID

	var assemblyUUID pgtype.UUID
	if err := assemblyUUID.Scan(model.AssemblyUUID); err != nil {
		return nil, err
	}
	res.AssemblyUUID = assemblyUUID

	res.Status = ShipAssemblyStatus(model.Status)

	var createdAt pgtype.Timestamp
	if err := createdAt.Scan(model.CreatedAt); err != nil {
		return nil, err
	}
	res.CreatedAt = createdAt

	var updatedAt pgtype.Timestamp
	if err := updatedAt.Scan(model.UpdatedAt); err != nil {
		return nil, err
	}
	res.UpdatedAt = updatedAt

	return res, nil
}
