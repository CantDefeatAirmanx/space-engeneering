package repository_part_mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	constants_mongo "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/constants/mongo"
)

func (r *RepositoryPartMongoImpl) GetParts(
	ctx context.Context,
	filter model_part.Filter,
) ([]*model_part.Part, error) {
	mongoFilter := createFilter(&filter)
	cursor, err := r.collection.Find(ctx, mongoFilter)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", model_part.ErrPartInternal, err)
	}

	parts := make([]*model_part.Part, 0)
	for cursor.Next(ctx) {
		var part Part
		if err := cursor.Decode(&part); err != nil {
			return nil, fmt.Errorf("%w: %v", model_part.ErrPartInternal, err)
		}

		modelPart := ToModel(&part)
		parts = append(parts, &modelPart)
	}

	return parts, nil
}

func createFilter(filter *model_part.Filter) bson.M {
	repoFilter := ToRepoFilter(filter)
	mongoFilter := bson.M{}

	if len(repoFilter.Uuids) > 0 {
		mongoFilter[UUIDField] = bson.M{
			constants_mongo.InOperator: repoFilter.Uuids,
		}
	}

	if len(repoFilter.Categories) > 0 {
		mongoFilter[CategoryField] = bson.M{
			constants_mongo.InOperator: repoFilter.Categories,
		}
	}

	if len(repoFilter.Names) > 0 {
		mongoFilter[NameField] = bson.M{
			constants_mongo.InOperator: repoFilter.Names,
		}
	}

	if len(repoFilter.ManufacturerCountries) > 0 {
		mongoFilter[ManufacturerCountryField] = bson.M{
			constants_mongo.InOperator: repoFilter.ManufacturerCountries,
		}
	}

	if len(repoFilter.Tags) > 0 {
		mongoFilter[TagsField] = bson.M{
			constants_mongo.InOperator: repoFilter.Tags,
		}
	}

	return mongoFilter
}
