package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/api/inventory/v1"
	repository_part_mongo "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/mongo_impl"
	service_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/shared/test_data"
	configs_inventory "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/inventory"
	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interceptor"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func main() {
	ctx := context.Background()
	// ToDo: Add working with .env abstactions (4 module)
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}
	envPath := filepath.Join(workingDir, "inventory", ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	connStr := os.Getenv(configs_inventory.EnvMongoDbURI)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		discErr := client.Disconnect(ctx)
		if discErr != nil {
			log.Printf("Error disconnecting from database: %v", discErr)
		}
		log.Fatalf("Error connecting to database: %v", err)
	}

	partRepo := repository_part_mongo.NewRepositoryPartMongoImpl(
		repository_part_mongo.NewRepositoryPartMongoImplParams{
			Db:           client.Database(os.Getenv(configs_inventory.EnvMongoDbName)),
			InitialParts: test_data.GetInitialParts(),
		},
	)
	partService := service_part.NewPartService(
		service_part.NewPartServiceParams{
			Repository: partRepo,
		},
	)
	inventoryAPI := api_inventory_v1.NewApi(
		partService,
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", configs_inventory.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			log.Fatalf("Failed to close listener: %v", err)
		}
	}()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryErrorInterceptor(),
			interceptor.ValidateInterceptor(),
		),
	)
	reflection.Register(grpcServer)

	inventory_v1.RegisterInventoryServiceServer(grpcServer, inventoryAPI)

	go func() {
		fmt.Println("Inventory service started")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	grpcServer.GracefulStop()
	log.Println("Inventory service stopped")
}
