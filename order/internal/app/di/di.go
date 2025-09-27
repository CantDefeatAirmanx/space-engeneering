package di

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"

	"github.com/CantDefeatAirmanx/space-engeneering/order/config"
	api_order_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/api/order/v1"
	client_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1"
	client_payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1"
	repository_order_postgre "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/postgre_impl"
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	service_order_producer "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order/producer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/closer"
	platform_kafka_producer "github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/kafka/producer"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	order_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/openapi/order/v1"
)

type DiContainer struct {
	closer closer.Closer

	ordersDb *pgxpool.Pool

	orderHandler    *api_order_v1.Api
	orderService    service_order.OrderService
	orderRepository service_order.OrderRepository
	orderProducer   service_order.OrderProducer
	inventoryClient client_inventory_v1.InventoryV1Client
	paymentClient   client_payment_v1.PaymentV1Client
}

func NewDiContainer(closer closer.Closer) *DiContainer {
	return &DiContainer{
		closer: closer,
	}
}

func (d *DiContainer) GetOrderServer(ctx context.Context) http.Handler {
	server, err := order_v1.NewServer(
		d.GetOrderHandler(ctx),
	)
	if err != nil {
		logger.Logger().Error("Failed to create order server", zap.Error(err))
		panic(err)
	}

	return server
}

func (d *DiContainer) GetOrderHandler(ctx context.Context) *api_order_v1.Api {
	if d.orderHandler != nil {
		return d.orderHandler
	}

	api := api_order_v1.NewApi(
		api_order_v1.NewApiParams{
			OrderService: d.GetOrderService(ctx),
		},
	)
	d.orderHandler = api

	return api
}

func (d *DiContainer) GetOrderService(ctx context.Context) service_order.OrderService {
	if d.orderService != nil {
		return d.orderService
	}

	orderService := service_order.NewOrderService(
		service_order.NewOrderServiceParams{
			OrderRepository: d.GetOrderRepository(ctx),
			InventoryClient: d.GetInventoryClient(ctx),
			PaymentClient:   d.GetPaymentClient(ctx),
			OrderProducer:   d.GetOrderProducer(ctx),
		},
	)

	d.orderService = orderService

	return orderService
}

func (d *DiContainer) GetOrderRepository(
	ctx context.Context,
) service_order.OrderRepository {
	if d.orderRepository != nil {
		return d.orderRepository
	}

	orderRepository := repository_order_postgre.NewOrderRepositoryPostgre(
		d.GetDB(ctx),
	)

	d.orderRepository = orderRepository

	return orderRepository
}

func (d *DiContainer) GetDB(ctx context.Context) *pgxpool.Pool {
	if d.ordersDb != nil {
		return d.ordersDb
	}

	db, err := pgxpool.New(ctx, config.Config.Postgres().Uri())
	d.closer.AddNamed("Orders Postgres Db", func(ctx context.Context) error {
		db.Close()

		return nil
	})
	if err != nil {
		logger.Logger().Error("Failed to create orders db", zap.Error(err))
		panic(err)
	}

	d.ordersDb = db

	return db
}

func (d *DiContainer) GetInventoryClient(
	ctx context.Context,
) client_inventory_v1.InventoryV1Client {
	inventoryClient, err := client_inventory_v1.NewInventoryClient(
		ctx,
		config.Config.InventoryClient().Url(),
	)
	d.closer.AddNamed("Inventory GRPC Client", func(ctx context.Context) error {
		return inventoryClient.Close()
	})
	if err != nil {
		logger.Logger().Error("Failed to create inventory client", zap.Error(err))
		panic(err)
	}

	d.inventoryClient = inventoryClient

	return inventoryClient
}

func (d *DiContainer) GetPaymentClient(
	ctx context.Context,
) client_payment_v1.PaymentV1Client {
	paymentClient, err := client_payment_v1.NewPaymentClient(
		ctx,
		config.Config.PaymentClient().Url(),
	)
	d.closer.AddNamed("Payment GRPC Client", func(ctx context.Context) error {
		return paymentClient.Close()
	})
	if err != nil {
		logger.Logger().Error("Failed to create payment client", zap.Error(err))
		panic(err)
	}

	d.paymentClient = paymentClient

	return paymentClient
}

func (d *DiContainer) GetOrderProducer(
	ctx context.Context,
) service_order.OrderProducer {
	if d.orderProducer != nil {
		return d.orderProducer
	}

	kafkaProducer, err := platform_kafka_producer.NewKafkaProducer(
		ctx,
		config.Config.Kafka().Brokers(),
	)
	if err != nil {
		logger.Logger().Error("Failed to create kafka producer", zap.Error(err))
		panic(err)
	}

	orderProducer := service_order_producer.NewOrderProducer(
		kafkaProducer,
	)
	d.orderProducer = orderProducer

	return orderProducer
}
