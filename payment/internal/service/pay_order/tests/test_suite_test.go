package service_pay_order_tests

import (
	"context"
	"testing"

	service_pay_order "github.com/CantDefeatAirmanx/space-engeneering/payment/internal/service/pay_order"
	"github.com/stretchr/testify/suite"
)

type TestingSuite struct {
	suite.Suite
	ctx     context.Context
	service service_pay_order.PayOrderService
}

func (s *TestingSuite) SetupTest() {
	ctx := context.Background()

	s.ctx = ctx
	s.service = service_pay_order.NewPayOrderServiceImpl()
}

func (s *TestingSuite) TearDownTest() {
	s.ctx = nil
	s.service = nil
}

func TestPayOrderSuite(t *testing.T) {
	suite.Run(t, new(TestingSuite))
}
