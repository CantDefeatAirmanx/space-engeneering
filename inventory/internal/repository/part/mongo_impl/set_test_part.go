package repository_part_mongo

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

// setTestPart sets or replaces a test part in the database.
//
// part should have UUID field set
func (r *RepositoryPartMongoImpl) setTestPart(
	ctx context.Context,
	part *model_part.Part,
) error {
	if part.UUID == "" {
		return fmt.Errorf("%w: %v", model_part.ErrPartInternal, "UUID is empty")
	}

	repoPart := ToRepo(part)
	filter := bson.M{UUIDField: repoPart.UUID}

	res := r.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			_, err := r.collection.InsertOne(ctx, repoPart)
			if err != nil {
				return fmt.Errorf("%w: %v", model_part.ErrPartInternal, err)
			}
			return nil
		}

		return res.Err()
	}

	_, err := r.collection.ReplaceOne(ctx, filter, repoPart)
	if err != nil {
		return fmt.Errorf("%w: %v", model_part.ErrPartInternal, err)
	}

	return nil
}
