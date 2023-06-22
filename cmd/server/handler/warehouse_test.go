package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/warehousesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/warehouse/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		warehouseFound := &domain.Warehouse{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(warehouseFound, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

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
		assert.Equal(t, *&warehouseFound, responseWarehouse)

	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Warehouse{}, warehouse.ErrNotFound)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetOne", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Warehouse{}, assert.PanicMatches)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create_Ok", func(t *testing.T) {
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

		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Warehouse")).Return(expectedWarehouse, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		req.GetBody()
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		bodyReturn, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Warehouse `json:"data"`
		}

		json.Unmarshal(bodyReturn, &responseDTO)

		actualWarehouse := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedWarehouse, actualWarehouse)
	})
	t.Run("Create BadRequest", func(t *testing.T) {
		createWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/warehouses", handler.Create())

		requestBody, _ := json.Marshal(createWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/warehouses", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		type expectedMensageResponseDTO struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		expectedMensageResponse := expectedMensageResponseDTO{
			Code:    "bad_request",
			Message: "Field Warehouse Code is required: ",
		}

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var actualMessageResponse expectedMensageResponseDTO
		json.Unmarshal(body, &actualMessageResponse)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedMensageResponse, actualMessageResponse)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("Get All", func(t *testing.T) {
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
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(warehousesFounds, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Warehouse `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		respondeWarehouses := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, warehousesFounds, respondeWarehouses)
	})

	t.Run("empty database", func(t *testing.T) {
		warehousesFounds := &[]domain.Warehouse{}
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(warehousesFounds, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/warehouses", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Deleted_sucess", func(t *testing.T) {
		warehouseFound := &domain.Warehouse{}
		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(warehouseFound, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/warehouses/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("Deleted_error", func(t *testing.T) {

		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(warehouse.ErrNotFound, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/warehouses/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
func TestUpdate(t *testing.T) {
	t.Run("Update_Ok", func(t *testing.T) {
		expectedWarehouse := &domain.Warehouse{
			ID:                 1,
			Address:            "Rua Teste",
			Telephone:          "11938473125",
			WarehouseCode:      "CX-2281-TCD",
			MinimumCapacity:    12,
			MinimumTemperature: 18,
		}
		updateWarehouseRequestDTO := dtos.WarehouseRequestDTO{
			Telephone:     "11938",
			WarehouseCode: "CX-1206-TCD",
		}

		//Configurar o mock do service
		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(expectedWarehouse, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/warehouses/:id", handler.Update())

		requestBody, _ := json.Marshal(updateWarehouseRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", request)
		req.GetBody()
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		bodyReturn, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Warehouse `json:"data"`
		}

		json.Unmarshal(bodyReturn, &responseDTO)

		actualWarehouse := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *expectedWarehouse, actualWarehouse)
	})

	t.Run("Update_Error", func(t *testing.T) {

		warehouseServiceMock := new(mocks.WarehouseServiceMock)
		warehouseServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(warehouse.ErrNotFound, nil)
		handler := handler.NewWarehouse(warehouseServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/warehouses/:id", handler.Update())

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/warehouses/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})
}
