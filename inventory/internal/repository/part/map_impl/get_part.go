package repository_part_map

import (
	"context"

	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

func (r *repositoryPartImpl) GetPart(ctx context.Context, uuid string) (*repository_model_part.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[uuid]
	if !ok {
		return nil, &repository_part.ErrPartNotFound{
			UUID: part.UUID,
		}
	}

	return &part, nil
}
