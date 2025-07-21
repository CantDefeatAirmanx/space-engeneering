package model_order

import (
	"fmt"
)

type ErrOrderNotFound struct {
	OrderUUID string
	Err       error
}

func (e *ErrOrderNotFound) Error() string {
	return fmt.Sprintf("order not found: %s", e.OrderUUID)
}

func (e *ErrOrderNotFound) Unwrap() error {
	return e.Err
}

func (e *ErrOrderNotFound) Is(target error) bool {
	_, ok := target.(*ErrOrderNotFound)
	return ok
}

type ErrOrderConflict struct {
	OrderUUID  string
	Err        error
	ErrMessage string
}

func (e *ErrOrderConflict) Error() string {
	return fmt.Sprintf("order conflict: %s, %s", e.OrderUUID, e.ErrMessage)
}

func (e *ErrOrderConflict) Unwrap() error {
	return e.Err
}

func (e *ErrOrderConflict) Is(target error) bool {
	_, ok := target.(*ErrOrderConflict)
	return ok
}

type ErrOrderInvalidArguments struct {
	OrderUUID string
	Err       error
}

func (e *ErrOrderInvalidArguments) Error() string {
	return fmt.Sprintf("order invalid arguments: %s", e.OrderUUID)
}

func (e *ErrOrderInvalidArguments) Unwrap() error {
	return e.Err
}

func (e *ErrOrderInvalidArguments) Is(target error) bool {
	_, ok := target.(*ErrOrderInvalidArguments)
	return ok
}

type ErrOrderInternal struct {
	OrderUUID string
	Err       error
}

func (e *ErrOrderInternal) Error() string {
	return fmt.Sprintf("order internal: %s", e.OrderUUID)
}

func (e *ErrOrderInternal) Unwrap() error {
	return e.Err
}

func (e *ErrOrderInternal) Is(target error) bool {
	_, ok := target.(*ErrOrderInternal)
	return ok
}
