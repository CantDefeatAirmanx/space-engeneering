package repository_part_map_tests

import (
	"slices"
	"strings"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/converter"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

func (s *TestingSuite) TestGetParts() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	tcases := []struct {
		name     string
		expected []repository_model_part.Part
		err      error
		filter   repository_part.Filter
	}{
		{
			name:     "success get parts without filters",
			expected: partsValues,
			err:      nil,
			filter:   repository_part.Filter{},
		},

		// Uuids filter tests
		{
			name: "success get parts with uuids filter",
			expected: []repository_model_part.Part{
				partsValues[1],
				partsValues[2],
			},
			err: nil,
			filter: repository_part.Filter{
				Uuids: []string{part2UUID, part3UUID},
			},
		},
		{
			name:     "not found parts with uuids filter",
			expected: []repository_model_part.Part{},
			err:      nil,
			filter: repository_part.Filter{
				Uuids: []string{"random_uuid1", "random_uuid2"},
			},
		},

		// Tags filter tests
		{
			name: "success get parts with single tag filter",
			expected: []repository_model_part.Part{
				partsValues[4],
			},
			err: nil,
			filter: repository_part.Filter{
				Tags: []string{tag5},
			},
		},
		{
			name: "success get parts with multiple tags filter",
			expected: []repository_model_part.Part{
				partsValues[0],
				partsValues[1],
			},
			err: nil,
			filter: repository_part.Filter{
				Tags: []string{tag1, tag2},
			},
		},
		{
			name:     "not found parts with tag filter",
			expected: []repository_model_part.Part{},
			err:      nil,
			filter: repository_part.Filter{
				Tags: []string{"random_tag"},
			},
		},
		{
			name:     "not found parts with multiple tags filter",
			expected: []repository_model_part.Part{},
			err:      nil,
			filter: repository_part.Filter{
				Tags: []string{"random_tag", "random_tag2"},
			},
		},

		// Categories filter tests
		{
			name: "success get parts with category filter",
			expected: []repository_model_part.Part{
				partsValues[0],
				partsValues[2],
			},
			err: nil,
			filter: repository_part.Filter{
				Categories: []repository_model_part.Category{repository_model_part.CategoryEngine},
			},
		},
		{
			name: "success get parts with multiple categories filter",
			expected: []repository_model_part.Part{
				partsValues[0],
				partsValues[1],
				partsValues[2],
			},
			err: nil,
			filter: repository_part.Filter{
				Categories: []repository_model_part.Category{
					repository_model_part.CategoryEngine,
					repository_model_part.CategoryFuel,
				},
			},
		},
		{
			name:     "not found parts with category filter",
			expected: []repository_model_part.Part{},
			err:      nil,
			filter: repository_part.Filter{
				Categories: []repository_model_part.Category{
					repository_model_part.CategoryUnknown,
				},
			},
		},

		// Names filter tests
		{
			name: "success get parts with name filter",
			expected: []repository_model_part.Part{
				partsValues[0],
			},
			err: nil,
			filter: repository_part.Filter{
				Names: []string{part1Name},
			},
		},
		{
			name: "success get parts with multiple names filter",
			expected: []repository_model_part.Part{
				partsValues[0],
				partsValues[2],
			},
			err: nil,
			filter: repository_part.Filter{
				Names: []string{part1Name, part3Name},
			},
		},
		{
			name:     "not found parts with name filter",
			expected: []repository_model_part.Part{},
			err:      nil,
			filter: repository_part.Filter{
				Names: []string{"random_name"},
			},
		},
	}

	for _, tcase := range tcases {
		s.Run(tcase.name, func() {
			parts, err := s.repo.GetParts(
				s.ctx,
				tcase.filter,
			)

			partsValues := getPartsValues(parts)

			sortParts(partsValues)
			sortParts(tcase.expected)

			s.NoError(err)
			s.Equal(tcase.expected, partsValues)
		})
	}
}

const (
	part1Name = "part1"
	part2Name = "part2"
	part3Name = "part3"
	part4Name = "part4"
	part5Name = "part5"
)

const (
	part1UUID = "part_1"
	part2UUID = "part_2"
	part3UUID = "part_3"
	part4UUID = "part_4"
	part5UUID = "part_5"
)

const (
	tag1 = "tag1"
	tag2 = "tag2"
	tag3 = "tag3"
	tag4 = "tag4"
	tag5 = "tag5"
)

var (
	part1Tags = []string{tag1, tag2, tag3}
	part2Tags = []string{tag1, tag2}
	part3Tags = []string{tag3}
	part4Tags = []string{tag3, tag4}
	part5Tags = []string{tag4, tag5}
)

func initParts(s *TestingSuite) []*model_part.Part {
	parts := []*model_part.Part{
		helpers_test_data.GenerateRandomPart(
			helpers_test_data.WithUUID(part1UUID),
			helpers_test_data.WithTags(part1Tags),
			helpers_test_data.WithCategory(model_part.CategoryEngine),
			helpers_test_data.WithName(part1Name),
		),
		helpers_test_data.GenerateRandomPart(
			helpers_test_data.WithUUID(part2UUID),
			helpers_test_data.WithTags(part2Tags),
			helpers_test_data.WithCategory(model_part.CategoryFuel),
			helpers_test_data.WithName(part2Name),
		),
		helpers_test_data.GenerateRandomPart(
			helpers_test_data.WithUUID(part3UUID),
			helpers_test_data.WithTags(part3Tags),
			helpers_test_data.WithCategory(model_part.CategoryEngine),
			helpers_test_data.WithName(part3Name),
		),
		helpers_test_data.GenerateRandomPart(
			helpers_test_data.WithUUID(part4UUID),
			helpers_test_data.WithTags(part4Tags),
			helpers_test_data.WithCategory(model_part.CategoryWing),
			helpers_test_data.WithName(part4Name),
		),
		helpers_test_data.GenerateRandomPart(
			helpers_test_data.WithUUID(part5UUID),
			helpers_test_data.WithTags(part5Tags),
			helpers_test_data.WithCategory(model_part.CategoryPortHole),
			helpers_test_data.WithName(part5Name),
		),
	}

	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}

	for _, part := range repoParts {
		s.repo.SetPart(s.ctx, part)
	}

	return parts
}

func getPartsValues(parts []*repository_model_part.Part) []repository_model_part.Part {
	partsValues := []repository_model_part.Part{}
	for _, part := range parts {
		partsValues = append(partsValues, *part)
	}
	return partsValues
}

func sortParts(parts []repository_model_part.Part) {
	slices.SortFunc(parts, func(a, b repository_model_part.Part) int {
		return strings.Compare(a.UUID, b.UUID)
	})
}
