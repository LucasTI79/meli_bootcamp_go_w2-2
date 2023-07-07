package services

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
)

type WarehouseService interface {
	Create(c *context.Context, dto dtos.WarehouseRequestDTO) (*entities.Warehouse, error)
	GetAll(c *context.Context) (*[]entities.Warehouse, error)
	GetOne(c *context.Context, id int) (*entities.Warehouse, error)
	Update(c *context.Context, id int, dto dtos.WarehouseRequestDTO) (*entities.Warehouse, error)
	Delete(c *context.Context, id int) error
}

type warehouseService struct {
	repository repositories.WarehouseRepository
}

func NewWarehouseService(r repositories.WarehouseRepository) WarehouseService {
	return &warehouseService{repository: r}
}

func (s *warehouseService) Create(c *context.Context, dto dtos.WarehouseRequestDTO) (*entities.Warehouse, error) {

	exists := s.repository.Exists(*c, dto.WarehouseCode)
	if exists {
		return nil, ErrConflict
	}

	var formatter entities.Warehouse = entities.Warehouse{
		ID:                 0,
		Address:            dto.Address,
		Telephone:          dto.Telephone,
		WarehouseCode:      dto.WarehouseCode,
		MinimumCapacity:    dto.MinimumCapacity,
		MinimumTemperature: dto.MinimumTemperature,
	}

	id, err := s.repository.Save(*c, formatter)

	if err != nil {
		return nil, err
	}

	formatter.ID = id

	return &formatter, nil
}

func (s *warehouseService) GetAll(c *context.Context) (*[]entities.Warehouse, error) {
	warehouses, err := s.repository.GetAll(*c)

	if err != nil {
		return nil, err
	}

	return &warehouses, nil
}

func (s *warehouseService) GetOne(c *context.Context, id int) (*entities.Warehouse, error) {
	result, err := s.repository.Get(*c, id)
	if err != nil {
		return nil, ErrNotFound
	}

	return &result, nil
}

func (s *warehouseService) Update(c *context.Context, id int, dto dtos.WarehouseRequestDTO) (*entities.Warehouse, error) {
	newWarehouse, er := s.GetOne(c, id)

	if er != nil {
		return nil, ErrNotFound
	}

	exists := s.repository.Exists(*c, dto.WarehouseCode)

	if exists {
		if newWarehouse.WarehouseCode != dto.WarehouseCode {
			return nil, ErrConflict
		}
	}

	newWarehouse = updateFormatter(dto, *newWarehouse)

	err := s.repository.Update(*c, *newWarehouse)

	if err != nil {
		return nil, err
	}

	return newWarehouse, nil
}

func (s *warehouseService) Delete(c *context.Context, id int) error {
	result := s.repository.Delete(*c, id)

	if result != nil {
		return ErrNotFound
	}

	return nil
}

func updateFormatter(dto dtos.WarehouseRequestDTO, newWarehouse entities.Warehouse) *entities.Warehouse {
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
