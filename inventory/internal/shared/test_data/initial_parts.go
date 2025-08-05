package test_data

import (
	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/converter"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

func GetRepoInitialParts() []repository_model_part.Part {
	parts := []repository_model_part.Part{}
	for _, data := range PartsData {
		randomPart := helpers_test_data.GenerateRandomPart(
			helpers_test_data.WithUUID(data.UUID),
			helpers_test_data.WithName(data.Name),
			helpers_test_data.WithCategory(data.Category),
		)
		parts = append(parts, repository_converter_part.ToRepository(randomPart))
	}
	return parts
}

var PartsData = []model_part.Part{
	{
		UUID:     "ca8da3f5-64b9-4958-9894-c28ce39d561a",
		Name:     "engine-01",
		Category: model_part.CategoryEngine,
	},
	{
		UUID:     "c6287976-02ad-4c4d-a196-d7b13eaf30b0",
		Name:     "engine-02",
		Category: model_part.CategoryEngine,
	},
	{
		UUID:     "bdf11c57-5c1a-4eaa-ae70-b2502b55923b",
		Name:     "fuel-01",
		Category: model_part.CategoryFuel,
	},
	{
		UUID:     "6c3ca46a-ab5c-4766-817f-c871c8cd2422",
		Name:     "fuel-02",
		Category: model_part.CategoryFuel,
	},
	{
		UUID:     "5020b49b-53f2-4454-acb8-0371becbab7e",
		Name:     "port-hole-01",
		Category: model_part.CategoryPortHole,
	},
	{
		UUID:     "7f6ed7d0-e517-4cbf-b814-12a301c3fe5e",
		Name:     "wing-01",
		Category: model_part.CategoryWing,
	},
}
