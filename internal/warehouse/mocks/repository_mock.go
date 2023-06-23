package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type WarehouseRepositoryMock struct {
	mock.Mock
}

func NewWarehouseRepositoryMock() *WarehouseRepositoryMock {
	return &WarehouseRepositoryMock{}
}

func (repository *WarehouseRepositoryMock) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *WarehouseRepositoryMock) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	args := repository.Called(ctx)

	return args.Get(0).(domain.Warehouse), args.Error(1)
}

func (repository *WarehouseRepositoryMock) Exists(ctx context.Context, warehouseCode string) bool {
	//TODO implement me
	panic("implement me")
}

func (repository *WarehouseRepositoryMock) Save(ctx context.Context, b domain.Warehouse) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *WarehouseRepositoryMock) Update(ctx context.Context, b domain.Warehouse) error {
	//TODO implement me
	panic("implement me")
}

func (repository *WarehouseRepositoryMock) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")

}
