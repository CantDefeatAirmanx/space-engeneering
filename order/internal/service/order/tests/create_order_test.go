package service_order_tests

import (
	"context"
	"errors"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/mock"

	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
	client_inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1"
	model_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/order"
	model_part "github.com/CantDefeatAirmanx/space-engeneering/order/internal/model/part"
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
)

func (s *TestingSuite) TestCreateOrderSuccess() {
	parts := []*model_part.Part{}
	for range 10 {
		randomPart := helpers_test_data.GenerateRandomPart()
		parts = append(parts, &model_part.Part{
			UUID:  randomPart.UUID,
			Price: randomPart.Price,
		})
	}

	partUuids := []string{}
	for _, part := range parts {
		partUuids = append(partUuids, part.UUID)
	}

	s.inventoryClientMock.EXPECT().ListParts(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		client_inventory_v1.ListPartsParams{
			Uuids: partUuids,
		},
	).Return(parts, nil)

	totalPrice := 0.0
	for _, part := range parts {
		totalPrice += part.Price
	}

	userUUID := gofakeit.UUID()

	s.repoMock.EXPECT().CreateOrder(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		mock.MatchedBy(func(repoOrder model_order.Order) bool {
			return true
		}),
	).Return(nil)

	result, err := s.service.CreateOrder(
		s.ctx,
		service_order.CreateOrderParams{
			UserUUID:  userUUID,
			PartUuids: partUuids,
		},
	)

	s.NoError(err)
	s.NotEmpty(result.OrderUUID)
	s.Equal(totalPrice, result.TotalPrice)
}

func (s *TestingSuite) TestCreateOrderInvalidArgumentsClientInventory() {
	s.inventoryClientMock.EXPECT().ListParts(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		client_inventory_v1.ListPartsParams{
			Uuids: []string{},
		},
	).Return(nil, model_part.ErrPartInvalidArguments)

	result, err := s.service.CreateOrder(
		s.ctx,
		service_order.CreateOrderParams{
			UserUUID:  gofakeit.UUID(),
			PartUuids: []string{},
		},
	)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderInvalidArguments)
	s.Nil(result)
}

func (s *TestingSuite) TestCreateOrderClientInventoryUnknownError() {
	s.inventoryClientMock.EXPECT().ListParts(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		client_inventory_v1.ListPartsParams{
			Uuids: []string{},
		},
	).Return(nil, errors.New("unknown error"))

	result, err := s.service.CreateOrder(
		s.ctx,
		service_order.CreateOrderParams{
			UserUUID:  gofakeit.UUID(),
			PartUuids: []string{},
		},
	)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderInternal)
	s.Nil(result)
}

func (s *TestingSuite) TestCreateOrderInvalidPartUuids() {
	parts := []*model_part.Part{}
	for range 10 {
		randomPart := helpers_test_data.GenerateRandomPart()
		parts = append(parts, &model_part.Part{
			UUID:  randomPart.UUID,
			Price: randomPart.Price,
		})
	}

	partUuids := []string{}
	for _, part := range parts {
		partUuids = append(partUuids, part.UUID)
	}

	s.inventoryClientMock.EXPECT().ListParts(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		client_inventory_v1.ListPartsParams{
			Uuids: partUuids,
		},
	).Return(parts[:5:5], nil)

	result, err := s.service.CreateOrder(
		s.ctx,
		service_order.CreateOrderParams{
			UserUUID:  gofakeit.UUID(),
			PartUuids: partUuids,
		},
	)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderInvalidArguments)
	s.Nil(result)
}

func (s *TestingSuite) TestCreateOrderRepositoryUnknownError() {
	s.inventoryClientMock.EXPECT().ListParts(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		client_inventory_v1.ListPartsParams{
			Uuids: []string{},
		},
	).Return([]*model_part.Part{}, nil)

	s.repoMock.EXPECT().CreateOrder(
		mock.MatchedBy(func(ctx context.Context) bool {
			_, ok := ctx.Deadline()
			return ok
		}),
		mock.MatchedBy(func(repoOrder model_order.Order) bool {
			return true
		}),
	).Return(model_order.ErrOrderInternal)

	result, err := s.service.CreateOrder(
		s.ctx,
		service_order.CreateOrderParams{
			UserUUID:  gofakeit.UUID(),
			PartUuids: []string{},
		},
	)

	s.Error(err)
	s.ErrorIs(err, model_order.ErrOrderInternal)
	s.Nil(result)
}
