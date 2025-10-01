package api_v1

import (
	"context"

	model_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/model/ship_assembly"
	service_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

func (a *Api) GetAssemblyInfo(
	ctx context.Context,
	req *assembly_v1.GetAssemblyInfoRequest,
) (*assembly_v1.GetAssemblyInfoResponse, error) {
	res, err := a.shipAssemblyService.GetAssemblyInfo(
		ctx,
		service_ship_assembly.GetAssemblyInfoParams{
			AssemblyUUID: req.AssemblyUuid,
			OrderUUID:    req.OrderUuid,
		},
	)
	if err != nil {
		return nil, err
	}

	return &assembly_v1.GetAssemblyInfoResponse{
		AssemblyInfo: model_ship_assembly.ConvertToProto(*res),
	}, nil
}
