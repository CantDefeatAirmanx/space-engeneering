package api_inventory_v1

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	model_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part/converter"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func (api *api) GetPart(
	ctx context.Context,
	req *inventory_v1.GetPartRequest,
) (*inventory_v1.GetPartResponse, error) {
	part, err := api.partService.GetPart(ctx, req.Uuid)
	if err != nil {
		switch {
		case errors.Is(err, &model_part.ErrPartNotFound{}):
			return nil, status.Errorf(codes.NotFound, "Part %s is not found. %v", req.Uuid, err)
		default:
			return nil, status.Errorf(codes.Internal, "Internal server error. %v", err)
		}
	}

	protoPart := model_converter_part.ToProto(part)

	return &inventory_v1.GetPartResponse{
		Part: protoPart,
	}, nil
}
