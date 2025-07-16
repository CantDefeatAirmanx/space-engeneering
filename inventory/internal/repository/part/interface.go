package repository_part

import (
	"context"

	repository_model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/repository/model/part"
)

type PartRepository interface {
	GetPart(ctx context.Context, uuid string) (*repository_model_part.Part, error)
	GetParts(ctx context.Context, filter Filter) ([]*repository_model_part.Part, error)
}

type Filter struct {
	Uuids                 []string
	Categories            []repository_model_part.Category
	ManufacturerCountries []string
	Tags                  []string
	Names                 []string
}
