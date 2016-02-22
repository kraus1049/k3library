package k3library

import (
	"errors"
)

var (
	ErrInvalid      = errors.New("Invalid argument")
	ErrInfiniteLoop = errors.New("Infinite loop")
)
