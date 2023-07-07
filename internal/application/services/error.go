package services

import "errors"

var (
	ErrNotFound            = errors.New("locality not found")
	ErrConflict            = errors.New("ID already exists")
	ErrUnprocessableEntity = errors.New("error processing entity")
)
