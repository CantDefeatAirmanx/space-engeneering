package api_v1

import (
	"context"

	service_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

func (a *Api) AssemblyCompleted(
	ctx context.Context,
	req *assembly_v1.AssemblyCompletedRequest,
) (
	*assembly_v1.AssemblyCompletedResponse, error,
) {
	res, err := a.shipAssemblyService.AssemblyCompleted(
		ctx,
		service_ship_assembly.AssemblyCompletedParams{
			OrderUUID:    req.OrderUuid,
			AssemblyUUID: req.AssemblyUuid,
		},
	)
	if err != nil {
		return nil, err
	}

	return &assembly_v1.AssemblyCompletedResponse{
		AssemblyUuid: res.AssemblyUUID,
	}, nil
}
