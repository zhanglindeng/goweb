package error

import "errors"

var (
	ErrNotJson = errors.New("content type is not application/json")
)
