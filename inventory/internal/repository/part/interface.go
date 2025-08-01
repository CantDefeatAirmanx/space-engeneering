package repository_part

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

type PartRepository interface {
	// GetPart returns a part by its UUID.
	//
	// Errors:
	//
	// - ([model_part.ErrPartNotFound]): if the part is not found.
	//
	// - ([model_part.ErrPartInternal](ErrPartInternal)): if the part failed to get.
	GetPart(ctx context.Context, uuid string) (*model_part.Part, error)

	// GetParts returns a list of parts by the given filter.
	//
	// Errors:
	//
	// - ([model_part.ErrPartInternal]): if the parts failed to get.
	GetParts(ctx context.Context, filter Filter) ([]*model_part.Part, error)

	// SetPart sets a part by its UUID.
	//
	// Errors:
	//
	// - ([model_part.ErrPartInternal]): if the part failed to set.
	SetPart(ctx context.Context, part *model_part.Part) error
}

type Filter struct {
	Uuids                 []string
	Categories            []repository_model_part.Category
	ManufacturerCountries []string
	Tags                  []string
	Names                 []string
}
