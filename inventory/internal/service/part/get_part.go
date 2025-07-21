package service_part

import (
	"context"

	repository_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part"
	repository_converter_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/converter"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

func (s *partServiceImpl) GetPart(
	ctx context.Context,
	uuid string,
) (*model_part.Part, error) {
	part, err := s.repository.GetPart(ctx, uuid)

	if err != nil {
		switch err {
		case &repository_part.ErrPartNotFound{}:
			return nil, model_part.ErrPartNotFound{
				UUID: uuid,
				Err:  err,
			}
		default:
			return nil, model_part.ErrPartInternal{
				UUID: uuid,
				Err:  err,
			}
		}
	}

	modelPart := repository_converter_part.ToModel(part)

	return &modelPart, nil
}
