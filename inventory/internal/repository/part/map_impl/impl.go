package repository_part_map

import (
	"context"
	"sync"

	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl/converter"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl/model"
)

var _ repository_part.PartRepository = (*RepositoryPartImpl)(nil)

type RepositoryPartImpl struct {
	mu    sync.RWMutex
	parts map[string]repository_model_part.Part
}

type NewRepositoryPartParams struct {
	InitialParts []repository_model_part.Part
}

func NewRepositoryPart(params NewRepositoryPartParams) *RepositoryPartImpl {
	repo := &RepositoryPartImpl{
		parts: make(map[string]repository_model_part.Part),
	}

	for _, part := range params.InitialParts {
		modelPart := repository_converter_part.ToModel(&part)
		err := repo.SetPart(context.Background(), &modelPart)
		if err != nil {
			panic("failed to initialize repository part")
		}
	}

	return repo
}
