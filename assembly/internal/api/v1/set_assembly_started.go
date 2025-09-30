package api_v1

import (
	"context"

	service_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

func (a *Api) SetAssemblyStarted(
	ctx context.Context,
	req *assembly_v1.SetAssemblyStartedRequest,
) (*assembly_v1.SetAssemblyStartedResponse, error) {
	res, err := a.shipAssemblyService.SetAssemblyStarted(
		ctx,
		service_ship_assembly.AssemblyStartedParams{
			OrderUUID:    req.OrderUuid,
			AssemblyUUID: req.AssemblyUuid,
		},
	)
	if err != nil {
		return nil, err
	}

	return &assembly_v1.SetAssemblyStartedResponse{
		AssemblyUuid: res.AssemblyUUID,
	}, nil
}
