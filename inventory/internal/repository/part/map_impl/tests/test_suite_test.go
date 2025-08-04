package repository_part_map_tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	repository_part_map "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl"
)

type TestingSuite struct {
	suite.Suite
	ctx  context.Context
	repo *repository_part_map.RepositoryPartImpl
}

func (s *TestingSuite) SetupTest() {
	s.ctx = context.Background()
	s.repo = repository_part_map.NewRepositoryPart(
		repository_part_map.NewRepositoryPartParams{},
	)
}

func (s *TestingSuite) TearDownTest() {
	s.ctx = nil
	s.repo = nil
}

func TestTestingSuite(t *testing.T) {
	suite.Run(t, new(TestingSuite))
}
