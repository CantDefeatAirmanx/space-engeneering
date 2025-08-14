package repository_part_mongo

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

func (r *RepositoryPartMongoImpl) SetPart(
	ctx context.Context,
	part *model_part.Part,
) error {
	repoPart := ToRepo(part)
	uuid := uuid.Must(uuid.NewV7()).String()
	repoPart.UUID = uuid

	_, err := r.collection.InsertOne(ctx, repoPart)
	if err != nil {
		return fmt.Errorf("%w: %v", model_part.ErrPartInternal, err)
	}

	return nil
}
