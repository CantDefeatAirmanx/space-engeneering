package service_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyServiceImpl) AssemblyStarted(
	ctx context.Context,
	params AssemblyStartedParams,
) (*AssemblyStartedReturn, error) {
	err := s.repository.UpdateShipAssembly(
		ctx,

		model_ship_assembly.SelectShipAssemblyParams{
			AssemblyUUID: params.AssemblyUUID,
			OrderUUID:    params.OrderUUID,
		},

		model_ship_assembly.UpdateShipAssemblyFields{
			Status: model_ship_assembly.ShipAssemblyStatusPending,
		},
	)
	if err != nil {
		return nil, err
	}

	return &AssemblyStartedReturn{
		AssemblyUUID: params.AssemblyUUID,
	}, nil
}
