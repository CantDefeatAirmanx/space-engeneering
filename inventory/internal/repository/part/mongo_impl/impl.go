package repository_part_mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
)

var _ repository_part.PartRepository = (*RepositoryPartMongoImpl)(nil)

type RepositoryPartMongoImpl struct {
	collection *mongo.Collection
}

type NewRepositoryPartMongoImplParams struct {
	Db           *mongo.Database
	InitialParts []*model_part.Part
}

func NewRepositoryPartMongoImpl(params NewRepositoryPartMongoImplParams) *RepositoryPartMongoImpl {
	repo := RepositoryPartMongoImpl{
		collection: params.Db.Collection(partsCollectionName),
	}

	err := initIndexes(&repo)
	if err != nil {
		log.Fatalf("Failed to create indexes: %v", err)
	}

	// ToDo: IS_DEV check
	for _, part := range params.InitialParts {
		err := repo.setTestPart(context.Background(), part)
		if err != nil {
			log.Fatalf("Failed to insert initial part: %v", err)
		}
	}

	return &repo
}

func initIndexes(repo *RepositoryPartMongoImpl) error {
	indexModels := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: UUIDField, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{
				{Key: NameField, Value: 1},
			},
			Options: options.Index().SetUnique(false),
		},
		{
			Keys: bson.D{
				{Key: CategoryField, Value: 1},
			},
			Options: options.Index().SetUnique(false),
		},
	}

	_, err := repo.collection.Indexes().CreateMany(
		context.Background(),
		indexModels,
	)

	return err
}
