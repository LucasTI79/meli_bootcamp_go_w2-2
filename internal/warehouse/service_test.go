package warehouse_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/warehousesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		expectedWarehouse := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedWarehouse, nil)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseReceived, err := service.GetOne(&ctx, 1)

		assert.Equal(t, *expectedWarehouse, *warehouseReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Warehouse{}, sql.ErrNoRows)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseReceived, err := service.GetOne(&ctx, 1)

		assert.Nil(t, warehouseReceived)
		assert.Equal(t, warehouse.ErrNotFound, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		expectedWarehouses := &[]domain.Warehouse{
			{
				ID:                 1,
				Address:            "Rua Teste2",
				Telephone:          "11938473322",
				WarehouseCode:      "CX-2281-TCD",
				MinimumCapacity:    12,
				MinimumTemperature: 18,
			},
			{
				ID:                 1,
				Address:            "Rua Teste2",
				Telephone:          "11938473322",
				WarehouseCode:      "CX-2281-TCD",
				MinimumCapacity:    12,
				MinimumTemperature: 18,
			},
		}

		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("GetAll", ctx).Return(*expectedWarehouses, nil)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehousesReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedWarehouses, *warehousesReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("GetAll", ctx).Return([]domain.Warehouse{}, errors.New("error"))

		service := warehouse.NewService(warehouseRepositoryMock)
		warehousesReceived, err := service.GetAll(&ctx)

		assert.Nil(t, warehousesReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		expectedWarehouse := &domain.Warehouse{

			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, 1).Return(expectedWarehouse, nil)
		warehouseRepositoryMock.On("Delete", ctx, 1).Return(nil)

		service := warehouse.NewService(warehouseRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, nil, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Delete", ctx, 1).Return(warehouse.ErrNotFound)

		service := warehouse.NewService(warehouseRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, warehouse.ErrNotFound, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehousetSaved, err := service.Create(&ctx, createWarehouseRequestDTO)

		assert.Equal(t, warehouse.ErrConflict, err)
		assert.Nil(t, warehousetSaved)

	})

	t.Run("create_error", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()

		warehousetRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehousetRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehousetRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Warehouse")).Return(0, errors.New("error"))

		service := warehouse.NewService(warehousetRepositoryMock)
		warehouseSaved, err := service.Create(&ctx, createWarehouseRequestDTO)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, warehouseSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedWarehouse := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()

		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Warehouse")).Return(1, nil)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseSaved, err := service.Create(&ctx, createWarehouseRequestDTO)

		assert.Equal(t, warehouseSaved, expectedWarehouse)
		assert.Nil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		originalWarehouse := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		updateWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            originalWarehouse.Address,
			Telephone:          originalWarehouse.Telephone,
			WarehouseCode:      originalWarehouse.WarehouseCode,
			MinimumCapacity:    originalWarehouse.MinimumCapacity,
			MinimumTemperature: originalWarehouse.MinimumTemperature,
		}

		ctx := context.TODO()
		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalWarehouse, nil)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Warehouse")).Return(nil)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseUpdated, err := service.Update(&ctx, 1, updateWarehouseRequestDTO)

		assert.Equal(t, originalWarehouse, warehouseUpdated)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		updateWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()
		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Warehouse{}, warehouse.ErrNotFound)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseUpdated, err := service.Update(&ctx, 1, updateWarehouseRequestDTO)

		assert.Equal(t, warehouse.ErrNotFound, err)
		assert.Nil(t, warehouseUpdated)
	})

	t.Run("update_unexpected_error", func(t *testing.T) {
		updateWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()
		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Warehouse{}, warehouse.ErrNotFound)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Warehouse")).Return(errors.New(""))

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseUpdated, err := service.Update(&ctx, 1, updateWarehouseRequestDTO)
		assert.Nil(t, warehouseUpdated)
		assert.Error(t, err)
	})

	t.Run("update_get_conflit_error", func(t *testing.T) {
		updateWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()
		warehouseRepositoryMock := new(mocks.WarehouseRepositoryMock)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Warehouse{}, warehouse.ErrNotFound)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := warehouse.NewService(warehouseRepositoryMock)
		warehouseUpdated, err := service.Update(&ctx, 1, updateWarehouseRequestDTO)
		assert.Nil(t, warehouseUpdated)
		assert.Error(t, err)
	})
}
