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

type Service interface {
	Create(c context.Context, dto dtos.WarehouseRequestDTO) (*domain.Warehouse, error)
	GetAll(c context.Context) (*[]domain.Warehouse, error)
	GetOne(c context.Context, id int) (*domain.Warehouse, error)
	Update(c context.Context, id int, dto dtos.WarehouseRequestDTO) (*domain.Warehouse, error)
	Delete(c context.Context, id int) error
}

type service struct {
	repository Repository
}

func (s *service) Create(c context.Context, dto dtos.WarehouseRequestDTO) (*domain.Warehouse, error) {
	exists := s.repository.Exists(c, dto.WarehouseCode)
	if exists {
		return nil, errors.New("a warehouse with this warehouse_code already exists")
	}

	var formatter domain.Warehouse = domain.Warehouse{
		ID:                 0,
		Address:            dto.Address,
		Telephone:          dto.Telephone,
		WarehouseCode:      dto.WarehouseCode,
		MinimumCapacity:    dto.MinimumCapacity,
		MinimumTemperature: dto.MinimumTemperature,
	}

	id, err := s.repository.Save(c, formatter)

	if err != nil {
		return nil, err
	}

	formatter.ID = id

	return &formatter, nil
}

func (s *service) GetAll(c context.Context) (*[]domain.Warehouse, error) {
	warehouses, err := s.repository.GetAll(c)

	if err != nil {
		return nil, err
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

func (s *service) Update(c context.Context, id int, dto dtos.WarehouseRequestDTO) (*domain.Warehouse, error) {
	newWarehouse, er := s.GetOne(c, id)

	if er != nil {
		return nil, ErrNotFound
	}

	newWarehouse = updateFormatter(dto, *newWarehouse)

	err := s.repository.Update(c, *newWarehouse)

	if err != nil {
		return nil, err
	}

	return newWarehouse, nil
}

func (s *service) Delete(c context.Context, id int) error {
	result := s.repository.Delete(c, id)

	if result != nil {
		return ErrNotFound
	}

	return nil
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func updateFormatter(dto dtos.WarehouseRequestDTO, newWarehouse domain.Warehouse) *domain.Warehouse {
	if dto.Address != "" {
		newWarehouse.Address = dto.Address
	}
	if dto.Telephone != "" {
		newWarehouse.Telephone = dto.Telephone
	}
	if dto.WarehouseCode != "" {
		newWarehouse.WarehouseCode = dto.WarehouseCode
	}
	if dto.MinimumCapacity != 0 {
		newWarehouse.MinimumCapacity = dto.MinimumCapacity
	}
	if dto.MinimumTemperature != 0 {
		newWarehouse.MinimumTemperature = dto.MinimumTemperature
	}

	return &newWarehouse
}
