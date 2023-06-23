package warehouse_test

import (
	"database/sql"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	warehouseFound := domain.Warehouse{
		Address:            "Rua Teste",
		Telephone:          "11938473125",
		WarehouseCode:      "CX-2281-TCD",
		MinimumCapacity:    12,
		MinimumTemperature: 18,
	}
	tests := []struct {
		name string
		id   int
		//Mocking repository.Get
		expectedGetResult domain.Warehouse
		expectedGetError  error
		expectedGetCalls  int
		//Asserting function
		warehouseFound *domain.Warehouse
		expectedError  error
	}{
		{
			name:              "Successfully get warehouse from db",
			id:                1,
			expectedGetResult: warehouseFound,
			expectedGetError:  nil,
			expectedGetCalls:  1,
			warehouseFound:    &warehouseFound,
			expectedError:     nil,
		},
		{
			name:              "Warehouse not found in db",
			id:                1,
			expectedGetResult: domain.Warehouse{},
			expectedGetError:  sql.ErrNoRows,
			expectedGetCalls:  1,
			warehouseFound:    &domain.Warehouse{},
			expectedError:     warehouse.ErrNotFound,
		},
		{
			name:              "Error connecting db",
			id:                1,
			expectedGetResult: domain.Warehouse{},
			expectedGetError:  assert.AnError,
			expectedGetCalls:  1,
			warehouseFound:    &domain.Warehouse{},
			expectedError:     assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			ctx := c.Request.Context()

			warehouseRepositoryMock := mocks.NewWarehouseRepositoryMock()
			warehouseRepositoryMock.On("Get", mock.AnythingOfType("context.Context"), mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)

			service := warehouse.NewService(warehouseRepositoryMock)
			warehouseReceived, err := service.GetOne(&ctx, test.id)

			assert.Equal(t, test.warehouseFound, *warehouseReceived)
			assert.Equal(t, test.expectedError, err)

			warehouseRepositoryMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
		})
	}
}
