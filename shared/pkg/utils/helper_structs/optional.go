package helper_structs

type OptionalInterface[T any] interface {
	IsSet() bool
	GetValue() T
}

type Optional[T any] struct {
	value T
	isSet bool
}

func NewOptional[T any](value T) *Optional[T] {
	return &Optional[T]{
		value: value,
		isSet: true,
	}
}

func NewOptionalEmpty[T any]() *Optional[T] {
	var zero T

	return &Optional[T]{
		value: zero,
		isSet: false,
	}
}

func (o *Optional[T]) IsSet() bool {
	if o == nil {
		return false
	}

	return o.isSet
}

func (o *Optional[T]) GetValue() T {
	if o == nil {
		var zero T
		return zero
	}

	return o.value
}
