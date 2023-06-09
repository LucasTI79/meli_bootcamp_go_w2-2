package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("warehouse not found")
)

type Service interface{
	Create(dto domain.WarehouseRequestDTO) (domain.Warehouse, error)
}

type service struct{
	repository 
}

func (s *service) Create(c context.Context,  dto domain.Warehouse) (domain.Warehouse, error) {
	result, err := s.repository.Save(c, dto)
	return result, err
}

func NewService(r Repository) Service {
	return &service{}
}
