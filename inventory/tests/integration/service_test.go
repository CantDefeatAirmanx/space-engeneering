//go:build integration

package integration

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"

	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	grpcConn        *grpc.ClientConn
	inventoryClient inventory_v1.InventoryServiceClient
)

var _ = Describe("Inventory Service Integration Tests", Ordered, func() {
	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		url := env.App.Address()

		conn, err := grpc.NewClient(
			url,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		Expect(err).To(Not(HaveOccurred()))

		conn.WaitForStateChange(ctx, connectivity.Ready)

		grpcClient := inventory_v1.NewInventoryServiceClient(conn)

		grpcConn = conn
		inventoryClient = grpcClient
	})

	AfterEach(func() {
		grpcConn.Close()
		cancel()
	})

	It("should successfully start", func() {
		Expect(true).To(BeTrue())
	})

	It("should list all parts", func() {
		_, err := inventoryClient.ListParts(
			ctx,
			&inventory_v1.ListPartsRequest{},
		)
		Expect(err).To(Not(HaveOccurred()))
	})
})
