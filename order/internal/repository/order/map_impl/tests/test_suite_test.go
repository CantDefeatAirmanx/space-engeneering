package repository_order_map_tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	repository_order "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order"
	repository_order_map "github.com/CantDefeatAirmanx/space-engeneering/order/internal/repository/order/map_impl"
)

type TestingSuite struct {
	suite.Suite
	ctx  context.Context
	repo repository_order.OrderRepository
}

func (s *TestingSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = repository_order_map.NewOrderRepositoryMap()
}

func (s *TestingSuite) TearDownTest() {
	s.ctx = nil
	s.repo = nil
}

func TestTestingSuite(t *testing.T) {
	suite.Run(t, new(TestingSuite))
}
