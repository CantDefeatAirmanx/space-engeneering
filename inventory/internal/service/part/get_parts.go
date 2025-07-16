package service_part

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/converter/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/model/part"
	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	"github.com/samber/lo"
)

func (s *partServiceImpl) GetParts(
	ctx context.Context,
	filter Filter,
) ([]*model_part.Part, error) {
	repositoryFilter := repository_part.Filter{
		Uuids:                 filter.Uuids,
		ManufacturerCountries: filter.ManufacturerCountries,
		Tags:                  filter.Tags,
		Names:                 filter.Names,
		Categories: lo.Map(
			filter.Categories,
			func(category model_part.Category, _ int) repository_model_part.Category {
				return repository_model_part.Category(category)
			},
		),
	}
	repoParts, err := s.repository.GetParts(ctx, repositoryFilter)

	if err != nil {
		return nil, err
	}

	modelParts := lo.Map(
		repoParts,
		func(part *repository_model_part.Part, _ int) *model_part.Part {
			modelPart := repository_converter_part.ToModel(part)

			return &modelPart
		},
	)

	return modelParts, nil
}
