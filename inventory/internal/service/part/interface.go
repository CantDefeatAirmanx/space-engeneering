package service_part

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

type PartService interface {
	// GetPart returns a part by its UUID.
	//
	// Errors:
	//
	// - ([model_part.ErrPartNotFound]): if the part is not found.
	//
	// - ([model_part.ErrPartInternal]): if the part failed to get.
	GetPart(ctx context.Context, uuid string) (*model_part.Part, error)

	// GetParts returns a list of parts by the given filter.
	//
	// Errors:
	//
	// - ([model_part.ErrPartInternal]): if the part failed to get.
	GetParts(ctx context.Context, filter model_part.Filter) ([]*model_part.Part, error)
}
