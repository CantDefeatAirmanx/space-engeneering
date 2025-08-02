package api_inventory_v1

import (
	"context"

	"github.com/samber/lo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	model_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part/converter"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func (api *api) ListParts(
	ctx context.Context,
	req *inventory_v1.ListPartsRequest,
) (*inventory_v1.ListPartsResponse, error) {
	categories := categoriesToModel(req.Filter.Categories)

	parts, err := api.partService.GetParts(ctx, model_part.Filter{
		Uuids:                 req.Filter.Uuids,
		ManufacturerCountries: req.Filter.ManufacturerCountries,
		Tags:                  req.Filter.Tags,
		Names:                 req.Filter.Names,
		Categories:            categories,
	})
	if err != nil {
		return nil, err
	}

	protoParts := partsToProtoParts(parts)

	return &inventory_v1.ListPartsResponse{
		Parts: protoParts,
	}, nil
}

func categoriesToModel(categories []inventory_v1.Category) []model_part.Category {
	return lo.Map(
		categories,
		func(category inventory_v1.Category, _ int) model_part.Category {
			return model_part.Category(category)
		},
	)
}

func partsToProtoParts(parts []*model_part.Part) []*inventory_v1.Part {
	return lo.Map(parts, func(part *model_part.Part, _ int) *inventory_v1.Part {
		protoPart := model_converter_part.ToProto(part)

		return protoPart
	})
}
