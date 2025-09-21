package model_ship_assembly

import "errors"

var (
	ErrAssemblyNotFound         = errors.New("assembly not found")
	ErrAssemblyInvalidArguments = errors.New("assembly invalid arguments")
	ErrAssemblyConflict         = errors.New("assembly conflict")
)
