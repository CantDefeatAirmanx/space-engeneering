//go:build integration

package integration

import (
	"context"
	"slices"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/config"
	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/shared/test_data"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
	"github.com/CantDefeatAirmanx/space-engeneering/platform/pkg/logger"
	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
	dbOps  *DbOps

	grpcConn        *grpc.ClientConn
	inventoryClient inventory_v1.InventoryServiceClient
)

var _ = Describe("Inventory Service Integration Tests", Ordered, func() {
	BeforeAll(func() {
		ctx, cancel = context.WithCancel(suiteCtx)
		suiteCtx, cancel = context.WithCancel(context.Background())
		dbOps = NewDbOps(ctx, config.Config.Mongo().DBName(), "parts")

		url := testEnvironment.App.Address()
		logger.Logger().Info("Connecting to GRPC server", zap.String("url", url))

		conn, err := grpc.NewClient(
			url,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		Expect(err).To(Not(HaveOccurred()))

		logger.Logger().Info("Creating inventory client...")

		grpcClient := inventory_v1.NewInventoryServiceClient(conn)

		grpcConn = conn
		inventoryClient = grpcClient
	})

	AfterAll(func() {
		grpcConn.Close()
		dbOps.ClearCollection(ctx)
		cancel()
	})

	It("should successfully start", func() {
		Expect(true).To(BeTrue())
	})

	Describe("Get Part", func() {
		It("should get part by id", func() {
			part := dbOps.AddGeneratedPart(ctx)

			response, err := inventoryClient.GetPart(ctx, &inventory_v1.GetPartRequest{
				Uuid: part.UUID,
			})
			logger.Logger().Info("response", zap.Any("response", response.Part.Name))
			dbOps.ClearCollection(ctx)

			Expect(err).To(Not(HaveOccurred()))
		})

		It("should return error if part not found", func() {
			_, err := inventoryClient.GetPart(
				ctx, &inventory_v1.GetPartRequest{
					Uuid: "123",
				},
			)
			Expect(err).To(HaveOccurred())
			statusErr, ok := status.FromError(err)
			Expect(ok).To(BeTrue())
			Expect(statusErr.Code()).To(Equal(codes.NotFound))
		})
	})

	Describe("List Parts", func() {
		testData := test_data.GetInitialParts()

		BeforeAll(func() {
			for _, part := range testData {
				dbOps.AddPart(ctx, part)
			}
		})

		AfterAll(func() {
			for _, part := range testData {
				dbOps.RemovePart(ctx, part.UUID)
			}
		})

		It("should list all parts", func() {
			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{},
				},
			)
			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(len(testData)))
			for _, part := range response.Parts {
				Expect(hasPartWithUuid(response, part.Uuid)).To(BeTrue())
			}
		})

		It("should filter by UUIDs", func() {
			filterParts := testData[:2]
			uuids := lo.Map(filterParts, func(item *model_part.Part, _ int) string {
				return item.UUID
			})

			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Uuids: uuids,
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(len(uuids)))
			for _, part := range response.Parts {
				Expect(hasPartWithUuid(response, part.Uuid)).To(BeTrue())
			}
		})

		It("should return empty list if uuids not found", func() {
			randomUuid := uuid.Must(uuid.NewV7())
			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Uuids: []string{randomUuid.String()},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(0))
		})

		It("should filter by names", func() {
			partsToAddInfo := map[string]struct {
				name string
				uuid string
			}{
				"somePartName1":   {name: "somePartName1", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName2":   {name: "somePartName2", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName2-1": {name: "somePartName2", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName2-2": {name: "somePartName2", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName3":   {name: "somePartName3", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName3-1": {name: "somePartName3", uuid: uuid.Must(uuid.NewV7()).String()},
			}
			defer func() {
				for _, info := range partsToAddInfo {
					dbOps.RemovePart(ctx, info.uuid)
				}
			}()

			for _, info := range partsToAddInfo {
				part := helpers_test_data.GenerateRandomPart(
					helpers_test_data.WithName(info.name),
					helpers_test_data.WithUUID(info.uuid),
				)
				dbOps.AddPart(ctx, part)
			}

			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Names: []string{partsToAddInfo["somePartName1"].name},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(1))
			Expect(hasPartWithUuid(response, partsToAddInfo["somePartName1"].uuid)).To(BeTrue())

			response, err = inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Names: []string{partsToAddInfo["somePartName2"].name},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(3))
			Expect(hasPartWithUuid(response, partsToAddInfo["somePartName2"].uuid)).To(BeTrue())
			Expect(hasPartWithUuid(response, partsToAddInfo["somePartName2-1"].uuid)).To(BeTrue())
			Expect(hasPartWithUuid(response, partsToAddInfo["somePartName2-2"].uuid)).To(BeTrue())

			response, err = inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Names: []string{partsToAddInfo["somePartName3"].name},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(2))
			Expect(hasPartWithUuid(response, partsToAddInfo["somePartName3"].uuid)).To(BeTrue())
			Expect(hasPartWithUuid(response, partsToAddInfo["somePartName3-1"].uuid)).To(BeTrue())

			response, err = inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Names: []string{
							partsToAddInfo["somePartName1"].name,
							partsToAddInfo["somePartName2"].name,
							partsToAddInfo["somePartName3"].name,
						},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(len(partsToAddInfo)))
		})

		It("should return empty list if names not found", func() {
			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Names: []string{uuid.Must(uuid.NewV7()).String()},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(0))
		})

		It("should filter by tags", func() {
			partsToAdd := map[string]struct {
				tagName string
				uuid    string
			}{
				"somePartName1":   {tagName: "someTagName1", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName1-1": {tagName: "someTagName1", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName2":   {tagName: "someTagName2", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName3":   {tagName: "someTagName3", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName4":   {tagName: "someTagName4", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName5":   {tagName: "someTagName5", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName6":   {tagName: "someTagName6", uuid: uuid.Must(uuid.NewV7()).String()},
				"somePartName7":   {tagName: "someTagName7", uuid: uuid.Must(uuid.NewV7()).String()},
			}

			for _, info := range partsToAdd {
				part := helpers_test_data.GenerateRandomPart(
					helpers_test_data.WithUUID(info.uuid),
					helpers_test_data.WithTags([]string{info.tagName}),
				)
				dbOps.AddPart(ctx, part)
			}
			defer func() {
				for _, info := range partsToAdd {
					dbOps.RemovePart(ctx, info.uuid)
				}
			}()

			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Tags: []string{partsToAdd["somePartName1"].tagName},
					},
				},
			)

			Expect(hasPartWithUuid(response, partsToAdd["somePartName1"].uuid)).To(BeTrue())
			Expect(hasPartWithUuid(response, partsToAdd["somePartName1-1"].uuid)).To(BeTrue())
			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(2))

			response, err = inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Tags: []string{partsToAdd["somePartName2"].tagName},
					},
				},
			)

			Expect(hasPartWithUuid(response, partsToAdd["somePartName2"].uuid)).To(BeTrue())
			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(1))

			response, err = inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Tags: []string{
							partsToAdd["somePartName1"].tagName,
							partsToAdd["somePartName2"].tagName,
						},
					},
				},
			)

			Expect(hasPartWithUuid(response, partsToAdd["somePartName1"].uuid)).To(BeTrue())
			Expect(hasPartWithUuid(response, partsToAdd["somePartName2"].uuid)).To(BeTrue())
			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(3))
		})

		It("should return empty list if tags not found", func() {
			response, err := inventoryClient.ListParts(
				ctx,
				&inventory_v1.ListPartsRequest{
					Filter: &inventory_v1.PartsFilter{
						Tags: []string{uuid.Must(uuid.NewV7()).String()},
					},
				},
			)

			Expect(err).To(Not(HaveOccurred()))
			Expect(response.Parts).To(HaveLen(0))
		})
	})
})

func hasPartWithUuid(
	response *inventory_v1.ListPartsResponse,
	uuid string,
) bool {
	return slices.ContainsFunc(response.Parts, func(part *inventory_v1.Part) bool {
		return part.Uuid == uuid
	})
}
