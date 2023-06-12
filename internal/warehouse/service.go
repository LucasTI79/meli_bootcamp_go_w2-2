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
	Create(c context.Context,dto domain.WarehouseRequestDTO) (*domain.Warehouse, error)
}

type service struct{
	repository Repository
}

func (s *service) Create(c context.Context, dto domain.WarehouseRequestDTO) (*domain.Warehouse, error) {
	exists := s.repository.Exists(c, dto.WarehouseCode)
	if exists {
		return nil, errors.New("já existe um warehouse com este código")
	}

	var formatter domain.Warehouse = domain.Warehouse{ID: 0, Address: dto.Address, Telephone: dto.Telephone, WarehouseCode: dto.WarehouseCode, MinimumCapacity: dto.MinimumCapacity, MinimumTemperature: dto.MinimumTemperature}

	id, err := s.repository.Save(c, formatter)

	if err != nil {
		return nil, err
	}

	formatter.ID = id

	return &formatter, nil
}

func NewService(r Repository) Service {
	return &service{ repository: r }
}
