package repository_part_mongo

import (
	"github.com/samber/lo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

func ToModel(part *Part) model_part.Part {
	result := model_part.Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      model_part.Category(part.Category),
		Dimensions: &model_part.Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: &model_part.Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
		Metadata:  map[string]*model_part.MetaDataValue{},
	}

	for key, value := range part.Metadata {
		result.Metadata[key] = &model_part.MetaDataValue{
			StringValue: value.StringValue,
			DoubleValue: value.DoubleValue,
			Int64Value:  value.Int64Value,
			BoolValue:   value.BoolValue,
		}
	}

	return result
}

func ToRepo(part *model_part.Part) Part {
	result := Part{
		UUID:          part.UUID,
		Name:          part.Name,
		Description:   part.Description,
		Price:         part.Price,
		StockQuantity: part.StockQuantity,
		Category:      Category(part.Category),
		Dimensions: &Dimensions{
			Length: part.Dimensions.Length,
			Width:  part.Dimensions.Width,
			Height: part.Dimensions.Height,
			Weight: part.Dimensions.Weight,
		},
		Manufacturer: &Manufacturer{
			Name:    part.Manufacturer.Name,
			Country: part.Manufacturer.Country,
			Website: part.Manufacturer.Website,
		},
		Tags:      part.Tags,
		CreatedAt: part.CreatedAt,
		UpdatedAt: part.UpdatedAt,
		Metadata:  map[string]*MetaDataValue{},
	}

	for key, value := range part.Metadata {
		result.Metadata[key] = &MetaDataValue{
			StringValue: value.StringValue,
			DoubleValue: value.DoubleValue,
			Int64Value:  value.Int64Value,
			BoolValue:   value.BoolValue,
		}
	}

	return result
}

func ToRepoFilter(filter *model_part.Filter) Filter {
	result := Filter{}

	if filter.Uuids != nil {
		result.Uuids = filter.Uuids
	}

	if filter.Categories != nil {
		result.Categories = lo.Map(
			filter.Categories, func(category model_part.Category, _ int) Category {
				return Category(category)
			},
		)
	}

	if filter.ManufacturerCountries != nil {
		result.ManufacturerCountries = filter.ManufacturerCountries
	}

	if filter.Tags != nil {
		result.Tags = filter.Tags
	}

	if filter.Names != nil {
		result.Names = filter.Names
	}

	return result
}
