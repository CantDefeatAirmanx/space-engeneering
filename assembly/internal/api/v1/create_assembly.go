package api_v1

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
	service_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

func (a *Api) CreateAssembly(
	ctx context.Context,
	req *assembly_v1.CreateAssemblyRequest,
) (*assembly_v1.CreateAssemblyResponse, error) {
	res, err := a.shipAssemblyService.CreateShipAssembly(
		ctx,
		service_ship_assembly.CreateShipAssemblyParams{
			OrderUUID: req.OrderUuid,
		},
	)
	if err != nil {
		return nil, err
	}

	return &assembly_v1.CreateAssemblyResponse{
		AssemblyInfo: model_ship_assembly.ConvertToProto(*res),
	}, nil
}
