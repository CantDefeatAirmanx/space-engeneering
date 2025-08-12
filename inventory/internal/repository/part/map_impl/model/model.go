package repository_model_part

import (
	"time"
)

// Part описывает деталь
type Part struct {
	// UUID идентификатор детали
	UUID string
	// Name название детали
	Name string
	// Description описание детали
	Description string
	// Price цена детали
	Price float64
	// StockQuantity количество деталей
	StockQuantity int64
	// Category категория детали
	Category Category
	// Dimensions размеры детали
	Dimensions *Dimensions
	// Manufacturer производитель детали
	Manufacturer *Manufacturer
	// Tags теги детали
	Tags []string
	// CreatedAt дата создания детали
	CreatedAt *time.Time
	// UpdatedAt дата обновления детали
	UpdatedAt *time.Time
	// Metadata метаданные детали
	Metadata map[string]*MetaDataValue
}

// Category описывает категорию детали
type Category int

const (
	// CategoryUnknown неизвестная категория
	CategoryUnknown Category = iota
	// CategoryEngine детали для двигателей
	CategoryEngine
	// CategoryFuel детали для топлива
	CategoryFuel
	// CategoryPortHole детали для отверстий
	CategoryPortHole
	// CategoryWing детали для крыльев
	CategoryWing
)

// String возвращает строковое представление категории
func (c Category) String() string {
	switch c {
	case CategoryEngine:
		return "ENGINE"
	case CategoryFuel:
		return "FUEL"
	case CategoryPortHole:
		return "PORT_HOLE"
	case CategoryWing:
		return "WING"
	default:
		return "UNKNOWN"
	}
}

// Dimensions описывает размеры детали
type Dimensions struct {
	// Length длина
	Length float64
	// Width ширина
	Width float64
	// Height высота
	Height float64
	// Weight вес
	Weight float64
}

// Manufacturer описывает производителя детали
type Manufacturer struct {
	// Name название производителя
	Name string
	// Country страна производителя
	Country string
	// Website URL сайта производителя
	Website string
}

// MetaDataValue описывает значение в метаданных
type MetaDataValue struct {
	StringValue *string
	DoubleValue *float64
	Int64Value  *int64
	BoolValue   *bool
}

// NewStringValue создает Value со строковым значением
func NewStringValue(v string) *MetaDataValue {
	return &MetaDataValue{StringValue: &v}
}

// NewDoubleValue создает Value со значением с плавающей точкой
func NewDoubleValue(v float64) *MetaDataValue {
	return &MetaDataValue{DoubleValue: &v}
}

// NewInt64Value создает Value с целочисленным значением
func NewInt64Value(v int64) *MetaDataValue {
	return &MetaDataValue{Int64Value: &v}
}

// NewBoolValue создает Value с булевым значением
func NewBoolValue(v bool) *MetaDataValue {
	return &MetaDataValue{BoolValue: &v}
}

// GetStringValue возвращает строковое значение
func (v *MetaDataValue) GetStringValue() string {
	if v != nil && v.StringValue != nil {
		return *v.StringValue
	}
	return ""
}

// GetDoubleValue возвращает значение с плавающей точкой
func (v *MetaDataValue) GetDoubleValue() float64 {
	if v != nil && v.DoubleValue != nil {
		return *v.DoubleValue
	}
	return 0
}

// GetInt64Value возвращает целочисленное значение
func (v *MetaDataValue) GetInt64Value() int64 {
	if v != nil && v.Int64Value != nil {
		return *v.Int64Value
	}
	return 0
}

// GetBoolValue возвращает булево значение
func (v *MetaDataValue) GetBoolValue() bool {
	if v != nil && v.BoolValue != nil {
		return *v.BoolValue
	}
	return false
}

// IsString проверяет, является ли значение строкой
func (v *MetaDataValue) IsString() bool {
	return v != nil && v.StringValue != nil
}

// IsDouble проверяет, является ли значение числом с плавающей точкой
func (v *MetaDataValue) IsDouble() bool {
	return v != nil && v.DoubleValue != nil
}

// IsInt64 проверяет, является ли значение целым числом
func (v *MetaDataValue) IsInt64() bool {
	return v != nil && v.Int64Value != nil
}

// IsBool проверяет, является ли значение булевым
func (v *MetaDataValue) IsBool() bool {
	return v != nil && v.BoolValue != nil
}
