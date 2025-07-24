package repository_part_map_tests

import (
	"math/rand"

	"github.com/brianvoe/gofakeit/v7"

	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

func (s *TestingSuite) TestGetPart() {
	randomParts := []*repository_model_part.Part{}
	for range 10 {
		randomParts = append(randomParts, generateRandomPart())
	}
	rndIdx := rand.Intn(len(randomParts))
	randomPartId := randomParts[rndIdx].UUID

	partsWithSameId := []*repository_model_part.Part{}
	for range 10 {
		partsWithSameId = append(partsWithSameId, generateRandomPart(WithUUID("dummy_uuid")))
	}

	tcases := []struct {
		name     string
		expected *repository_model_part.Part
		err      error
		setUp    func()
		uuid     string
	}{
		{
			name:     "success get",
			expected: randomParts[rndIdx],
			err:      nil,
			setUp: func() {
				for _, part := range randomParts {
					s.repo.SetPart(s.ctx, part)
				}
			},
			uuid: randomPartId,
		},
		{
			name:     "not found",
			expected: nil,
			err:      &repository_part.ErrPartNotFound{},
			setUp:    func() {},
			uuid:     gofakeit.UUID(),
		},
		{
			name:     "success get with same id set",
			expected: partsWithSameId[len(partsWithSameId)-1],
			err:      nil,
			setUp: func() {
				for _, part := range partsWithSameId {
					s.repo.SetPart(s.ctx, part)
				}
			},
			uuid: partsWithSameId[len(partsWithSameId)-1].UUID,
		},
	}

	for _, tc := range tcases {
		s.Run(tc.name, func() {
			if tc.setUp != nil {
				tc.setUp()
			}

			part, err := s.repo.GetPart(
				s.ctx,
				tc.uuid,
			)

			if tc.expected != nil {
				s.Equal(*tc.expected, *part)
			} else {
				s.Nil(part)
			}

			s.ErrorIs(err, tc.err)
		})
	}
}
