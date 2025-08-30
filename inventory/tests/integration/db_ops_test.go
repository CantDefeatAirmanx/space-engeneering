//go:build integration

package integration

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_part_mongo "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/mongo_impl"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

type DbOps struct {
	db *mongo.Collection
}

func NewDbOps(ctx context.Context, dbName, collectionName string) *DbOps {
	return &DbOps{
		db: testEnvironment.Mongo.Client().Database(dbName).Collection(collectionName),
	}
}

func (d *DbOps) AddGeneratedPart(ctx context.Context, options ...helpers_test_data.PartOption) repository_part_mongo.Part {
	randomPart := helpers_test_data.GenerateRandomPart(options...)
	repoPart := repository_part_mongo.ToRepo(randomPart)

	_, err := d.db.InsertOne(ctx, repoPart)
	if err != nil {
		panic(err)
	}

	return repoPart
}

func (d *DbOps) AddPart(ctx context.Context, part *model_part.Part) {
	repoPart := repository_part_mongo.ToRepo(part)

	_, err := d.db.InsertOne(ctx, repoPart)
	if err != nil {
		panic(err)
	}
}

func (d *DbOps) RemovePart(ctx context.Context, uuid string) {
	_, err := d.db.DeleteOne(ctx, bson.M{"uuid": uuid})
	if err != nil {
		panic(err)
	}
}

func (d *DbOps) ClearCollection(ctx context.Context) {
	_, err := d.db.DeleteMany(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
}
