package repository_part_map

import (
	"context"
	"slices"

	"github.com/samber/lo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/converter"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

type FilterFunc func(part *repository_model_part.Part) bool

func (r *RepositoryPartImpl) GetParts(
	ctx context.Context,
	filter model_part.Filter,
) ([]*model_part.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	repositoryFilter := repository_converter_part.ToRepositoryFilter(filter)

	filterFuncs := []FilterFunc{
		createFilterByCategories(repositoryFilter.Categories),
		createFilterByManufacturerCountries(repositoryFilter.ManufacturerCountries),
		createFilterByTags(repositoryFilter.Tags),
		createFilterByUuids(repositoryFilter.Uuids),
		createFilterByNames(repositoryFilter.Names),
	}

	parts := make([]*repository_model_part.Part, 0)

outer:
	for _, part := range r.parts {
		for _, filterFunc := range filterFuncs {
			if !filterFunc(&part) {
				continue outer
			}
		}

		parts = append(parts, &part)
	}

	modelParts := convertPartsToModel(parts)

	return modelParts, nil
}

func convertPartsToModel(parts []*repository_model_part.Part) []*model_part.Part {
	return lo.Map(
		parts,
		func(part *repository_model_part.Part, _ int) *model_part.Part {
			modelPart := repository_converter_part.ToModel(part)

			return &modelPart
		},
	)
}

func createFilterByCategories(categories []repository_model_part.Category) FilterFunc {
	return func(part *repository_model_part.Part) bool {
		if len(categories) == 0 {
			return true
		}

		return slices.Contains(categories, part.Category)
	}
}

func createFilterByManufacturerCountries(manufacturerCountries []string) FilterFunc {
	return func(part *repository_model_part.Part) bool {
		if len(manufacturerCountries) == 0 {
			return true
		}

		return slices.Contains(manufacturerCountries, part.Manufacturer.Country)
	}
}

func createFilterByTags(tags []string) FilterFunc {
	return func(part *repository_model_part.Part) bool {
		if len(tags) == 0 {
			return true
		}

		for _, tag := range tags {
			if slices.Contains(part.Tags, tag) {
				return true
			}
		}

		return false
	}
}

func createFilterByUuids(uuids []string) FilterFunc {
	return func(part *repository_model_part.Part) bool {
		if len(uuids) == 0 {
			return true
		}

		return slices.Contains(uuids, part.UUID)
	}
}

func createFilterByNames(names []string) FilterFunc {
	return func(part *repository_model_part.Part) bool {
		if len(names) == 0 {
			return true
		}

		return slices.Contains(names, part.Name)
	}
}
