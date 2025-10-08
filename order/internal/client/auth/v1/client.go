package client_auth_v1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	auth_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/auth/v1"
)

var _ AuthV1Client = (*authV1GrpcClient)(nil)

type authV1GrpcClient struct {
	grpcClient auth_v1.AuthServiceClient
	conn       *grpc.ClientConn
}

func NewAuthClient(
	ctx context.Context,
	url string,
) (*authV1GrpcClient, error) {
	conn, err := grpc.NewClient(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	conn.Connect()
	conn.WaitForStateChange(ctx, connectivity.Ready)

	grpcClient := auth_v1.NewAuthServiceClient(conn)

	return &authV1GrpcClient{
		grpcClient: grpcClient,
		conn:       conn,
	}, nil
}
