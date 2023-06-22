package handler_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		productFound := &domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Teste",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		//Configurar o mock do service
		productServiceMock := new(mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(productFound, nil)
		handler := handler.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseProduct := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *productFound, responseProduct)
	})

	t.Run("find_by_id_non_existent", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrNotFound)
		handler := handler.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("invalid_id", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrNotFound)
		handler := handler.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/xyz", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("internal_server_error", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Product{}, assert.AnError)
		handler := handler.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)

	})

}
