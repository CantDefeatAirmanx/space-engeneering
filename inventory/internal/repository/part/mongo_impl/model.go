package repository_part_mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Part struct {
	ID            primitive.ObjectID        `bson:"_id,omitempty"`
	UUID          string                    `bson:"uuid"`
	Name          string                    `bson:"name"`
	Description   string                    `bson:"description"`
	Price         float64                   `bson:"price"`
	StockQuantity int64                     `bson:"stock_quantity"`
	Category      Category                  `bson:"category"`
	Dimensions    *Dimensions               `bson:"dimensions,omitempty"`
	Manufacturer  *Manufacturer             `bson:"manufacturer,omitempty"`
	Tags          []string                  `bson:"tags"`
	CreatedAt     *time.Time                `bson:"created_at,omitempty"`
	UpdatedAt     *time.Time                `bson:"updated_at,omitempty"`
	Metadata      map[string]*MetaDataValue `bson:"metadata"`
}

type Category int

const (
	CategoryUnknown Category = iota
	CategoryEngine
	CategoryFuel
	CategoryPortHole
	CategoryWing
)

type Dimensions struct {
	Length float64 `bson:"length"`
	Width  float64 `bson:"width"`
	Height float64 `bson:"height"`
	Weight float64 `bson:"weight"`
}

type Manufacturer struct {
	Name    string `bson:"name"`
	Country string `bson:"country"`
	Website string `bson:"website"`
}

type MetaDataValue struct {
	StringValue *string  `bson:"string_value,omitempty"`
	DoubleValue *float64 `bson:"double_value,omitempty"`
	Int64Value  *int64   `bson:"int64_value,omitempty"`
	BoolValue   *bool    `bson:"bool_value,omitempty"`
}

type Filter struct {
	Uuids                 []string
	Categories            []Category
	ManufacturerCountries []string
	Tags                  []string
	Names                 []string
}
