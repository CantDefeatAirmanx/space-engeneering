package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/test_data"
	configs_inventory "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/inventory"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

func main() {
	inventoryRepo := NewInventoryRepositoryMap()
	for idx := range test_data.InitialParts {
		err := inventoryRepo.CreatePart(context.Background(), &test_data.InitialParts[idx])
		if err != nil {
			log.Printf("Failed to create part: %v", err)
		}
	}

	inventoryService := NewInventoryServiceServer(
		NewInventoryServiceServerParams{
			Repository: inventoryRepo,
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

	inventory_v1.RegisterInventoryServiceServer(grpcServer, inventoryService)

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

type InventoryServiceServer struct {
	inventory_v1.UnimplementedInventoryServiceServer
	repository InventoryRepository
}

type NewInventoryServiceServerParams struct {
	Repository InventoryRepository
}

func NewInventoryServiceServer(params NewInventoryServiceServerParams) *InventoryServiceServer {
	return &InventoryServiceServer{
		repository: params.Repository,
	}
}

func (s *InventoryServiceServer) GetPart(
	ctx context.Context,
	req *inventory_v1.GetPartRequest,
) (*inventory_v1.GetPartResponse, error) {
	part, err := s.repository.GetPart(ctx, req.Uuid)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "part %s not found. %v", req.Uuid, err)
	}

	return &inventory_v1.GetPartResponse{
		Part: part,
	}, nil
}

func (s *InventoryServiceServer) ListParts(
	ctx context.Context,
	req *inventory_v1.ListPartsRequest,
) (*inventory_v1.ListPartsResponse, error) {
	parts, err := s.repository.ListAllParts(ctx, req.Filter)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to list parts. %v", err)
	}

	return &inventory_v1.ListPartsResponse{
		Parts: parts,
	}, nil
}

type InventoryRepository interface {
	CreatePart(ctx context.Context, part *inventory_v1.Part) error
	GetPart(ctx context.Context, partID string) (*inventory_v1.Part, error)
	ListAllParts(ctx context.Context, filter *inventory_v1.PartsFilter) ([]*inventory_v1.Part, error)
}

type InventoryRepositoryMap struct {
	mu    sync.RWMutex
	parts map[string]*inventory_v1.Part
}

type PartNotFoundError struct {
	PartID string
	Err    error
}

func (e *PartNotFoundError) Error() string {
	return fmt.Sprintf("part %s not found. %v", e.PartID, e.Err)
}

func NewInventoryRepositoryMap() *InventoryRepositoryMap {
	return &InventoryRepositoryMap{
		parts: make(map[string]*inventory_v1.Part),
	}
}

func (r *InventoryRepositoryMap) CreatePart(ctx context.Context, part *inventory_v1.Part) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	part.CreatedAt = timestamppb.New(time.Now())
	part.UpdatedAt = timestamppb.New(time.Now())

	r.parts[part.Uuid] = part

	return nil
}

func (r *InventoryRepositoryMap) GetPart(ctx context.Context, partID string) (*inventory_v1.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	part, ok := r.parts[partID]
	if !ok {
		return nil, &PartNotFoundError{
			PartID: partID,
			Err:    errors.New("not in map repo"),
		}
	}

	return part, nil
}

func (r *InventoryRepositoryMap) ListAllParts(
	ctx context.Context,
	filter *inventory_v1.PartsFilter,
) ([]*inventory_v1.Part, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	filters := []FilterFunc{
		createFilterByCategories(filter.Categories),
		createFilterByManufacturerCountries(filter.ManufacturerCountries),
		createFilterByTags(filter.Tags),
		createFilterByUuids(filter.Uuids),
	}
	parts := make([]*inventory_v1.Part, 0)

outer:
	for _, part := range r.parts {
		for _, filter := range filters {
			if !filter(part) {
				continue outer
			}
		}
		parts = append(parts, part)
	}

	return parts, nil
}

type FilterFunc func(part *inventory_v1.Part) bool

func createFilterByCategories(categories []inventory_v1.Category) FilterFunc {
	return func(part *inventory_v1.Part) bool {
		if len(categories) == 0 {
			return true
		}

		return slices.Contains(categories, part.Category)
	}
}

func createFilterByManufacturerCountries(manufacturerCountries []string) FilterFunc {
	return func(part *inventory_v1.Part) bool {
		if len(manufacturerCountries) == 0 {
			return true
		}

		return slices.Contains(manufacturerCountries, part.Manufacturer.Country)
	}
}

func createFilterByTags(tags []string) FilterFunc {
	return func(part *inventory_v1.Part) bool {
		if len(tags) == 0 {
			return true
		}

		for _, tag := range tags {
			if slices.Contains(part.Tags, tag) {
				return true
			}
		}

		return false
	}
}

func createFilterByUuids(uuids []string) FilterFunc {
	return func(part *inventory_v1.Part) bool {
		if len(uuids) == 0 {
			return true
		}

		return slices.Contains(uuids, part.Uuid)
	}
}
