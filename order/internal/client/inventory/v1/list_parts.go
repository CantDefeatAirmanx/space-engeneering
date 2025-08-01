package client_inventory_v1

import (
	"context"
	"fmt"

	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part"
	model_part_converter "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part/converter"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func (client *inventoryV1GrpcClient) ListParts(
	ctx context.Context,
	params ListPartsParams,
) ([]*model_part.Part, error) {
	response, err := client.grpcClient.ListParts(
		ctx,
		&inventory_v1.ListPartsRequest{
			Filter: &inventory_v1.PartsFilter{
				Uuids: params.Uuids,
				Categories: lo.Map(params.Categories, func(category model_part.Category, _ int) inventory_v1.Category {
					return inventory_v1.Category(category)
				}),
				ManufacturerCountries: params.ManufacturerCountries,
				Tags:                  params.Tags,
				Names:                 params.Names,
			},
		},
	)
	if err != nil {
		statusErr, ok := status.FromError(err)

		if !ok {
			return nil, model_part.ErrPartInternal
		}

		switch statusErr.Code() {
		case codes.InvalidArgument:
			return nil, fmt.Errorf("%w: %s", model_part.ErrPartInvalidArguments, statusErr.Message())
		case codes.Internal:
			return nil, fmt.Errorf("%w: %s", model_part.ErrPartInternal, statusErr.Message())
		default:
			return nil, fmt.Errorf("%w: %s", model_part.ErrPartInternal, statusErr.Message())
		}
	}

	return lo.Map(response.Parts, func(part *inventory_v1.Part, _ int) *model_part.Part {
		modelPart := model_part_converter.ToModel(part)

		return &modelPart
	}), nil
}
