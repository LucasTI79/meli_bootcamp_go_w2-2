package sellers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/sellers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sellers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller/mocks"
	"github.com/gin-gonic/gin"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		sellerFound := &domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(sellerFound, nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Seller `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseSeller := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *sellerFound, responseSeller)

	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Seller{}, seller.ErrNotFound)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)

	})

	t.Run("get_invalid_id", func(t *testing.T) {
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Seller{}, seller.ErrNotFound)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers/xyz", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("get_internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Seller{}, assert.AnError)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)

	})

}

func TestCreate(t *testing.T) {

	//"create_conflict Se o cid já existir, ele retornará um erro 409 Conflict."

	//"create_ok Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.""
	t.Run("Create_Ok", func(t *testing.T) {
		// Definir resultado da consulta
		expectedSeller := &domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}
		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Seller")).Return(expectedSeller, nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body2, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Seller `json:"data"`
		}

		json.Unmarshal(body2, &responseDTO)

		actualSeller := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedSeller, actualSeller)
	})

	//"create_bad_request Quando o JSON tiver um campo incorreto, um código 400 será retornado"
	t.Run("create_bad_request", func(t *testing.T) {
		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         0,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		// sellerServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Seller")).Return(expectedSeller, nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		type expectedMensageResponseDTO struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		expectedMensageResponse := expectedMensageResponseDTO{
			Code:    "bad_request",
			Message: "Field CID is required: ",
		}

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var actualMessageResponse expectedMensageResponseDTO
		json.Unmarshal(body, &actualMessageResponse)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, expectedMensageResponse, actualMessageResponse) //  "Field CID is required: %s", ""
	})
	// TODO: Finalizar cobertura para os outros campos da struct relacionada a struct.

	//"create_fail Se o objeto JSON não contiver os campos necessários, um código 422 será retornado"
	t.Run("create_error_conflict", func(t *testing.T) {

		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Seller")).Return(&domain.Seller{}, seller.ErrConflict)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body2, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Seller `json:"data"`
		}

		json.Unmarshal(body2, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusConflict, res.Code)
	})

	t.Run("create_internal_server_error", func(t *testing.T) {

		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Seller")).Return(&domain.Seller{}, errors.New("error"))
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body2, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Seller `json:"data"`
		}

		json.Unmarshal(body2, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("create_fail_companyName_nil", func(t *testing.T) {

		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         1,
			CompanyName: "",
			Address:     "Test",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("create_fail_CID_nil", func(t *testing.T) {

		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         0,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("create_fail_adress_nil", func(t *testing.T) {

		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         1,
			CompanyName: "Test",
			Address:     "",
			Telephone:   "Test",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})

	t.Run("create_fail_telephone_nil", func(t *testing.T) {

		createSellerRequestDTO := dtos.CreateSellerRequestDTO{
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "",
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)

	})
}

func TestGetAll(t *testing.T) {

	/*find_all Quando a solicitação for bem-sucedida, o back-end retornará uma lista de todos os vendedores existentes - 200*/
	t.Run("GetAll_find_all", func(t *testing.T) {
		sellersFounds := &[]domain.Seller{
			{
				ID:          1,
				CID:         1,
				CompanyName: "Test",
				Address:     "Test",
				Telephone:   "Test",
			},
			{
				ID:          1,
				CID:         1,
				CompanyName: "Test",
				Address:     "Test",
				Telephone:   "Test",
			},
		}
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(sellersFounds, nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Seller `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseSellers := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *sellersFounds, responseSellers)
	})

	t.Run("GetAll_empty_database", func(t *testing.T) {
		sellersFounds := &[]domain.Seller{}
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(sellersFounds, nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("GetAll_internal_error", func(t *testing.T) {
		sellersFounds := &[]domain.Seller{}
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(sellersFounds, assert.AnError)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/sellers", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/sellers", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_delete_ok", func(t *testing.T) {

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/sellers/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/sellers/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(seller.ErrNotFound)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/sellers/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/sellers/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("delete_error_parsing_id", func(t *testing.T) {

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/sellers/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/sellers/xyz", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update_update_ok", func(t *testing.T) {
		companyName := "Test"
		address := "Test"
		telephone := "Test"

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateSellerRequestDTO := dtos.UpdateSellerRequestDTO{
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"),
			mock.AnythingOfType("*dtos.UpdateSellerRequestDTO")).Return(&domain.Seller{}, nil)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/sellers/:id", handler.Update())

		requestBody, _ := json.Marshal(updateSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/sellers/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		companyName := "Test"
		address := "Test"
		telephone := "Test"

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateSellerRequestDTO := dtos.UpdateSellerRequestDTO{
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("*dtos.UpdateSellerRequestDTO")).Return(&domain.Seller{}, seller.ErrNotFound)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/sellers/:id", handler.Update())

		requestBody, _ := json.Marshal(updateSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/sellers/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("id_conversion_error", func(t *testing.T) {

		companyName := "Test"
		address := "Test"
		telephone := "Test"

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateSellerRequestDTO := dtos.UpdateSellerRequestDTO{
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Seller")).Return(&domain.Seller{}, errors.New("error"))
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/sellers/:id", handler.Update())

		requestBody, _ := json.Marshal(updateSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/sellers/a", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Seller `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("update_conflict", func(t *testing.T) {
		companyName := "Test"
		address := "Test"
		telephone := "Test"

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateSellerRequestDTO := dtos.UpdateSellerRequestDTO{
			CompanyName: &companyName,
			Address:     &address,
			Telephone:   &telephone,
		}

		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"),
			mock.AnythingOfType("*dtos.UpdateSellerRequestDTO")).Return(&domain.Seller{}, seller.ErrConflict)
		handler := sellers.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/sellers/:id", handler.Update())

		requestBody, _ := json.Marshal(updateSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/sellers/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusConflict, res.Code)
	})

}
