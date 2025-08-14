package repository_part_map_tests

import (
	"slices"
	"strings"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl/converter"
	"github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/map_impl/model"
	helpers_test_data "github.com/CantDefeatAirmanx/space-engeneering/inventory/pkg/lib/helpers/test_data"
)

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

func (s *TestingSuite) TestGetPartsWithoutFilters() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	expected := getPartsValues(repoParts)

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithUuidsFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[1],
		partsValues[2],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Uuids: []string{part2UUID, part3UUID},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithUuidsFilterNotFound() {
	initParts(s)

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Uuids: []string{"random_uuid1", "random_uuid2"},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))

	s.NoError(err)
	s.Empty(resultValues)
}

func (s *TestingSuite) TestGetPartsWithSingleTagFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[4],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Tags: []string{tag5},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithMultipleTagsFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[0],
		partsValues[1],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Tags: []string{tag1, tag2},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithTagFilterNotFound() {
	initParts(s)

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Tags: []string{"random_tag"},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))

	s.NoError(err)
	s.Empty(resultValues)
}

func (s *TestingSuite) TestGetPartsWithMultipleTagsFilterNotFound() {
	initParts(s)

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Tags: []string{"random_tag", "random_tag2"},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))

	s.NoError(err)
	s.Empty(resultValues)
}

func (s *TestingSuite) TestGetPartsWithCategoryFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[0],
		partsValues[2],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Categories: []model_part.Category{model_part.CategoryEngine},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithMultipleCategoriesFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[0],
		partsValues[1],
		partsValues[2],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Categories: []model_part.Category{
			model_part.CategoryEngine,
			model_part.CategoryFuel,
		},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithCategoryFilterNotFound() {
	initParts(s)

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Categories: []model_part.Category{
			model_part.CategoryUnknown,
		},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))

	s.NoError(err)
	s.Empty(resultValues)
}

func (s *TestingSuite) TestGetPartsWithNameFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[0],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Names: []string{part1Name},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithMultipleNamesFilter() {
	parts := initParts(s)
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	partsValues := getPartsValues(repoParts)

	expected := []repository_model_part.Part{
		partsValues[0],
		partsValues[2],
	}

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Names: []string{part1Name, part3Name},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))
	sortParts(resultValues)
	sortParts(expected)

	s.NoError(err)
	s.Equal(expected, resultValues)
}

func (s *TestingSuite) TestGetPartsWithNameFilterNotFound() {
	initParts(s)

	result, err := s.repo.GetParts(s.ctx, model_part.Filter{
		Names: []string{"random_name"},
	})

	resultValues := getPartsValues(modelPartsToRepositoryParts(result))

	s.NoError(err)
	s.Empty(resultValues)
}

func modelPartsToRepositoryParts(parts []*model_part.Part) []*repository_model_part.Part {
	repoParts := []*repository_model_part.Part{}
	for _, part := range parts {
		repoPart := repository_converter_part.ToRepository(part)
		repoParts = append(repoParts, &repoPart)
	}
	return repoParts
}

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
		modelPart := repository_converter_part.ToModel(part)
		s.repo.SetPart(s.ctx, &modelPart)
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
