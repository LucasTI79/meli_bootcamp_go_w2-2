package services_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_warehouseService_Get(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		expectedWarehouse := &entities.Warehouse{
			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		ctx := context.TODO()

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedWarehouse, nil)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehouseReceived, err := service.GetOne(&ctx, 1)

		assert.Equal(t, *expectedWarehouse, *warehouseReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Warehouse{}, sql.ErrNoRows)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehouseReceived, err := service.GetOne(&ctx, 1)

		assert.Nil(t, warehouseReceived)
		assert.Equal(t, services.ErrNotFound, err)
	})
}

func Test_warehouseService_GetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		expectedWarehouses := &[]entities.Warehouse{
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

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("GetAll", ctx).Return(*expectedWarehouses, nil)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehousesReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedWarehouses, *warehousesReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("GetAll", ctx).Return([]entities.Warehouse{}, errors.New("error"))

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehousesReceived, err := service.GetAll(&ctx)

		assert.Nil(t, warehousesReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func Test_warehouseService_Delete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		expectedWarehouse := &entities.Warehouse{

			ID:                 1,
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		ctx := context.TODO()

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, 1).Return(expectedWarehouse, nil)
		warehouseRepositoryMock.On("Delete", ctx, 1).Return(nil)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, nil, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Delete", ctx, 1).Return(services.ErrNotFound)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, services.ErrNotFound, err)
	})
}

func Test_warehouseService_Create(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste2",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		ctx := context.TODO()

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehousetSaved, err := service.Create(&ctx, createWarehouseRequestDTO)

		assert.Equal(t, services.ErrConflict, err)
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

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Warehouse")).Return(0, errors.New("error"))

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehouseSaved, err := service.Create(&ctx, createWarehouseRequestDTO)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, warehouseSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedWarehouse := &entities.Warehouse{
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

		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Warehouse")).Return(1, nil)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehouseSaved, err := service.Create(&ctx, createWarehouseRequestDTO)

		assert.Equal(t, warehouseSaved, expectedWarehouse)
		assert.Nil(t, err)
	})
}

func Test_warehouseService_Update(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		originalWarehouse := &entities.Warehouse{
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
		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalWarehouse, nil)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Warehouse")).Return(nil)

		service := services.NewWarehouseService(warehouseRepositoryMock)
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
		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Warehouse{}, services.ErrNotFound)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehouseUpdated, err := service.Update(&ctx, 1, updateWarehouseRequestDTO)

		assert.Equal(t, services.ErrNotFound, err)
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
		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Warehouse{}, services.ErrNotFound)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		warehouseRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Warehouse")).Return(errors.New(""))

		service := services.NewWarehouseService(warehouseRepositoryMock)
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
		warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock(t)
		warehouseRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Warehouse{}, services.ErrNotFound)
		warehouseRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewWarehouseService(warehouseRepositoryMock)
		warehouseUpdated, err := service.Update(&ctx, 1, updateWarehouseRequestDTO)
		assert.Nil(t, warehouseUpdated)
		assert.Error(t, err)
	})
}
