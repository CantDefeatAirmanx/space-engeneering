package helpers_test_data

import (
	"time"

	"github.com/brianvoe/gofakeit/v7"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

func GenerateRandomPart(options ...PartOption) *model_part.Part {
	now := time.Now()
	createdAt := gofakeit.DateRange(time.Now().AddDate(0, -6, 0), now)
	updatedAt := gofakeit.DateRange(createdAt, now)

	// Generate random category
	categories := []model_part.Category{
		model_part.CategoryEngine,
		model_part.CategoryFuel,
		model_part.CategoryPortHole,
		model_part.CategoryWing,
	}
	category := categories[gofakeit.Number(0, len(categories)-1)]

	dimensions := &model_part.Dimensions{
		Length: gofakeit.Float64Range(0.1, 100.0),
		Width:  gofakeit.Float64Range(0.1, 50.0),
		Height: gofakeit.Float64Range(0.1, 20.0),
		Weight: gofakeit.Float64Range(0.01, 500.0),
	}

	// Generate manufacturer
	manufacturer := &model_part.Manufacturer{
		Name:    gofakeit.Company(),
		Country: gofakeit.Country(),
		Website: gofakeit.URL(),
	}

	// Generate tags
	tags := []string{
		gofakeit.Word(),
		gofakeit.Word(),
		gofakeit.BuzzWord(),
	}

	// Generate metadata
	metadata := map[string]*model_part.MetaDataValue{
		"color":       model_part.NewStringValue(gofakeit.Color()),
		"material":    model_part.NewStringValue(gofakeit.Word()),
		"durability":  model_part.NewDoubleValue(gofakeit.Float64Range(0.0, 10.0)),
		"warranty":    model_part.NewInt64Value(int64(gofakeit.Number(12, 120))),
		"is_premium":  model_part.NewBoolValue(gofakeit.Bool()),
		"temperature": model_part.NewDoubleValue(gofakeit.Float64Range(-50.0, 200.0)),
	}

	result := &model_part.Part{
		UUID:          gofakeit.UUID(),
		Name:          gofakeit.ProductName(),
		Description:   gofakeit.ProductDescription(),
		Price:         gofakeit.Price(1.0, 10000.0),
		StockQuantity: int64(gofakeit.Number(0, 1000)),
		Category:      category,
		Dimensions:    dimensions,
		Manufacturer:  manufacturer,
		Tags:          tags,
		CreatedAt:     &createdAt,
		UpdatedAt:     &updatedAt,
		Metadata:      metadata,
	}

	// Apply options
	for _, option := range options {
		option(result)
	}

	return result
}

type PartOption func(*model_part.Part)

func WithUUID(uuid string) PartOption {
	return func(p *model_part.Part) {
		p.UUID = uuid
	}
}

func WithName(name string) PartOption {
	return func(p *model_part.Part) {
		p.Name = name
	}
}

func WithDescription(description string) PartOption {
	return func(p *model_part.Part) {
		p.Description = description
	}
}

func WithPrice(price float64) PartOption {
	return func(p *model_part.Part) {
		p.Price = price
	}
}

func WithStockQuantity(stockQuantity int64) PartOption {
	return func(p *model_part.Part) {
		p.StockQuantity = stockQuantity
	}
}

func WithCategory(category model_part.Category) PartOption {
	return func(p *model_part.Part) {
		p.Category = category
	}
}

func WithTags(tags []string) PartOption {
	return func(p *model_part.Part) {
		p.Tags = tags
	}
}
