package repository_part_mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

func (r *RepositoryPartMongoImpl) GetPart(
	ctx context.Context,
	uuid string,
) (*model_part.Part, error) {
	filter := bson.M{UUIDField: uuid}
	part := r.collection.FindOne(ctx, filter)

	if err := part.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, model_part.ErrPartNotFound
		}
		return nil, fmt.Errorf("%w: %s", model_part.ErrPartInternal, err.Error())
	}

	var result Part
	if err := part.Decode(&result); err != nil {
		return nil, fmt.Errorf("%w: %s", model_part.ErrPartInternal, err.Error())
	}

	modelPart := ToModel(&result)

	return &modelPart, nil
}
