package service_part_tests

import (
	"github.com/brianvoe/gofakeit/v7"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	service_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/service/part"
	helpers_mocks "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

func (s *TestingSuite) TestGetPartsSuccess() {
	partsUUID := []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()}

	modelGeneratedParts := []*model_part.Part{
		helpers_mocks.GenerateRandomPart(helpers_mocks.WithUUID(partsUUID[0])),
		helpers_mocks.GenerateRandomPart(helpers_mocks.WithUUID(partsUUID[1])),
		helpers_mocks.GenerateRandomPart(helpers_mocks.WithUUID(partsUUID[2])),
	}

	s.mockRepo.EXPECT().GetParts(s.ctx, repository_part.Filter{
		Uuids: partsUUID,
	}).Return(modelGeneratedParts, nil)

	parts, err := s.service.GetParts(s.ctx, service_part.Filter{
		Uuids: partsUUID,
	})

	s.Equal(modelGeneratedParts, parts)
	s.NoError(err)
}

func (s *TestingSuite) TestGetPartsInternalError() {
	partsUUID := []string{gofakeit.UUID(), gofakeit.UUID(), gofakeit.UUID()}

	s.mockRepo.EXPECT().GetParts(s.ctx, repository_part.Filter{
		Uuids: partsUUID,
	}).Return(nil, model_part.ErrPartInternal)

	parts, err := s.service.GetParts(s.ctx, service_part.Filter{
		Uuids: partsUUID,
	})

	s.ErrorIs(err, model_part.ErrPartInternal)
	s.Nil(parts)
}
