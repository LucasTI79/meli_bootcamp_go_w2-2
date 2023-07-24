package inbound_orders_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/inbound_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	inbound_order "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/inbound_order/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		inboundOrdersFound := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(inboundOrdersFound, nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/inboundOrders/:id", inboundOrders.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.InboundOrders `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseInboundOrders := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *inboundOrdersFound, responseInboundOrders)

	})

	// // find_by_id_non_existent
	// t.Run("find_by_id_non_existent", func(t *testing.T) {
	// 	//Configurar o mock do service
	// 	inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
	// 	inboundOrdersServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.InboundOrders{}, inbound_order.ErrNotFound)
	// 	inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

	// 	//Configurar o servidor
	// 	gin.SetMode(gin.TestMode)
	// 	r := gin.Default()
	// 	r.GET("/api/v1/inboundOrders/:id", inboundOrders.Get())

	// 	//Definir request e response
	// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders/1", nil)
	// 	res := httptest.NewRecorder()

	// 	//Executar request
	// 	r.ServeHTTP(res, req)

	// 	//Validar resultado
	// 	assert.Equal(t, http.StatusNotFound, res.Code)
	// })

	t.Run("invalid_id", func(t *testing.T) {
		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.InboundOrders{}, inbound_order.ErrNotFound)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/inboundOrders/:id", inboundOrders.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders/a", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.InboundOrders{}, assert.AnError)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/inboundOrders/:id", inboundOrders.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

}

func TestGetAll(t *testing.T) {
	t.Run("find_all", func(t *testing.T) {
		// Definir resultado da consulta
		expectedinboundOrders := &[]domain.InboundOrders{
			{
				ID:             1,
				OrderDate:      "teste",
				OrderNumber:    "teste",
				EmployeeID:     "teste",
				ProductBatchID: "teste",
				WarehouseID:    "teste",
			},
			{
				ID:             2,
				OrderDate:      "teste",
				OrderNumber:    "teste",
				EmployeeID:     "teste",
				ProductBatchID: "teste",
				WarehouseID:    "teste",
			},
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedinboundOrders, nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/inboundOrders", inboundOrders.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.InboundOrders `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseInboundOrders := responseDTO.Data

		//Validar resultado
		assert.Equal(t, *expectedinboundOrders, responseInboundOrders)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(&[]domain.InboundOrders{}, assert.AnError)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/inboundOrders", inboundOrders.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
	t.Run("find_no_content", func(t *testing.T) {
		inboundOrdersFound := &[]domain.InboundOrders{}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("GetAll", mock.Anything).Return(inboundOrdersFound, nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/inboundOrders", inboundOrders.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/inboundOrders", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

}

func TestSave(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {

		expectedInboundOrdersCreate := &domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		requestInboundOrdersCreate := &domain.RequestCreateInboundOrders{
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.InboundOrders")).Return(expectedInboundOrdersCreate, nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.InboundOrders `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		actualInboundOrders := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedInboundOrdersCreate, actualInboundOrders)

	})

	t.Run("create_fail", func(t *testing.T) {
		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Save", mock.AnythingOfType("*context.Context")).Return(requestInboundOrdersCreate, inbound_order.ErrUnprocessableEntity)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_conflit", func(t *testing.T) {
		expectedInboundOrdersCreate := &domain.InboundOrders{}

		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.InboundOrders")).Return(expectedInboundOrdersCreate, inbound_order.ErrConflict)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusConflict, res.Code)
	})

	t.Run("create_fail_order_date_nil", func(t *testing.T) {
		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			// OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_order_number_nil", func(t *testing.T) {
		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate: "teste",
			// OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_employee_id_nil", func(t *testing.T) {
		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate:   "teste",
			OrderNumber: "teste",
			// EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_product_batch_id_nil", func(t *testing.T) {
		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate:   "teste",
			OrderNumber: "teste",
			EmployeeID:  "teste",
			// ProductBatchID: "teste",
			WarehouseID: "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
	t.Run("create_fail_warehouse_id_nil", func(t *testing.T) {
		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			// WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_internal_server_error", func(t *testing.T) {
		inboundOrdersCreate := domain.InboundOrders{
			ID:             1,
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
			OrderDate:      "teste",
			OrderNumber:    "teste",
			EmployeeID:     "teste",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.InboundOrders")).Return(inboundOrdersCreate, assert.AnError)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/inboundOrders", inboundOrders.Save())

		requestBody, _ := json.Marshal(requestInboundOrdersCreate)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.InboundOrders `json:"data"`
		}
		json.Unmarshal(body, &responseDTO)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	// t.Run("create_bad_request", func(t *testing.T) {
	// 	requestInboundOrdersCreate := domain.RequestCreateInboundOrders{
	// 		OrderDate:      "teste",
	// 		OrderNumber:    "teste",
	// 		EmployeeID:     "teste",
	// 		ProductBatchID: "teste",
	// 		WarehouseID:    "teste",
	// 	}

	// 	//Configurar o mock do service
	// 	inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
	// 	inboundOrdersServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.InboundOrders")).Return(&domain.InboundOrders{}, assert.AnError)
	// 	inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

	// 	gin.SetMode(gin.TestMode)
	// 	r := gin.Default()
	// 	r.POST("/api/v1/inboundOrders", inboundOrders.Save())

	// 	requestBody, _ := json.Marshal(requestInboundOrdersCreate)
	// 	request := bytes.NewReader(requestBody)

	// 	req := httptest.NewRequest(http.MethodPost, "/api/v1/inboundOrders", request)
	// 	res := httptest.NewRecorder()

	// 	r.ServeHTTP(res, req)

	// 	body, _ := ioutil.ReadAll(res.Body)

	// 	var responseDTO struct {
	// 		Data domain.InboundOrders `json:"data"`
	// 	}
	// 	json.Unmarshal(body, &responseDTO)
	// 	assert.Equal(t, http.StatusBadRequest, res.Code)
	// })
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		inboundOrdersFound := &domain.InboundOrders{}

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(inboundOrdersFound, nil)
		inboundOrdersServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/inboundOrders/:id", inboundOrders.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/inboundOrders/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("delete_non_existent", func(t *testing.T) {
		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(domain.InboundOrders{}, nil)
		inboundOrdersServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(inbound_order.ErrNotFound)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/inboundOrders/:id", inboundOrders.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/inboundOrders/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("delete_error_parsing_id", func(t *testing.T) {

		//Configurar o mock do service
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/inboundOrders/:id", inboundOrders.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/inboundOrders/ww", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {

		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		inboundOrdersUpdated := domain.InboundOrders{
			ID:             1,
			OrderDate:      "updated",
			OrderNumber:    "updated",
			EmployeeID:     "updated",
			ProductBatchID: "teste",
			WarehouseID:    "teste",
		}

		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrdersServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("*domain.RequestUpdateInboundOrders")).Return(&inboundOrdersUpdated, nil)
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/inboundOrders/:id", inboundOrders.Update())

		requestBody, _ := json.Marshal(requestUpdateInboundOrders)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/inboundOrders/1", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.InboundOrders `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseInboundOrders := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, inboundOrdersUpdated, *responseInboundOrders)
	})

	t.Run("update_unprocessable_entity", func(t *testing.T) {
		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/inboundOrders/:id", inboundOrders.Update())

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/inboundOrders/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	// t.Run("update_not_found", func(t *testing.T) {
	// 	orderDate := "updated"
	// 	orderNumber := "updated"
	// 	employeeID := "updated"

	// 	requestUpdateInboundOrders := domain.RequestUpdateInboundOrders{
	// 		OrderDate:   &orderDate,
	// 		OrderNumber: &orderNumber,
	// 		EmployeeID:  &employeeID,
	// 	}

	// 	inboundOrdersUpdated := &domain.InboundOrders{}

	// 	inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
	// 	inboundOrdersServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("*domain.requestUpdateInboundOrders")).Return(inboundOrdersUpdated, inbound_order.ErrNotFound)
	// 	inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

	// 	//Configurar o servidor
	// 	gin.SetMode(gin.TestMode)
	// 	r := gin.Default()
	// 	r.PATCH("/api/v1/inboundOrders/:id", inboundOrders.Update())

	// 	requestBody, _ := json.Marshal(requestUpdateInboundOrders)
	// 	request := bytes.NewReader(requestBody)

	// 	req := httptest.NewRequest(http.MethodPatch, "/api/v1/inboundOrders/1", request)
	// 	res := httptest.NewRecorder()

	// 	r.ServeHTTP(res, req)

	// 	assert.Equal(t, http.StatusNotFound, res.Code)
	// })

	t.Run("update_status_bad_request", func(t *testing.T) {
		orderDate := "updated"
		orderNumber := "updated"
		employeeID := "updated"

		requestUpdateInboundOrders := domain.RequestUpdateInboundOrders{
			OrderDate:   &orderDate,
			OrderNumber: &orderNumber,
			EmployeeID:  &employeeID,
		}

		inboundOrdersServiceMock := mocks.NewInboundOrdersServiceMock()
		inboundOrders := inbound_orders.NewInboundOrders(inboundOrdersServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/inboundOrders/:id", inboundOrders.Update())

		requestBody, _ := json.Marshal(requestUpdateInboundOrders)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/inboundOrders/xx", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
