package service_part

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

func (s *partServiceImpl) GetParts(
	ctx context.Context,
	filter model_part.Filter,
) ([]*model_part.Part, error) {
	parts, err := s.repository.GetParts(ctx, filter)
	if err != nil {
		return nil, model_part.ErrPartInternal
	}

	return parts, nil
}
