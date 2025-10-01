package service_ship_assembly

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyServiceImpl) GetAssemblyInfo(
	ctx context.Context,
	params GetAssemblyInfoParams,
) (*model_ship_assembly.ShipAssembly, error) {
	assembly, err := s.repository.GetShipAssembly(
		ctx,
		model_ship_assembly.SelectShipAssemblyParams{
			AssemblyUUID: params.AssemblyUUID,
			OrderUUID:    params.OrderUUID,
		},
	)
	if err != nil {
		return nil, err
	}

	return assembly, nil
}
