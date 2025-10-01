package repository_ship_assembly_postgres

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyRepositoryPostgres) SetShipAssemblyStatusPending(
	ctx context.Context,
	selectParams model_ship_assembly.SelectShipAssemblyParams,
) error {
	return s.UpdateShipAssembly(
		ctx,
		selectParams,
		model_ship_assembly.UpdateShipAssemblyFields{
			Status: model_ship_assembly.ShipAssemblyStatusPending,
		},
	)
}
