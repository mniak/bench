package bench

import "errors"

var (
	ErrProgramNotFound = errors.New("program not found")
	ErrSourceNotFound  = errors.New("source not found")
)
