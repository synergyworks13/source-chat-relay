package protocol

import "errors"

var (
	ErrCannotReadString = errors.New("Cannot read string")
	ErrOutOfBound       = errors.New("Read out of bound")
)
