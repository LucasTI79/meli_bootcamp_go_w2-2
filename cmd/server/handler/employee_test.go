package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
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
		employeeServiceMock := new(mocks.EmployeeServiceMock)
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(employeeFound, nil)
		handler := handler.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", handler.Get())

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
		employeeServiceMock := new(mocks.EmployeeServiceMock)
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Employee{}, handler.ErrNotFound)
		handler := handler.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", handler.Get())

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
		employeeServiceMock := new(mocks.EmployeeServiceMock)
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Employee{}, handler.ErrNotFound)
		handler := handler.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", handler.Get())

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
		employeeServiceMock := new(mocks.EmployeeServiceMock)
		employeeServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Employee{}, assert.AnError)
		handler := handler.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees/:id", handler.Get())

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
				FirstName:    "Joao",
				LastName:     "Silva",
				WarehouseID:  2,
			},
		}

		//Configurar o mock do service
		employeeServiceMock := new(mocks.EmployeeServiceMock)
		employeeServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedEmployee, nil)
		handler := handler.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees", handler.GetAll())

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
		employeeServiceMock := new(mocks.EmployeeServiceMock)
		employeeServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(&[]domain.Employee{}, assert.AnError)
		handler := handler.NewEmployee(employeeServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/employees", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	// t.Run("find_no_content", func(t *testing.T) {
	// 	//Configurar o mock do service
	// 	employeeServiceMock := new(mocks.EmployeeServiceMock)
	// 	employeeServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(&[]domain.Employee{}, nil)
	// 	handler := handler.NewEmployee(employeeServiceMock)

	// 	//Configurar o servidor
	// 	gin.SetMode(gin.TestMode)
	// 	r := gin.Default()
	// 	r.GET("/api/v1/employees", handler.GetAll())

	// 	//Definir request e response
	// 	req := httptest.NewRequest(http.MethodGet, "/api/v1/employees", nil)
	// 	res := httptest.NewRecorder()

	// 	//Executar request
	// 	r.ServeHTTP(res, req)

	// 	//Validar resultado
	// 	assert.Equal(t, http.StatusNoContent, res.Code)
	// })
}
