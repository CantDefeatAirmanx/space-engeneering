package service_order_tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	inventory_v1_mocks "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/inventory/v1/mocks"
	payment_v1_mocks "github.com/CantDefeatAirmanx/space-engeneering/order/internal/client/payment/v1/mocks"
	repository_order_mocks "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/mocks"
	service_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order"
	service_order_mocks "github.com/CantDefeatAirmanx/space-engeneering/order/internal/service/order/mocks"
)

type TestingSuite struct {
	suite.Suite
	ctx     context.Context
	service service_order.OrderService

	orderProducerMock   *service_order_mocks.MockOrderProducer
	repoMock            *repository_order_mocks.MockOrderRepository
	inventoryClientMock *inventory_v1_mocks.MockInventoryV1Client
	paymentClientMock   *payment_v1_mocks.MockPaymentV1Client
}

func (s *TestingSuite) SetupTest() {
	s.ctx = context.Background()
	s.repoMock = repository_order_mocks.NewMockOrderRepository(s.T())
	s.inventoryClientMock = inventory_v1_mocks.NewMockInventoryV1Client(s.T())
	s.paymentClientMock = payment_v1_mocks.NewMockPaymentV1Client(s.T())
	s.orderProducerMock = service_order_mocks.NewMockOrderProducer(s.T())

	s.service = service_order.NewOrderService(
		service_order.NewOrderServiceParams{
			OrderRepository: s.repoMock,
			InventoryClient: s.inventoryClientMock,
			PaymentClient:   s.paymentClientMock,
			OrderProducer:   s.orderProducerMock,
		},
	)
}

func (s *TestingSuite) TearDownTest() {
	s.ctx = nil
	s.repoMock = nil
	s.inventoryClientMock = nil
	s.paymentClientMock = nil
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestingSuite))
}
