package repository_part_map_tests

import (
	"github.com/brianvoe/gofakeit/v7"

	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

func generateRandomPart(opts ...PartOption) *repository_model_part.Part {
	createdAt := gofakeit.Date()
	updatedAt := gofakeit.Date()
	metaVal1 := gofakeit.RandomString([]string{"value1", "value2", "value3"})
	metaVal2 := gofakeit.RandomString([]string{"value1", "value2", "value3"})

	result := repository_model_part.Part{
		UUID:        gofakeit.UUID(),
		Name:        gofakeit.Name(),
		Description: gofakeit.Sentence(10),
		Category:    repository_model_part.Category(gofakeit.RandomInt([]int{0, 1, 2, 3})),
		Manufacturer: &repository_model_part.Manufacturer{
			Country: gofakeit.Country(),
			Name:    gofakeit.Company(),
			Website: gofakeit.URL(),
		},
		Price:         gofakeit.Float64Range(100, 10000),
		StockQuantity: gofakeit.Int64(),
		Tags:          []string{gofakeit.RandomString([]string{"tag1", "tag2", "tag3"})},
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
		Metadata: map[string]*repository_model_part.MetaDataValue{
			"key1": {
				StringValue: &metaVal1,
			},
			"key2": {
				StringValue: &metaVal2,
			},
		},
		Dimensions: &repository_model_part.Dimensions{
			Length: gofakeit.Float64Range(10, 100),
			Width:  gofakeit.Float64Range(10, 100),
			Height: gofakeit.Float64Range(10, 100),
			Weight: gofakeit.Float64Range(10, 100),
		},
	}

	for _, opt := range opts {
		opt(&result)
	}

	return &result
}

type PartOption func(part *repository_model_part.Part)

func WithName(name string) PartOption {
	return func(part *repository_model_part.Part) {
		part.Name = name
	}
}

func WithCategory(category repository_model_part.Category) PartOption {
	return func(part *repository_model_part.Part) {
		part.Category = category
	}
}

func WithUUID(uuid string) PartOption {
	return func(part *repository_model_part.Part) {
		part.UUID = uuid
	}
}

func WithTags(tags []string) PartOption {
	return func(part *repository_model_part.Part) {
		part.Tags = tags
	}
}
