package api_v1

import (
	service_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

var _ assembly_v1.AssemblyServiceServer = (*Api)(nil)

type Api struct {
	assembly_v1.UnimplementedAssemblyServiceServer
	shipAssemblyService service_ship_assembly.ShipAssemblyService
}

func NewApi(shipAssemblyService service_ship_assembly.ShipAssemblyService) *Api {
	return &Api{
		shipAssemblyService: shipAssemblyService,
	}
}
