package api_inventory_v1

import (
	service_inventory "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

var _ inventory_v1.InventoryServiceServer = (*api)(nil)

type api struct {
	inventory_v1.UnimplementedInventoryServiceServer
	partService service_inventory.PartService
}

func NewApi(
	partService service_inventory.PartService,
) *api {
	return &api{
		partService: partService,
	}
}
