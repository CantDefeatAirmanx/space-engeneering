package di

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/config"
	api_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/api/inventory/v1"
	repository_part_mongo "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/mongo_impl"
	service_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/shared/test_data"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	platform_mongo "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/db/mongo"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

type DiContainer struct {
	closer closer.Closer

	inventoryAPI   inventory_v1.InventoryServiceServer
	partService    service_part.PartService
	partRepository service_part.PartRepository
	mongoClient    *mongo.Client
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}

func (d *DiContainer) GetInventoryAPI(ctx context.Context) inventory_v1.InventoryServiceServer {
	if d.inventoryAPI != nil {
		return d.inventoryAPI
	}

	d.inventoryAPI = api_inventory_v1.NewApi(
		d.GetPartService(ctx),
	)
	logger.DefaultInfoLogger().Info("Inventory API created")

	return d.inventoryAPI
}

func (d *DiContainer) GetPartService(ctx context.Context) service_part.PartService {
	if d.partService != nil {
		return d.partService
	}

	d.partService = service_part.NewPartService(
		service_part.NewPartServiceParams{
			Repository: d.GetPartRepository(ctx),
		},
	)
	logger.DefaultInfoLogger().Info("Part service created")

	return d.partService
}

func (d *DiContainer) GetPartRepository(ctx context.Context) service_part.PartRepository {
	if d.partRepository != nil {
		return d.partRepository
	}

	client := d.GetMongoClient(ctx).Database(config.Config.Mongo().DBName())

	params := repository_part_mongo.NewRepositoryPartMongoImplParams{
		Db: client,
	}
	if config.Config.IsDev() {
		params.InitialParts = test_data.GetInitialParts()
	}

	d.partRepository = repository_part_mongo.NewRepositoryPartMongoImpl(
		ctx,
		params,
	)
	logger.DefaultInfoLogger().Info("Part repository created")

	return d.partRepository
}

func (d *DiContainer) GetMongoClient(ctx context.Context) *mongo.Client {
	if d.mongoClient != nil {
		return d.mongoClient
	}

	mongoClient, err := platform_mongo.Connect(
		ctx,
		platform_mongo.WithURI(config.Config.Mongo().URI()),
	)
	logger.Logger().Debug(fmt.Sprintf(
		"Connected to MongoDB: %v, err: %v",
		config.Config.Mongo().URI(), err),
	)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	d.mongoClient = mongoClient

	d.closer.AddNamed("Mongo client", func(ctx context.Context) error {
		return d.mongoClient.Disconnect(ctx)
	})

	return d.mongoClient
}
