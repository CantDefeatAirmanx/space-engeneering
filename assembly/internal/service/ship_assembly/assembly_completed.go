package service_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyServiceImpl) AssemblyCompleted(
	ctx context.Context,
	params AssemblyCompletedParams,
) (*AssemblyCompletedReturn, error) {
	err := s.repository.UpdateShipAssembly(
		ctx,

		model_ship_assembly.SelectShipAssemblyParams{
			AssemblyUUID: params.AssemblyUUID,
			OrderUUID:    params.OrderUUID,
		},

		model_ship_assembly.UpdateShipAssemblyFields{
			Status: model_ship_assembly.ShipAssemblyStatusCompleted,
		},
	)
	if err != nil {
		return nil, err
	}

	return &AssemblyCompletedReturn{
		AssemblyUUID: params.AssemblyUUID,
	}, nil
}
