package repository_part

import (
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/converter/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/shared/test_data"
	"sync"

	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/converter/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/model/part"
)

var _ PartRepository = (*repositoryPartImpl)(nil)

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
