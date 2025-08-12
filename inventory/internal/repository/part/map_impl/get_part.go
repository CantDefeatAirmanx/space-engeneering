package repository_part_map

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl/converter"
)

func (r *RepositoryPartImpl) GetPart(
	ctx context.Context,
	uuid string,
) (*model_part.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[uuid]
	if !ok {
		return nil, model_part.ErrPartNotFound
	}

	modelPart := repository_converter_part.ToModel(&part)

	return &modelPart, nil
}
