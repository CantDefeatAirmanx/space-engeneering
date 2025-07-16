package model_part

import "fmt"

type ErrPartNotFound struct {
	UUID string
	Err  error
}

var ErrPartNotFoundInstance = &ErrPartNotFound{}

func (e ErrPartNotFound) Error() string {
	return fmt.Sprintf("%v. Part not found: %s", e.Err, e.UUID)
}

func (e ErrPartNotFound) Unwrap() error {
	return e.Err
}

type ErrPartInternal struct {
	UUID string
	Err  error
}

func (e ErrPartInternal) Error() string {
	return fmt.Sprintf("%v. Part internal error", e.Err)
}

func (e ErrPartInternal) Unwrap() error {
	return e.Err
}
