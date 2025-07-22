package repository_part_map

import (
	"sync"

	model_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part/converter"
	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/converter"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/shared/test_data"
)

var _ repository_part.PartRepository = (*repositoryPartImpl)(nil)

type repositoryPartImpl struct {
	mu    sync.RWMutex
	parts map[string]repository_model_part.Part
}

func NewRepositoryPart() *repositoryPartImpl {
	parts := make(map[string]repository_model_part.Part)

	for idx := range len(test_data.InitialParts) {
		part := &test_data.InitialParts[idx]
		modelPart := model_converter_part.ToModel(part)
		parts[modelPart.UUID] = repository_converter_part.ToRepository(&modelPart)
	}

	return &repositoryPartImpl{
		parts: parts,
	}
}
