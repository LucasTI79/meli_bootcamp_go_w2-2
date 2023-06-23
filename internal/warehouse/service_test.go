package warehouse_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	expectedWarehouses := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		},
		{
			ID:                 1,
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		},
	}

	tests := []struct {
		name string
		//Mocking repository.GetAll
		expectedGetAllResult []domain.Warehouse
		expectedGetAllError  error
		expectedGetAllCalls  int
		//Asserting function
		expectedWarehouses *[]domain.Warehouse
		expectedError      error
	}{
		{
			name:                 "Successfully get all Warehouses from db",
			expectedGetAllResult: expectedWarehouses,
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedWarehouses:   &expectedWarehouses,
			expectedError:        nil,
		},
		{
			name:                 "Error connecting db",
			expectedGetAllResult: []domain.Warehouse{},
			expectedGetAllError:  assert.AnError,
			expectedGetAllCalls:  1,
			expectedWarehouses:   &[]domain.Warehouse{},
			expectedError:        assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock()
			warehouseRepositoryMock.On("GetAll", ctx).Return(test.expectedGetAllResult, test.expectedGetAllError)

			service := warehouse.NewService(warehouseRepositoryMock)
			WarehouseReceived, err := service.GetAll(&ctx)

			assert.Equal(t, *test.expectedWarehouses, *WarehouseReceived)
			assert.Equal(t, test.expectedError, err)

			warehouseRepositoryMock.AssertNumberOfCalls(t, "GetAll", test.expectedGetAllCalls)
		})
	}
}
