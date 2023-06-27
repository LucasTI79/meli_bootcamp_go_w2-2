package warehouses_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	warehouse_handler "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/warehouses"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/warehousesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		warehouseFound := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(warehouseFound, nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Warehouse `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseWarehouse := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *warehouseFound, responseWarehouse)

	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Warehouse{}, warehouse.ErrNotFound)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses/:id", handler.Get())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("get_invalid_id", func(t *testing.T) {

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Warehouse{}, assert.AnError)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses/:id", handler.Get())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/xyz", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)

	})
	t.Run("not_found_error", func(t *testing.T) {

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Warehouse{}, assert.AnError)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses/:id", handler.Get())

		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}

func TestCreate(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		expectedWarehouse := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("dtos.WarehouseRequestDTO")).Return(expectedWarehouse, nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		bodyReturn, _ := ioutil.ReadAll(res.Body)
		var responseDTO struct {
			Data *domain.Warehouse `json:"data"`
		}
		json.Unmarshal(bodyReturn, &responseDTO)
		actualWarehouse := responseDTO.Data

		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedWarehouse, *actualWarehouse)
	})
	t.Run("create_fail", func(t *testing.T) {
		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Create", mock.AnythingOfType("*context.Context")).Return(createWarehouseRequestDTO, warehouse.ErrUnprocessableEntity)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
	t.Run("create_conflit", func(t *testing.T) {
		expectedWarehouse := &domain.Warehouse{}
		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("dtos.WarehouseRequestDTO")).Return(expectedWarehouse, warehouse.ErrConflict)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusConflict, res.Code)
	})
	t.Run("create_fail_address_nil", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_telephone_nil", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_warehouse_code_nil", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_minimum_capacity_nil", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumTemperature: 18,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_minimum_temperature_nil", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:         "Rua Teste",
			Telephone:       "11938473125",
			WarehouseCode:   "CX-2281-TCD",
			MinimumCapacity: 12,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_internal_server_error", func(t *testing.T) {

		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Warehouse")).Return(&domain.Warehouse{}, errors.New("error"))
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/Warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/Warehouses", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body2, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Warehouse `json:"data"`
		}
		json.Unmarshal(body2, &responseDTO)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		warehousesFounds := &[]domain.Warehouse{
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
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(warehousesFounds, nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses", handler.GetAll())
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Warehouse `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseWarehouses := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *warehousesFounds, responseWarehouses)
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

func TestDelete(t *testing.T) {
	t.Run("delete_delete_ok", func(t *testing.T) {

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/warehouses/:id", handler.Delete())

		req := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("delete_non_existent", func(t *testing.T) {
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(warehouse.ErrNotFound, nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/warehouses/:id", handler.Delete())
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})
	t.Run("delete_error_parsing_id", func(t *testing.T) {

		WarehouseServiceMock := new(mocks.WarehouseServiceMock)
		WarehouseServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		handler := warehouse_handler.NewWarehouse(WarehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/warehouses/:id", handler.Delete())
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/warehouses/xyz", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		address := "teste2"
		telephone := "232039"
		warehouseCode := "CX-2281-TCD"
		minimumCapacity := 12
		minimumTemperature := 10

		updateWarehouseRequest := dtos.WarehouseRequestDTO{
			Address:            address,
			Telephone:          telephone,
			WarehouseCode:      warehouseCode,
			MinimumCapacity:    minimumCapacity,
			MinimumTemperature: minimumTemperature,
		}
		updatedWarehouse := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste4",
			Telephone:          "11938473322",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 8,
		}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On(
			"Update", mock.Anything, mock.Anything, mock.Anything).Return(updatedWarehouse, nil)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/warehouses/:id", handler.Update())

		requestBody, _ := json.Marshal(updateWarehouseRequest)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Warehouse `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseWarehouse := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *updatedWarehouse, *responseWarehouse)
	})

	t.Run("update_unprocessable_entity", func(t *testing.T) {

		warehouseServiceMock := new(mocks.WarehouseServiceMock)

		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/warehouses/:id", handler.Update())

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("update_not_found", func(t *testing.T) {

		address := "teste2"
		telephone := "232039"
		warehouseCode := "CX-2281-TCD"
		minimumCapacity := 12
		minimumTemperature := 10

		updateWarehouseRequest := dtos.WarehouseRequestDTO{
			Address:            address,
			Telephone:          telephone,
			WarehouseCode:      warehouseCode,
			MinimumCapacity:    minimumCapacity,
			MinimumTemperature: minimumTemperature,
		}
		updatedWarehouse := &domain.Warehouse{}
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On(
			"Update", mock.Anything, mock.Anything, mock.Anything).Return(updatedWarehouse, warehouse.ErrNotFound)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/warehouses/:id", handler.Update())

		requestBody, _ := json.Marshal(updateWarehouseRequest)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("update_status_bad_request", func(t *testing.T) {
		address := "teste2"
		telephone := "232039"
		warehouseCode := "CX-2281-TCD"
		minimumCapacity := 12
		minimumTemperature := 10

		updateWarehouseRequest := dtos.WarehouseRequestDTO{
			Address:            address,
			Telephone:          telephone,
			WarehouseCode:      warehouseCode,
			MinimumCapacity:    minimumCapacity,
			MinimumTemperature: minimumTemperature,
		}

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := warehouse_handler.NewWarehouse(warehouseServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/warehouses/:id", handler.Update())

		requestBody, _ := json.Marshal(updateWarehouseRequest)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/a", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

}
