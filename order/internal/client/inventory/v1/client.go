package client_inventory_v1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/interfaces"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

var (
	_ InventoryV1Client    = (*inventoryV1GrpcClient)(nil)
	_ interfaces.WithClose = (*inventoryV1GrpcClient)(nil)
)

type inventoryV1GrpcClient struct {
	grpcClient inventory_v1.InventoryServiceClient
	conn       *grpc.ClientConn
}

func NewInventoryClient(
	ctx context.Context,
	url string,
) (*inventoryV1GrpcClient, error) {
	conn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	conn.Connect()
	conn.WaitForStateChange(ctx, connectivity.Ready)

	grpcClient := inventory_v1.NewInventoryServiceClient(conn)

	return &inventoryV1GrpcClient{
		grpcClient: grpcClient,
		conn:       conn,
	}, nil
}
