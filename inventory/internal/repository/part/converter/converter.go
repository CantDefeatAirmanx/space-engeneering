package repository_converter_part

import (
	"github.com/samber/lo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

func ToModel(repo *repository_model_part.Part) model_part.Part {
	metadata := make(
		map[string]*model_part.MetaDataValue,
		len(repo.Metadata),
	)
	for key, value := range repo.Metadata {
		metadata[key] = &model_part.MetaDataValue{
			StringValue: value.StringValue,
			DoubleValue: value.DoubleValue,
			Int64Value:  value.Int64Value,
			BoolValue:   value.BoolValue,
		}
	}

	return model_part.Part{
		UUID:          repo.UUID,
		Name:          repo.Name,
		Description:   repo.Description,
		Price:         repo.Price,
		StockQuantity: repo.StockQuantity,
		Category:      model_part.Category(repo.Category),
		Dimensions: &model_part.Dimensions{
			Length: repo.Dimensions.Length,
			Width:  repo.Dimensions.Width,
			Height: repo.Dimensions.Height,
			Weight: repo.Dimensions.Weight,
		},
		Manufacturer: &model_part.Manufacturer{
			Name:    repo.Manufacturer.Name,
			Country: repo.Manufacturer.Country,
			Website: repo.Manufacturer.Website,
		},
		Tags:      repo.Tags,
		CreatedAt: repo.CreatedAt,
		UpdatedAt: repo.UpdatedAt,
		Metadata:  metadata,
	}
}

func ToRepository(model *model_part.Part) repository_model_part.Part {
	metadata := make(
		map[string]*repository_model_part.MetaDataValue,
		len(model.Metadata),
	)
	for key, value := range model.Metadata {
		metadata[key] = &repository_model_part.MetaDataValue{
			StringValue: value.StringValue,
			DoubleValue: value.DoubleValue,
			Int64Value:  value.Int64Value,
			BoolValue:   value.BoolValue,
		}
	}

	return repository_model_part.Part{
		UUID:          model.UUID,
		Name:          model.Name,
		Description:   model.Description,
		Price:         model.Price,
		StockQuantity: model.StockQuantity,
		Category:      repository_model_part.Category(model.Category),
		Dimensions: &repository_model_part.Dimensions{
			Length: model.Dimensions.Length,
			Width:  model.Dimensions.Width,
			Height: model.Dimensions.Height,
			Weight: model.Dimensions.Weight,
		},
		Manufacturer: &repository_model_part.Manufacturer{
			Name:    model.Manufacturer.Name,
			Country: model.Manufacturer.Country,
			Website: model.Manufacturer.Website,
		},
		Tags:      model.Tags,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Metadata:  metadata,
	}
}

func ToRepositoryFilter(model model_part.Filter) repository_model_part.Filter {
	return repository_model_part.Filter{
		Uuids: model.Uuids,
		Categories: lo.Map(
			model.Categories,
			func(category model_part.Category, _ int) repository_model_part.Category {
				return repository_model_part.Category(category)
			},
		),
		ManufacturerCountries: model.ManufacturerCountries,
		Tags:                  model.Tags,
		Names:                 model.Names,
	}
}
