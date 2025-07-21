package repository_order

import (
	"errors"
	"fmt"
)

var ErrOrder = errors.New("order error")

type ErrOrderNotFound struct {
	OrderID string
	Err     error
}

func (e *ErrOrderNotFound) Error() string {
	return fmt.Sprintf("order %s not found. %v", e.OrderID, e.Err)
}

func (e *ErrOrderNotFound) Unwrap() error {
	return e.Err
}

func (e *ErrOrderNotFound) Is(target error) bool {
	_, ok := target.(*ErrOrderNotFound)
	return ok
}

type ErrOrderInternal struct {
	Err error
}

func (e *ErrOrderInternal) Error() string {
	return fmt.Sprintf("order internal error. %v", e.Err)
}

func (e *ErrOrderInternal) Unwrap() error {
	return e.Err
}

func (e *ErrOrderInternal) Is(target error) bool {
	_, ok := target.(*ErrOrderInternal)
	return ok
}
