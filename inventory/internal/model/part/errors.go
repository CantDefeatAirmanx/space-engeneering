package model_part

import "fmt"

type ErrPartNotFound struct {
	UUID string
	Err  error
}

func (e ErrPartNotFound) Error() string {
	return fmt.Sprintf("Part not found: %s", e.UUID)
}

func (e ErrPartNotFound) Unwrap() error {
	return e.Err
}

func (e ErrPartNotFound) Is(target error) bool {
	_, ok := target.(*ErrPartNotFound)
	return ok
}

type ErrPartInternal struct {
	UUID string
	Err  error
}

func (e ErrPartInternal) Error() string {
	return fmt.Sprintf("Part internal error")
}

func (e ErrPartInternal) Unwrap() error {
	return e.Err
}

func (e ErrPartInternal) Is(target error) bool {
	_, ok := target.(*ErrPartInternal)
	return ok
}
