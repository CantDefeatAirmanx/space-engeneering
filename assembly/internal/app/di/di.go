package di

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/assembly/config"
	api_v1 "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/api/v1"
	repository_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/repository/ship_assembly"
	repository_ship_assembly_postgres "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/repository/ship_assembly/postgres"
	service_ship_assembly "github.com/CantDefeatAirmanx/space-engeneering/assembly/internal/service/ship_assembly"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	platform_kafka "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka"
	platform_kafka_consumer "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/consumer"
	platform_kafka_producer "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/producer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	assembly_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/assembly/v1"
)

type DiContainer struct {
	closer closer.Closer

	assemblyApi            assembly_v1.AssemblyServiceServer
	shipAssemblyService    service_ship_assembly.ShipAssemblyService
	shipAssemblyRepository repository_ship_assembly.ShipAssemblyRepository
	consumer               platform_kafka.Consumer
	producer               platform_kafka.Producer
	postgres               *pgxpool.Pool
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}

func (d *DiContainer) GetAssemblyApi(
	ctx context.Context,
) assembly_v1.AssemblyServiceServer {
	if d.assemblyApi != nil {
		return d.assemblyApi
	}

	d.assemblyApi = api_v1.NewApi(
		d.GetShipAssemblyService(ctx),
	)

	return d.assemblyApi
}

func (d *DiContainer) GetShipAssemblyService(
	ctx context.Context,
) service_ship_assembly.ShipAssemblyService {
	if d.shipAssemblyService != nil {
		return d.shipAssemblyService
	}

	service := service_ship_assembly.NewShipAssemblyService(
		ctx,
		d.GetShipAssemblyRepository(ctx),
		d.GetConsumer(ctx),
		d.GetProducer(ctx),
	)
	d.shipAssemblyService = service

	return service
}

func (d *DiContainer) GetShipAssemblyRepository(
	ctx context.Context,
) repository_ship_assembly.ShipAssemblyRepository {
	if d.shipAssemblyRepository != nil {
		return d.shipAssemblyRepository
	}

	shipAssemblyRepository := repository_ship_assembly_postgres.NewShipAssemblyRepositoryPostgres(
		d.getPostgres(ctx),
	)
	d.shipAssemblyRepository = shipAssemblyRepository

	return shipAssemblyRepository
}

func (d *DiContainer) getPostgres(
	ctx context.Context,
) *pgxpool.Pool {
	if d.postgres != nil {
		return d.postgres
	}

	postgres, err := pgxpool.New(ctx, config.Config.Postgres().Uri())
	if err != nil {
		logger.Logger().Error("Failed to create postgres", zap.Error(err))
		panic(err)
	}
	d.postgres = postgres

	d.closer.AddNamed("Postgres", func(ctx context.Context) error {
		postgres.Close()

		return nil
	})

	return d.postgres
}

func (d *DiContainer) GetProducer(
	ctx context.Context,
) platform_kafka.Producer {
	if d.producer != nil {
		return d.producer
	}

	producer, err := platform_kafka_producer.NewKafkaProducer(
		ctx,
		config.Config.Kafka().Brokers(),
	)
	if err != nil {
		logger.Logger().Error("Failed to create kafka producer", zap.Error(err))
		panic(err)
	}
	d.producer = producer

	return producer
}

func (d *DiContainer) GetConsumer(
	ctx context.Context,
) platform_kafka.Consumer {
	if d.consumer != nil {
		return d.consumer
	}

	consumer, err := platform_kafka_consumer.NewKafkaConsumer(
		config.Config.Kafka().Brokers(),
		"Assembly Consumer",
	)
	if err != nil {
		logger.Logger().Error("Failed to create kafka consumer", zap.Error(err))
		panic(err)
	}
	d.consumer = consumer

	return consumer
}
