package service_part

import (
	"context"

	model_part "github.com/CantDefeatAirmanx/space-engeneering/inventory/internal/model/part"
)

type PartService interface {
	GetPart(ctx context.Context, uuid string) (*model_part.Part, error)
	GetParts(ctx context.Context, filter Filter) ([]*model_part.Part, error)
}

type Filter struct {
	Uuids                 []string
	Categories            []model_part.Category
	ManufacturerCountries []string
	Tags                  []string
	Names                 []string
}
