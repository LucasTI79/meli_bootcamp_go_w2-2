package domain

import "errors"

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

type Section struct {
	ID                 int `json:"id"`
	SectionNumber      int `json:"section_number"`
	CurrentTemperature int `json:"current_temperature"`
	MinimumTemperature int `json:"minimum_temperature"`
	CurrentCapacity    int `json:"current_capacity"`
	MinimumCapacity    int `json:"minimum_capacity"`
	MaximumCapacity    int `json:"maximum_capacity"`
	WarehouseID        int `json:"warehouse_id"`
	ProductTypeID      int `json:"product_type_id"`
}
