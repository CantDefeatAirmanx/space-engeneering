package repository_part_map_tests

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	helpers_mocks "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

func (s *TestingSuite) TestGetPartSuccess() {
	randomParts := []*model_part.Part{}
	for range 10 {
		modelPart := helpers_mocks.GenerateRandomPart()
		randomParts = append(randomParts, modelPart)
	}
	rndIdx := rand.Intn(len(randomParts))
	randomPartId := randomParts[rndIdx].UUID

	for _, part := range randomParts {
		s.repo.SetPart(s.ctx, part)
	}

	result, err := s.repo.GetPart(s.ctx, randomPartId)

	s.NoError(err)
	s.NotNil(result)
	s.Equal(*randomParts[rndIdx], *result)
}

func (s *TestingSuite) TestGetPartNotFound() {
	randomUUID := gofakeit.UUID()

	result, err := s.repo.GetPart(s.ctx, randomUUID)

	s.Error(err)
	s.ErrorIs(err, model_part.ErrPartNotFound)
	s.Nil(result)
}

func (s *TestingSuite) TestGetPartSuccessWithSameIdSet() {
	partsWithSameId := []*model_part.Part{}
	for range 10 {
		modelPart := helpers_mocks.GenerateRandomPart(
			helpers_mocks.WithUUID("dummy_uuid"),
		)
		partsWithSameId = append(partsWithSameId, modelPart)
	}

	for _, part := range partsWithSameId {
		s.repo.SetPart(s.ctx, part)
	}

	expectedPart := partsWithSameId[len(partsWithSameId)-1]
	result, err := s.repo.GetPart(s.ctx, expectedPart.UUID)

	s.NoError(err)
	s.NotNil(result)
	s.Equal(*expectedPart, *result)
}
