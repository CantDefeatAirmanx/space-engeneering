package repository_part

import "fmt"

type ErrPartNotFound struct {
	Err  error
	UUID string
}

func (e *ErrPartNotFound) Error() string {
	return fmt.Sprintf("part not found: %s", e.UUID)
}

func (e *ErrPartNotFound) Unwrap() error {
	return e.Err
}

func (e *ErrPartNotFound) Is(target error) bool {
	_, ok := target.(*ErrPartNotFound)
	return ok
}

type ErrPartInternal struct {
	Err  error
	UUID string
}

func (e *ErrPartInternal) Error() string {
	return fmt.Sprintf("internal part %s error", e.UUID)
}

func (e *ErrPartInternal) Unwrap() error {
	return e.Err
}

func (e *ErrPartInternal) Is(target error) bool {
	_, ok := target.(*ErrPartInternal)
	return ok
}
