package utils

import (
	"errors"
	"fmt"
)

var (
	ErrZeroDiv      = errors.New("")
	ErrIndex        = fmt.Errorf("")
	ErrTypeMismatch = errors.New("")
)
