package repository_part_map

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

func (r *RepositoryPartImpl) GetPart(ctx context.Context, uuid string) (*repository_model_part.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[uuid]
	if !ok {
		return nil, model_part.ErrPartNotFound
	}

	return &part, nil
}
