package service_ship_assembly

import (
	"context"

	"github.com/google/uuid"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyServiceImpl) CreateShipAssembly(
	ctx context.Context,
	params CreateShipAssemblyParams,
) (*model_ship_assembly.ShipAssembly, error) {
	assemblyUUID := uuid.Must(uuid.NewV7())

	assembly, err := s.repository.CreateShipAssembly(
		ctx,
		&model_ship_assembly.ShipAssembly{
			OrderUUID:    params.OrderUUID,
			AssemblyUUID: assemblyUUID.String(),
			Status:       model_ship_assembly.ShipAssemblyStatusNotStarted,
		},
	)
	if err != nil {
		return nil, err
	}

	return assembly, nil
}
