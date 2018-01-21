package error

import "errors"

var (
	ErrUserUnauthorized = errors.New("user unauthorized")
)
