package mocks

import (
	"context"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/warehousesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type WarehouseServiceMock struct {
	mock.Mock
}

func (service *WarehouseServiceMock) GetOne(ctx *context.Context, id int) (*domain.Warehouse, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (service *WarehouseServiceMock) GetAll(ctx *context.Context) (*[]domain.Warehouse, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.Warehouse), args.Error(1)
}

func (service *WarehouseServiceMock) Create(ctx *context.Context, warehouse dtos.WarehouseRequestDTO) (*domain.Warehouse, error) {
	args := service.Called(ctx, warehouse)

	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (service *WarehouseServiceMock) Update(ctx *context.Context, id int, updateWarehouseRequest dtos.WarehouseRequestDTO) (*domain.Warehouse, error) {
	args := service.Called(ctx, updateWarehouseRequest)

	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (service *WarehouseServiceMock) Delete(ctx *context.Context, id int) error {
	args := service.Called(ctx, id)

	return args.Error(1)
}
