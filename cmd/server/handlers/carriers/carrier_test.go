package carriers_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		carriersFounds := &[]domain.Carrier{
			{
				ID:          1,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6700,
			},
			{
				ID:          2,
				CID:         "CID#1",
				CompanyName: "some name",
				Address:     "corrientes 800",
				Telephone:   "4567-4567",
				LocalityId:  6701,
			},
		}
		carrierServiceMock := new(mocks.CarriersServiceMock)
		carrierServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(carriersFounds, nil)
		handler := carrier_handler.NewWarehouse(carrierServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/carriers", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/carriers", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Carrier `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseCarriers := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *carriersFounds, responseCarriers)
	})

	t.Run("empty_database", func(t *testing.T) {
		warehousesFounds := &[]domain.Warehouse{}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(warehousesFounds, nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		warehousesFounds := &[]domain.Warehouse{}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(warehousesFounds, assert.AnError)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusInternalServerError, res.Code)

	})
}
