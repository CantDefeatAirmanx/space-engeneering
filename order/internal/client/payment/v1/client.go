package client_payment_v1

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	configs_order "github.com/CantDefeatAirmanx/space-engeneering/shared/configs/server/order"
	payment_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/payment/v1"
)

var _ PaymentV1Client = (*paymentV1GrpcClient)(nil)

type paymentV1GrpcClient struct {
	grpcClient payment_v1.PaymentServiceClient
	Conn       *grpc.ClientConn
}

func NewPaymentClient(ctx context.Context) (*paymentV1GrpcClient, error) {
	//nolint:staticcheck
	conn, err := grpc.Dial(
		configs_order.PaymentServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	conn.Connect()
	conn.WaitForStateChange(ctx, connectivity.Ready)

	grpcClient := payment_v1.NewPaymentServiceClient(conn)

	return &paymentV1GrpcClient{
		grpcClient: grpcClient,
		Conn:       conn,
	}, nil
}
