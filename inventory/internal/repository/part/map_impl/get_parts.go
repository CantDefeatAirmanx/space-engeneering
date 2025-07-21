package repository_part_map

import (
	"context"
	"slices"

	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

type FilterFunc func(part *repository_model_part.Part) bool

func (r *repositoryPartImpl) GetParts(ctx context.Context, filter repository_part.Filter) ([]*repository_model_part.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	filterFuncs := []FilterFunc{
		createFilterByCategories(filter.Categories),
		createFilterByManufacturerCountries(filter.ManufacturerCountries),
		createFilterByTags(filter.Tags),
		createFilterByUuids(filter.Uuids),
		createFilterByNames(filter.Names),
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

	return parts, nil
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
