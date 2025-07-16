package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	api_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/api/inventory/v1"
	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	service_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	configs_inventory "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/inventory"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func main() {
	partRepo := repository_part.NewRepositoryPart()
	partService := service_part.NewPartService(
		service_part.NewPartServiceParams{
			Repository: partRepo,
		},
	)
	inventoryAPI := api_inventory_v1.NewApi(
		api_inventory_v1.NewApiParams{
			PartService: partService,
		},
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

	grpcServer := grpc.NewServer()
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
