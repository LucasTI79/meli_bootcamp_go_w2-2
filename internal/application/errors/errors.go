package errors

import "errors"

// Errors
var (
	ErrNotFound = errors.New("locality not found")
	ErrConflict = errors.New("ID already exists")
)
