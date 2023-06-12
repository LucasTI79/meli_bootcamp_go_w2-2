package warehouse

import (
	"context"
	"errors"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/warehousesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("warehouse not found")
)

type Service interface{
	Create(c context.Context, dto dtos.WarehouseRequestDTO) (*domain.Warehouse, error)
	GetAll(c context.Context) (*[]domain.Warehouse, error)
	GetOne(c context.Context, id int) (*domain.Warehouse, error)
}

type service struct{
	repository Repository
}

func (s *service) Create(c context.Context, dto dtos.WarehouseRequestDTO) (*domain.Warehouse, error) {
	exists := s.repository.Exists(c, dto.WarehouseCode)
	if exists {
		return nil, errors.New("a warehouse with this warehouse_code already exists")
	}

	var formatter domain.Warehouse = domain.Warehouse{ID: 0, Address: dto.Address, Telephone: dto.Telephone, WarehouseCode: dto.WarehouseCode, MinimumCapacity: dto.MinimumCapacity, MinimumTemperature: dto.MinimumTemperature}

	id, err := s.repository.Save(c, formatter)

	if err != nil {
		return nil, err
	}

	formatter.ID = id

	return &formatter, nil
}

func (s *service) GetAll(c context.Context) (*[]domain.Warehouse, error) {
	warehouses, err  := s.repository.GetAll(c)

	if err != nil {
		return nil, err
	}

	if len(warehouses) < 1 {
		return nil, ErrNotFound
	}

	return &warehouses, nil
}

func (s *service) GetOne(c context.Context, id int) (*domain.Warehouse, error) {
	result, err := s.repository.Get(c, id)
	if err != nil {
		return nil, ErrNotFound
	}

	return &result, nil
}


func NewService(r Repository) Service {
	return &service{ repository: r }
}
