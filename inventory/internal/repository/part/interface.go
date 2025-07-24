package repository_part

import (
	"context"

	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/part/model"
)

type PartRepository interface {
	// GetPart returns a part by its UUID.
	//
	// Errors:
	//
	// - ([repository_part.ErrPartNotFound]): if the part is not found.
	//
	// - ([repository_part.ErrPartInternal](ErrPartInternal)): if the part failed to get.
	GetPart(ctx context.Context, uuid string) (*repository_model_part.Part, error)

	// GetParts returns a list of parts by the given filter.
	//
	// Errors:
	//
	// - ([repository_part.ErrPartInternal]): if the parts failed to get.
	GetParts(ctx context.Context, filter Filter) ([]*repository_model_part.Part, error)

	// SetPart sets a part by its UUID.
	//
	// Errors:
	//
	// - ([repository_part.ErrPartInternal]): if the part failed to set.
	SetPart(ctx context.Context, part *repository_model_part.Part) error
}

type Filter struct {
	Uuids                 []string
	Categories            []repository_model_part.Category
	ManufacturerCountries []string
	Tags                  []string
	Names                 []string
}
