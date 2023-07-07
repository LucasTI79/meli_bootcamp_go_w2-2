package handlers_test

//
//import (
//	"bytes"
//	"encoding/json"
//	"errors"
//	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services/mocks"
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
//	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/handlers"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//)
//
//func Test_sectionHandler_GetAll(t *testing.T) {
//	t.Run("READ - find_all: Should return status 200 with all sections", func(t *testing.T) {
//		// Definir resultado da consulta
//		expectedSections := &[]entities.Section{
//			{
//				ID:                 1,
//				SectionNumber:      65473,
//				CurrentTemperature: 15,
//				MinimumTemperature: 5,
//				CurrentCapacity:    10,
//				MinimumCapacity:    12,
//				MaximumCapacity:    20,
//				WarehouseID:        1234,
//				ProductTypeID:      874893,
//			},
//			{
//				ID:                 2,
//				SectionNumber:      4653,
//				CurrentTemperature: 20,
//				MinimumTemperature: 50,
//				CurrentCapacity:    1000,
//				MinimumCapacity:    120,
//				MaximumCapacity:    200,
//				WarehouseID:        9878,
//				ProductTypeID:      87489223,
//			},
//		}
//
//		//Configurar o mock do service
//		sectionServiceMock := mocks.NewSectionServiceMock(t)
//		sectionServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedSections, nil)
//		handler := handlers.NewSection(sectionServiceMock)
//
//		//Configurar o servidor
//		server := createServer()
//		server.GET("/api/v1/sections", handler.GetAll())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
//		response := httptest.NewRecorder()
//
//		//Executar o servidor
//		server.ServeHTTP(response, request)
//
//		//Parsear response
//		body, _ := io.ReadAll(response.Body)
//		var responseDTO struct {
//			Data *[]entities.Section `json:"data"`
//		}
//		json.Unmarshal(body, &responseDTO)
//		actualSections := responseDTO.Data
//
//		//Validar resultado
//		assert.Equal(t, http.StatusOK, response.Code)
//		assert.Equal(t, *expectedSections, *actualSections)
//	})
//
//	t.Run("READ - Empty Database - should return status 204 when the list returns empty", func(t *testing.T) {
//		// Definir resultado da consulta
//		expectedSections := &[]entities.Section{}
//
//		//Configurar o mock do service
//		sectionServiceMock := mocks.NewSectionServiceMock(t)
//		sectionServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(expectedSections, nil)
//		handler := handlers.NewSection(sectionServiceMock)
//
//		//Configurar o servidor
//		server := createServer()
//		server.GET("/api/v1/sections", handler.GetAll())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
//		response := httptest.NewRecorder()
//
//		//Executar o servidor
//		server.ServeHTTP(response, request)
//
//		//Validar resultado
//		assert.Equal(t, http.StatusNoContent, response.Code)
//	})
//
//	t.Run("READ - ServerInternalError - Should return error 500 when http side error occurs with the database.", func(t *testing.T) {
//		//Definir resultado da consulta
//		sectionServiceMock := mocks.NewSectionServiceMock(t)
//		sectionServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(&[]entities.Section{}, assert.AnError)
//		handler := handlers.NewSection(sectionServiceMock)
//
//		//Configurar o servidor
//		gin.SetMode(gin.TestMode)
//		r := gin.Default()
//		r.GET("/api/v1/sections", handler.GetAll())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections", nil)
//		response := httptest.NewRecorder()
//
//		//Executar request
//		r.ServeHTTP(response, request)
//
//		//Validar resultado
//		assert.Equal(t, http.StatusInternalServerError, response.Code)
//	})
//}
//
//func Test_sectionHandler_Get(t *testing.T) {
//	t.Run("READ - find_by_id_existent - Should return status 200 when finding a section by ID", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		expectedSection := &entities.Section{
//			ID:                 2,
//			SectionNumber:      2,
//			CurrentTemperature: 2,
//			MinimumTemperature: 2,
//			CurrentCapacity:    2,
//			MinimumCapacity:    2,
//			MaximumCapacity:    2,
//			WarehouseID:        2,
//			ProductTypeID:      2,
//		}
//		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(expectedSection, nil)
//		server.GET("/api/v1/sections/:id", handler.Get())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		body, _ := io.ReadAll(response.Body)
//		responseResult := &dtos.SectionResponse{}
//		json.Unmarshal(body, responseResult)
//
//		assert.Equal(t, http.StatusOK, response.Code)
//		assert.Equal(t, *expectedSection, responseResult.Data)
//	})
//
//	t.Run("READ - find_by_id_non_existent - Should return error 404 when not finding a section by id", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&entities.Section{}, services.ErrNotFound)
//		server.GET("/api/v1/sections/:id", handler.Get())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/1", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusNotFound, response.Code)
//	})
//
//	t.Run("READ - invalid id - should return error 400 when an invalid ID is entered.", func(t *testing.T) {
//		server, _, handler := InitServerWithGetSections(t)
//
//		server.GET("/api/v1/sections/:id", handler.Get())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/a", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusBadRequest, response.Code)
//	})
//
//	t.Run("READ - should return error 500 when an internal http error occurs when looking up a section by id", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		mockService.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&entities.Section{}, assert.AnError)
//		//Dúvida sobre o retorno section.ErrEmployeeNotFound ao invés de um invalid id personalizado...
//		server.GET("/api/v1/sections/:id", handler.Get())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodGet, "/api/v1/sections/10", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusInternalServerError, response.Code)
//	})
//}
//
//func Test_sectionHandler_Delete(t *testing.T) {
//	t.Run("DELETE - OK - When the deletion is successful, a 204 code is returned.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		mockService.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
//
//		server.DELETE("/api/v1/sections/:id", handler.Delete())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusNoContent, response.Code)
//	})
//	t.Run("DELETE - Delete_non_existent - Should return status 404 when deleting a section that does not exist.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		mockService.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(services.ErrNotFound)
//
//		server.DELETE("/api/v1/sections/:id", handler.Delete())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusNotFound, response.Code)
//	})
//	t.Run("DELETE - ID invalid - Should return error 400 when trying to delete a section with invalid ID.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		mockService.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
//
//		server.DELETE("/api/v1/sections/:id", handler.Delete())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/x", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusBadRequest, response.Code)
//	})
//
//	t.Run("DELETE - Server Internal Error - Should return error 500 when an internal http error occurs while deleting a section.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//
//		mockService.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(assert.AnError)
//
//		server.DELETE("/api/v1/sections/:id", handler.Delete())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodDelete, "/api/v1/sections/1", nil)
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, request)
//
//		assert.Equal(t, http.StatusInternalServerError, response.Code)
//	})
//}
//
//func Test_sectionHandler_Create(t *testing.T) {
//	t.Run("CREATE - OK - When data entry is successful, a 201 code will be returned along with the inserted object", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(expectedSection, nil)
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		//Parsear response
//		bodyResponse, _ := io.ReadAll(response.Body)
//
//		var responseSection struct {
//			Data *entities.Section `json:"data"`
//		}
//		json.Unmarshal(bodyResponse, &responseSection)
//		actualSection := responseSection.Data
//		//Validar resultado
//		assert.Equal(t, http.StatusCreated, response.Code)
//		assert.Equal(t, *expectedSection, *actualSection)
//	})
//	t.Run("CREATE - StatusUnprocessableEntity", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", nil)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//
//	t.Run("CREATE - Create_Conflict - If the section_number already exists, it will return a 409 Conflict error.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, services.ErrConflict)
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusConflict, response.Code)
//	})
//
//	t.Run("CREATE - Create_Internal_Server_Error -  return status code 500", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusInternalServerError, response.Code)
//	})
//
//	t.Run("CREATE - Create_Fail_SectionNumber_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{
//			CurrentTemperature: 10,
//			MinimumTemperature: 10,
//			CurrentCapacity:    10,
//			MinimumCapacity:    10,
//			MaximumCapacity:    10,
//			WarehouseID:        10,
//			ProductTypeID:      10,
//		}
//
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_CurrentTemperature_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{SectionNumber: 10, CurrentTemperature: 0, MinimumTemperature: 10, CurrentCapacity: 10, MinimumCapacity: 10, MaximumCapacity: 10, WarehouseID: 10, ProductTypeID: 10}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_MinimumTemperature_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{SectionNumber: 10, CurrentTemperature: 10, CurrentCapacity: 10, MinimumCapacity: 10, MaximumCapacity: 10, WarehouseID: 10, ProductTypeID: 10}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_CurrentCapacity_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{10, 10, 10, 0, 10, 10, 10, 10}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_MinimumCapacity_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{10, 10, 10, 10, 0, 10, 10, 10}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_MaximumCapacity_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{10, 10, 10, 10, 10, 0, 10, 10}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_WarehouseID_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{10, 10, 10, 10, 10, 10, 0, 10}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("CREATE - Create_Fail_ProductTypeID_Nil - Status Code 422", func(t *testing.T) {
//		requestSection := dtos.CreateSectionRequestDTO{10, 10, 10, 10, 10, 10, 10, 0}
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Save",
//			mock.AnythingOfType("*context.Context"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//			mock.AnythingOfType("int"),
//		).Return(&entities.Section{}, errors.New("error"))
//		server.POST("/api/v1/sections", handler.Create())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		request := httptest.NewRequest(http.MethodPost, "/api/v1/sections", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//}
//
//func Test_sectionHandler_Update(t *testing.T) {
//	t.Run("UPDATE - OK - When the data update is successful, the section with the updated information is returned along with a 200 code.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Update",
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//		).Return(expectedSection, nil)
//		server.PATCH("/api/v1/sections/:id", handler.Update())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/1", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		//Parsear response
//		bodyResponse, _ := io.ReadAll(response.Body)
//
//		var responseSection struct {
//			Data *entities.Section `json:"data"`
//		}
//		json.Unmarshal(bodyResponse, &responseSection)
//		actualSection := responseSection.Data
//		//Validar resultado
//		assert.Equal(t, http.StatusOK, response.Code)
//		assert.Equal(t, *expectedSection, *actualSection)
//	})
//	t.Run("UPDATE - ID_No_Existent - If the section to be updated does not exist, a 404 code is returned.", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Update",
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//		).Return(&entities.Section{}, services.ErrNotFound)
//		server.PATCH("/api/v1/sections/:id", handler.Update())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/2", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		//Parsear response
//		bodyResponse, _ := io.ReadAll(response.Body)
//
//		var responseSection struct {
//			Data *entities.Section `json:"data"`
//		}
//		json.Unmarshal(bodyResponse, &responseSection)
//
//		assert.Equal(t, http.StatusNotFound, response.Code)
//	})
//	t.Run("UPDATE - StatusUnprocessableEntity", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Update",
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//		).Return(&entities.Section{}, errors.New("error"))
//		server.PATCH("/api/v1/sections/:id", handler.Update())
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/2", nil)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		//Parsear response
//		bodyResponse, _ := io.ReadAll(response.Body)
//
//		var responseSection struct {
//			Data *entities.Section `json:"data"`
//		}
//		json.Unmarshal(bodyResponse, &responseSection)
//
//		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
//	})
//	t.Run("UPDATE - Conflit - Should return conflict error", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Update",
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//		).Return(&entities.Section{}, services.ErrConflict)
//		server.PATCH("/api/v1/sections/:id", handler.Update())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/1", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//
//		//Parsear response
//		bodyResponse, _ := io.ReadAll(response.Body)
//
//		var responseSection struct {
//			Data *entities.Section `json:"data"`
//		}
//		json.Unmarshal(bodyResponse, &responseSection)
//		//Validar resultado
//		assert.Equal(t, http.StatusConflict, response.Code)
//	})
//	t.Run("UPDATE - Invalid ID - Should return bad request error when id is invalid", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Update",
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//		).Return(&entities.Section{}, errors.New("error"))
//		server.PATCH("/api/v1/sections/:id", handler.Update())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/x", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusBadRequest, response.Code)
//	})
//	t.Run("UPDATE - Server_Internal_Error - Should return bad request error when id is invalid", func(t *testing.T) {
//		server, mockService, handler := InitServerWithGetSections(t)
//		mockService.On("Update",
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//			mock.Anything,
//		).Return(&entities.Section{}, errors.New("error"))
//		server.PATCH("/api/v1/sections/:id", handler.Update())
//
//		requestBody, _ := json.Marshal(requestSection)
//		req := bytes.NewReader(requestBody)
//
//		//Definir request e response
//		request := httptest.NewRequest(http.MethodPatch, "/api/v1/sections/1", req)
//		response := httptest.NewRecorder()
//
//		server.ServeHTTP(response, request)
//		assert.Equal(t, http.StatusInternalServerError, response.Code)
//	})
//}
