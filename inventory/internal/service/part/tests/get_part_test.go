package service_part_tests

import (
	"github.com/brianvoe/gofakeit/v7"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/converter"
	helpers_mocks "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

func (s *TestingSuite) TestGetPartSuccess() {
	partUUID := gofakeit.UUID()

	modelGeneratedPart := helpers_mocks.GenerateRandomPart(
		helpers_mocks.WithUUID(partUUID),
	)
	repositoryGeneratedPart := repository_converter_part.ToRepository(modelGeneratedPart)
	s.mockRepo.EXPECT().GetPart(s.ctx, partUUID).Return(&repositoryGeneratedPart, nil)

	part, err := s.service.GetPart(s.ctx, partUUID)
	s.NoError(err)
	s.Equal(*modelGeneratedPart, *part)
}

func (s *TestingSuite) TestGetPartNotFound() {
	partUUID := gofakeit.UUID()
	s.mockRepo.EXPECT().GetPart(s.ctx, partUUID).Return(nil, &model_part.ErrPartNotFound{})

	part, err := s.service.GetPart(s.ctx, partUUID)

	s.ErrorIs(err, &model_part.ErrPartNotFound{})
	s.Nil(part)
}

func (s *TestingSuite) TestGetPartInternalError() {
	partUUID := gofakeit.UUID()
	s.mockRepo.EXPECT().GetPart(s.ctx, partUUID).Return(nil, &model_part.ErrPartInternal{})

	part, err := s.service.GetPart(s.ctx, partUUID)

	s.ErrorIs(err, &model_part.ErrPartInternal{})
	s.Nil(part)
}
