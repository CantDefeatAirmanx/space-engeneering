package repository_part_map

import (
	"context"

	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

func (r *RepositoryPartImpl) SetPart(
	ctx context.Context,
	part *repository_model_part.Part,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.parts[part.UUID] = *part

	return nil
}
