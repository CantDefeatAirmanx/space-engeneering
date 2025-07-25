package service_part_tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/mocks"
	service_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"
)

type TestingSuite struct {
	suite.Suite
	ctx      context.Context
	service  service_part.PartService
	mockRepo *mocks.MockPartRepository
}

func (s *TestingSuite) SetupTest() {
	s.ctx = context.Background()
	s.mockRepo = mocks.NewMockPartRepository(s.T())
	service := service_part.NewPartService(
		service_part.NewPartServiceParams{
			Repository: s.mockRepo,
		},
	)
	s.service = service
}

func (s *TestingSuite) TearDownTest() {
	s.ctx = nil
	s.mockRepo = nil
}

func TestTestingSuite(t *testing.T) {
	suite.Run(t, new(TestingSuite))
}
