package service_ship_assembly

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
)

func (s *ShipAssemblyServiceImpl) CreateShipAssembly(
	ctx context.Context,
	params CreateShipAssemblyParams,
) (*model_ship_assembly.ShipAssembly, error) {
	existingAssembly, err := s.repository.GetShipAssembly(
		ctx,
		model_ship_assembly.SelectShipAssemblyParams{
			OrderUUID: params.OrderUUID,
		},
	)
	if err != nil && !errors.Is(err, model_ship_assembly.ErrAssemblyNotFound) {
		return nil, err
	}
	if existingAssembly != nil {
		return nil, fmt.Errorf("%w: assembly already exists", model_ship_assembly.ErrAssemblyConflict)
	}

	assemblyUUID := uuid.Must(uuid.NewV7())

	err = s.repository.CreateShipAssembly(ctx, &model_ship_assembly.ShipAssembly{
		OrderUUID:    params.OrderUUID,
		AssemblyUUID: assemblyUUID.String(),
		Status:       model_ship_assembly.ShipAssemblyStatusNotStarted,
	})
	if err != nil {
		return nil, err
	}

	return &model_ship_assembly.ShipAssembly{
		OrderUUID: params.OrderUUID,
	}, nil
}
