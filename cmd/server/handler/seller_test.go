package handler_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sellers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller/mocks"
	"github.com/gin-gonic/gin"
	_ "github.com/stretchr/testify"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
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
		handler := handler.NewSeller(sellerServiceMock)

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

	t.Run("find_by_id_non_existent", func(t *testing.T) {
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Seller{}, seller.ErrNotFound)
		handler := handler.NewSeller(sellerServiceMock)

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

	t.Run("invalid_id", func(t *testing.T) {
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Seller{}, seller.ErrNotFound)
		handler := handler.NewSeller(sellerServiceMock)

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

	t.Run("internal_server_error", func(t *testing.T) {
		//Configurar o mock do service
		sellerServiceMock := new(mocks.SellerServiceMock)
		sellerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Seller{}, assert.AnError)
		handler := handler.NewSeller(sellerServiceMock)

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
	//"create_ok Quando a entrada de dados for bem-sucedida, um código 201 será retornado junto com o objeto inserido.""
	//"create_bad_request Quando o JSON tiver um campo incorreto, um código 400 será retornado"
	//"create_fail Se o objeto JSON não contiver os campos necessários, um código 422 será retornado"
	//"create_conflict Se o cid já existir, ele retornará um erro 409 Conflict."
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
		handler := handler.NewSeller(sellerServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/sellers", handler.Create())

		requestBody, _ := json.Marshal(createSellerRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/sellers", request)
		req.GetBody()
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
}
