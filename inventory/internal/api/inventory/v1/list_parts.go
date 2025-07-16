package api_inventory_v1

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	model_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part/converter"
	service_inventory "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
	"github.com/samber/lo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (api *api) ListParts(
	ctx context.Context,
	req *inventory_v1.ListPartsRequest,
) (*inventory_v1.ListPartsResponse, error) {
	categories := lo.Map(
		req.Filter.Categories,
		func(category inventory_v1.Category, _ int) model_part.Category {
			return model_part.Category(category)
		},
	)

	parts, err := api.partService.GetParts(ctx, service_inventory.Filter{
		Uuids:                 req.Filter.Uuids,
		ManufacturerCountries: req.Filter.ManufacturerCountries,
		Tags:                  req.Filter.Tags,
		Names:                 req.Filter.Names,
		Categories:            categories,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error. %v", err)
	}

	protoParts := lo.Map(parts, func(part *model_part.Part, _ int) *inventory_v1.Part {
		protoPart := model_converter_part.ToProto(part)

		return &protoPart
	})

	return &inventory_v1.ListPartsResponse{
		Parts: protoParts,
	}, nil
}
