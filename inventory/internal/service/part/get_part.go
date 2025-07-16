package service_part

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/converter/part"
)

func (s *partServiceImpl) GetPart(
	ctx context.Context,
	uuid string,
) (*model_part.Part, error) {
	part, err := s.repository.GetPart(ctx, uuid)

	if err != nil {
		return nil, err
	}

	modelPart := repository_converter_part.ToModel(part)

	return &modelPart, nil
}
