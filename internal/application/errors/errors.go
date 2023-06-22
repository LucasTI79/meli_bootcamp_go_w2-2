package errors

import(
	"errors"
)

var (
	ErrNotFound        = errors.New("section not found")
	ErrInvalidId       = errors.New("invalid id")
	ErrTryAgain        = errors.New("error, try again %s")
	ErrAlreadyExists   = errors.New("section already exists")
	ErrModifySection   = errors.New("cannot modify Section")
	ErrGettingSectionById = errors.New("error getting Section by id")
	ErrDeletingSection = errors.New("error deleting section")
	ErrNoRows = errors.New("sql: no rows in result set")
)
