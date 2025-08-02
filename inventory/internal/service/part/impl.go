package service_part

var _ PartService = (*partServiceImpl)(nil)

type partServiceImpl struct {
	repository PartRepository
}

type NewPartServiceParams struct {
	Repository PartRepository
}

func NewPartService(params NewPartServiceParams) *partServiceImpl {
	return &partServiceImpl{
		repository: params.Repository,
	}
}
