package api_inventory_v1

import (
	"context"

	"github.com/samber/lo"
	"go.uber.org/zap"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	model_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part/converter"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/contexts"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func (api *api) ListParts(
	ctx context.Context,
	req *inventory_v1.ListPartsRequest,
) (*inventory_v1.ListPartsResponse, error) {
	contexts.GetLogParamsSetterFunc(ctx)(
		getLogParams(req.Filter),
	)

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

func getLogParams(filter *inventory_v1.PartsFilter) []zap.Field {
	fields := []zap.Field{}

	if len(filter.Categories) > 0 {
		fields = append(
			fields,
			zap.Strings(
				categoriesLogKey,
				lo.Map(filter.Categories, func(category inventory_v1.Category, _ int) string {
					return category.String()
				}),
			),
		)
	}

	if len(filter.Uuids) > 0 {
		fields = append(fields, zap.Strings(uuidsLogKey, filter.Uuids))
	}

	if len(filter.ManufacturerCountries) > 0 {
		fields = append(fields, zap.Strings(manufacturerCountriesLogKey, filter.ManufacturerCountries))
	}

	if len(filter.Names) > 0 {
		fields = append(fields, zap.Strings(namesLogKey, filter.Names))
	}

	if len(filter.Tags) > 0 {
		fields = append(fields, zap.Strings(tagsLogKey, filter.Tags))
	}

	return fields
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
