package client_inventory_v1

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part"
	model_part_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part/converter"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func (client *inventoryV1GrpcClient) GetPart(
	ctx context.Context,
	params *inventory_v1.GetPartRequest,
) (*model_part.Part, error) {
	response, err := client.grpcClient.GetPart(ctx, params)
	if err != nil {
		statusErr, ok := status.FromError(err)

		if !ok {
			return nil, err
		}

		partUuid := params.Uuid
		switch statusErr.Code() {
		case codes.NotFound:
			return nil, model_part.ErrPartNotFound
		case codes.InvalidArgument:
			return nil, fmt.Errorf(
				"%w: %s, partUuid: %s",
				model_part.ErrPartInvalidArguments,
				statusErr.Message(),
				partUuid,
			)
		case codes.Internal:
			return nil, fmt.Errorf(
				"%w: %s, partUuid: %s",
				err,
				statusErr.Message(),
				partUuid,
			)
		default:
			return nil, fmt.Errorf(
				"%w: %s, partUuid: %s",
				err,
				statusErr.Message(),
				partUuid,
			)
		}
	}

	modelPart := model_part_converter.ToModel(response.Part)

	return &modelPart, nil
}
