package service_part

import repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"

var _ PartService = (*partServiceImpl)(nil)

type partServiceImpl struct {
	repository repository_part.PartRepository
}

type NewPartServiceParams struct {
	Repository repository_part.PartRepository
}

func NewPartService(params NewPartServiceParams) *partServiceImpl {
	return &partServiceImpl{
		repository: params.Repository,
	}
}
