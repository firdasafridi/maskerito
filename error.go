package maskerito

import "errors"

var (
	ErrInputNil  = errors.New("input is nil")
	ErrConfigNil = errors.New("config is nil")
)
