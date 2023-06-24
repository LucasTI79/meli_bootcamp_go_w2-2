package handler_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/section/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T){
	t.Run("READ - find_all: Should return status 200 with all sections", func(t *testing.T) {
		// Definir resultado da consulta
		expectedSections := &[]domain.Section{
			{
				ID:                 1,
				SectionNumber:      65473,
				CurrentTemperature: 15,
				MinimumTemperature: 5,
				CurrentCapacity:    10,
				MinimumCapacity:    12,
				MaximumCapacity:    20,
				WarehouseID:        1234,
				ProductTypeID:      874893,
			},
			{
				ID:                 2,
				SectionNumber:      4653,
				CurrentTemperature: 20,
				MinimumTemperature: 50,
				CurrentCapacity:    1000,
				MinimumCapacity:    120,
				MaximumCapacity:    200,
				WarehouseID:        9878,
				ProductTypeID:      87489223,
			},
		}

		//Configurar o mock do service
		sectionServiceMock := new(mocks.SectionServiceMock)
		sectionServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedSections, nil)
		handler := handler.NewSection(sectionServiceMock)

		//Configurar o servidor
		server := createServer()
		server.GET("/api/v1/sections", handler.GetAll())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
		response := httptest.NewRecorder()

		//Executar o servidor
		server.ServeHTTP(response, request)

		//Parsear response
		body, _ := ioutil.ReadAll(response.Body)
		var responseDTO struct {
			Data *[]domain.Section `json:"data"`
		}
		json.Unmarshal(body, &responseDTO)
		actualSections := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, response.Code)
		fmt.Println(actualSections)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *expectedSections, *actualSections)
		assert.True(t, len(*actualSections) == 2)
	})
	// Not implemented in cod this usecase.
	t.Run("READ - Empty Database - should return status 204 when the list returns empty", func(t *testing.T) {
		// Definir resultado da consulta
		expectedSections := &[]domain.Section{}

		//Configurar o mock do service
		sectionServiceMock := new(mocks.SectionServiceMock)
		sectionServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedSections, nil)
		handler := handler.NewSection(sectionServiceMock)

		//Configurar o servidor
		server := createServer()
		server.GET("/api/v1/sections", handler.GetAll())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
		response := httptest.NewRecorder()

		//Executar o servidor
		server.ServeHTTP(response, request)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, response.Code)
	})

	t.Run("READ - ServerInternalError - Should return error 500 when server side error occurs with the database.",func(t *testing.T){
		//Definir resultado da consulta
		//expectedSection := &[]domain.Section{} @@@Verificar se esse objeto pode ser substituído pelo nil linha 110
		//Configurar o mock do service
		sectionServiceMock := new(mocks.SectionServiceMock)
		sectionServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(nil, assert.AnError)
		handler := handler.NewSection(sectionServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sections", handler.GetAll())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
		response := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(response, request)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
} 

func TestGet(t *testing.T) {
	t.Run("READ - find_by_id_existent - Should return status 200 when finding a section by ID", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		expectedSection := &domain.Section{
				ID: 2,
				SectionNumber: 2,
				CurrentTemperature: 2,
				MinimumTemperature: 2,
				CurrentCapacity: 2,
				MinimumCapacity: 2,
				MaximumCapacity: 2,
				WarehouseID: 2,
				ProductTypeID: 2,
		}
		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(expectedSection, nil)
		server.GET("/api/v1/sections/:id", handler.Get())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		body, _ := ioutil.ReadAll(response.Body)
		responseResult := &domain.SectionResponse{}
		json.Unmarshal(body, responseResult)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, *expectedSection, responseResult.Data)
	})

	t.Run("READ - find_by_id_non_existent - Should return error 404 when not finding a section by id", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Section{}, section.ErrNotFound)
		server.GET("/api/v1/sections/:id", handler.Get())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("READ - invalid id - should return error 400 when an invalid ID is entered.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Section{}, section.ErrNotFound)
		//Dúvida sobre o retorno section.ErrNotFound ao invés de um invalid id personalizado...
		server.GET("/api/v1/sections/:id", handler.Get())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/a", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("READ - invalid id - should return error 500 when an internal server error occurs when looking up a section by id", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Section{}, assert.AnError)
		//Dúvida sobre o retorno section.ErrNotFound ao invés de um invalid id personalizado...
		server.GET("/api/v1/sections/:id", handler.Get())

		//Definir request e response
		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/10", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func TestDelete(t *testing.T){
	t.Run("DELETE - OK - When the deletion is successful, a 204 code is returned.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Delete",  mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)

		server.DELETE("/api/v1/sections/:id", handler.Delete())

		//Definir request e response
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
	t.Run("DELETE - Delete_non_existent - Should return status 404 when deleting a section that does not exist.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Delete",  mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(section.ErrNotFound)

		server.DELETE("/api/v1/sections/:id", handler.Delete())

		//Definir request e response
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("DELETE - ID invalid - Should return error 400 when trying to delete a section with invalid ID.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Delete",  mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)

		server.DELETE("/api/v1/sections/:id", handler.Delete())

		//Definir request e response
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/x", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("DELETE - Server Internal Error - Should return error 500 when an internal server error occurs while deleting a section.", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Delete",  mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(assert.AnError)

		server.DELETE("/api/v1/sections/:id", handler.Delete())

		//Definir request e response
		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
}

func InitServerWithGetSections(t *testing.T) (*gin.Engine, *mocks.SectionServiceMock, *handler.Section) {
	t.Helper()
	server := createServer()
	mockService := new(mocks.SectionServiceMock)
	handler := handler.NewSection(mockService)
	return server, mockService, handler
}
func createServer() *gin.Engine {
	//Configurar o servidor
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	return server
}