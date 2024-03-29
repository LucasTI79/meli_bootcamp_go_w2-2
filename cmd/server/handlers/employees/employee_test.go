package employees_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/employees"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/employee/mocks"
	"github.com/gin-gonic/gin"

	//"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {
	t.Run("find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		employeeFound := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(employeeFound, nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", employees.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Employee `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseEmployee := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *employeeFound, responseEmployee)

	})
	// find_by_id_non_existent
	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Employee{}, employees.ErrNotFound)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", employees.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("invalid_id", func(t *testing.T) {
		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Employee{}, employees.ErrNotFound)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", employees.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/a", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Employee{}, assert.AnError)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", employees.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
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
		expectedEmployee := &[]domain.Employee{
			{
				ID:           1,
				CardNumberID: "123",
				FirstName:    "Maria",
				LastName:     "Silva",
				WarehouseID:  1,
			},
			{
				ID:           2,
				CardNumberID: "234",
				FirstName:    "Pedro",
				LastName:     "Silva",
				WarehouseID:  2,
			},
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedEmployee, nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees", employees.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Employee `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseEmployee := responseDTO.Data

		//Validar resultado
		assert.Equal(t, *expectedEmployee, responseEmployee)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(&[]domain.Employee{}, assert.AnError)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees", employees.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
	t.Run("find_no_content", func(t *testing.T) {
		employeeFound := &[]domain.Employee{}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("GetAll", mock.Anything).Return(employeeFound, nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees", employees.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

}

func TestSave(t *testing.T) {
	t.Run("create_ok", func(t *testing.T) {
		expectedEmployeeCreate := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Employee")).Return(expectedEmployeeCreate, nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Employee `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		actualEmployee := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedEmployeeCreate, actualEmployee)

	})

	t.Run("create_fail", func(t *testing.T) {
		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Save", mock.AnythingOfType("*context.Context")).Return(requestEmployeeCreated, employee.ErrUnprocessableEntity)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_conflit", func(t *testing.T) {
		expectedEmployeeCreate := &domain.Employee{}

		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Employee")).Return(expectedEmployeeCreate, employee.ErrConflict)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusConflict, res.Code)
	})

	t.Run("create_fail_card_number_id_nil", func(t *testing.T) {

		requestEmployeeCreated := domain.RequestCreateEmployee{
			FirstName:   "Maria",
			LastName:    "Silva",
			WarehouseID: 1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employees := employees.NewEmployee(employeeServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_first_name_nil", func(t *testing.T) {

		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employees := employees.NewEmployee(employeeServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_last_name_nil", func(t *testing.T) {

		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employees := employees.NewEmployee(employeeServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_fail_warehouse_id_nil", func(t *testing.T) {

		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employees := employees.NewEmployee(employeeServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("create_internal_server_error", func(t *testing.T) {
		employeeCreate := &domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Employee")).Return(*employeeCreate, assert.AnError)
		employees := employees.NewEmployee(employeeServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Warehouse `json:"data"`
		}
		json.Unmarshal(body, &responseDTO)
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("create_bad_request", func(t *testing.T) {
		requestEmployeeCreated := domain.RequestCreateEmployee{
			CardNumberID: "123",
			FirstName:    "Maria",
			LastName:     "Silva",
			WarehouseID:  1,
		}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Employee")).Return(&domain.Employee{}, assert.AnError)
		employees := employees.NewEmployee(employeeServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/employees", employees.Save())

		requestBody, _ := json.Marshal(requestEmployeeCreated)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Warehouse `json:"data"`
		}
		json.Unmarshal(body, &responseDTO)
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {
		employeeFound := &domain.Employee{}

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(employeeFound, nil)
		employeeServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/employees/:id", employees.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("delete_non_existent", func(t *testing.T) {
		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(domain.Employee{}, nil)
		employeeServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(employee.ErrNotFound)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/employees/:id", employees.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("delete_error_parsing_id", func(t *testing.T) {

		//Configurar o mock do service
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/employees/:id", employees.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/employees/ww", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {

		firstName := "teste"
		lastName := "Teste"

		RequestUpdateEmployee := domain.RequestUpdateEmployee{
			FirstName: &firstName,
			LastName:  &lastName,
		}

		employeeUpdated := domain.Employee{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "teste",
			LastName:     "Teste",
			WarehouseID:  1,
		}

		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("*domain.RequestUpdateEmployee")).Return(&employeeUpdated, nil)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/employees/:id", employees.Update())

		requestBody, _ := json.Marshal(RequestUpdateEmployee)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/employees/1", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Employee `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseEmployee := responseDTO.Data

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, employeeUpdated, *responseEmployee)
	})

	t.Run("update_unprocessable_entity", func(t *testing.T) {
		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/employees/:id", employees.Update())

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/employees/1", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
	t.Run("update_not_found", func(t *testing.T) {
		cardNumberID := "123"
		firstName := "teste"
		lastName := "Teste"
		warehouseId := 1

		RequestUpdateEmployee := domain.RequestUpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseId,
		}
		updatedEmployee := &domain.Employee{}

		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employeeServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("*domain.RequestUpdateEmployee")).Return(updatedEmployee, employee.ErrNotFound)
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/employees/:id", employees.Update())

		requestBody, _ := json.Marshal(RequestUpdateEmployee)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/employees/1", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("update_status_bad_request", func(t *testing.T) {
		cardNumberID := "123"
		firstName := "teste"
		lastName := "Teste"
		warehouseId := 1

		RequestUpdateEmployee := domain.RequestUpdateEmployee{
			CardNumberID: &cardNumberID,
			FirstName:    &firstName,
			LastName:     &lastName,
			WarehouseID:  &warehouseId,
		}

		employeeServiceMock := mocks.NewEmployeeServiceMock()
		employees := employees.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/employees/:id", employees.Update())

		requestBody, _ := json.Marshal(RequestUpdateEmployee)
		request := bytes.NewReader(requestBody)

		req := httptest.NewRequest(http.MethodPatch, "/api/v1/employees/xx", request)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusBadRequest, res.Code)
	})
}
