package service_ship_assembly

import (
	"context"

	repository_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/repository/ship_assembly"
)

var _ ShipAssemblyService = (*ShipAssemblyServiceImpl)(nil)

type ShipAssemblyServiceImpl struct {
	repository repository_ship_assembly.ShipAssemblyRepository
	cancel     context.CancelFunc
}

func NewShipAssemblyService(
	ctx context.Context,
	repository repository_ship_assembly.ShipAssemblyRepository,
) *ShipAssemblyServiceImpl {
	service := &ShipAssemblyServiceImpl{repository: repository}

	withCancel, cancel := context.WithCancel(ctx)
	service.cancel = cancel

	go service.watchOrderPaidEvent(withCancel)

	return service
}
