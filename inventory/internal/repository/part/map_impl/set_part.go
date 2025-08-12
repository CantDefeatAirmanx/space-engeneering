package repository_part_map

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl/converter"
)

func (r *RepositoryPartImpl) SetPart(
	ctx context.Context,
	part *model_part.Part,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.parts[part.UUID] = repository_converter_part.ToRepository(part)

	return nil
}
