package products_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/products"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product/product_mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
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
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(productFound, nil)
		handler := products.NewProduct(productServiceMock)

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

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrNotFound)
		handler := products.NewProduct(productServiceMock)

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

	t.Run("get_invalid_id", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrNotFound)
		handler := products.NewProduct(productServiceMock)

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

	t.Run("get_internal_server_error", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Get", mock.AnythingOfType("*context.Context"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, assert.AnError)
		handler := products.NewProduct(productServiceMock)

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

func TestCreate(t *testing.T) {

	t.Run("create_create_ok", func(t *testing.T) {
		// Definir resultado da consulta
		expectedProduct := &domain.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		createProductRequestDTO := products.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(expectedProduct, nil)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body2, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Product `json:"data"`
		}

		json.Unmarshal(body2, &responseDTO)

		actualProduct := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedProduct, *actualProduct)
	})

	t.Run("create_fail", func(t *testing.T) {
		createProductRequestDTO := products.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Create", mock.AnythingOfType("*context.Context")).Return(createProductRequestDTO, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_description_empty", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("", 2, 2, 2.2, 2.2, 2.2, "2222", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_expirationRate_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 0, 2, 2.2, 2.2, 2.2, "22222", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_freezinRate_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 0, 2.2, 2.2, 2.2, "22222", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_height_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 0, 2.2, 2.2, "22222", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_length_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 0, 2.2, "22222", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_netweight_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 2.2, 0, "22222", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_productCode_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 2.2, 2.2, "", 2.2, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_recomFreezTemp_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 2.2, 2.2, "teste", 0, 2.2, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_width_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 2.2, 2.2, "teste", 2.2, 0, 2, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_productTypeID_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 2.2, 2.2, "teste", 2.2, 2.2, 0, 2)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_sellerId_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := buildProductRequestDTO("teste", 2, 2, 2.2, 2.2, 2.2, "teste", 2.2, 2.2, 2, 0)

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_conflict", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := products.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, product.ErrConflict)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusConflict, res.Code)

	})

	t.Run("create_internal_server_error", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRequestDTO := products.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int"),
			mock.AnythingOfType("int")).Return(&domain.Product{}, errors.New("error"))
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/products", handler.Create())

		requestBody, _ := json.Marshal(createProductRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/products", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)

	})

}
func TestGetAll(t *testing.T) {

	t.Run("getAll_find_all", func(t *testing.T) {
		productsFounds := &[]domain.Product{
			{
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
			},
			{
				Description:    "Teste",
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
			},
		}
		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(productsFounds, nil)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseProducts := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *productsFounds, responseProducts)
	})

	t.Run("getAll_empty_database", func(t *testing.T) {
		productsFounds := &[]domain.Product{}
		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(productsFounds, nil)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("getAll_internal_error", func(t *testing.T) {
		productsFounds := &[]domain.Product{}
		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(productsFounds, assert.AnError)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/products", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestDelete(t *testing.T) {

	t.Run("delete_non_existent", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(product.ErrNotFound)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/products/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("delete_delete_ok", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/products/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("delete_error_parsing_id", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/products/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/xyz", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("delete_error_deleting_product", func(t *testing.T) {

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(assert.AnError)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/products/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/products/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update_non_existent", func(t *testing.T) {
		description := "teste2"
		expirationRate := 2
		freezingRate := 2
		var height float32 = 2.2
		var length float32 = 2.2
		var netweight float32 = 2.2
		productCode := "teste2"
		var recomFreezTemp float32 = 2.2
		var width float32 = 2.2
		productTypeID := 2
		sellerID := 2

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateProductRequest := products.RequestUpdateProduct{
			Description:    &description,
			ExpirationRate: &expirationRate,
			FreezingRate:   &freezingRate,
			Height:         &height,
			Length:         &length,
			Netweight:      &netweight,
			ProductCode:    &productCode,
			RecomFreezTemp: &recomFreezTemp,
			Width:          &width,
			ProductTypeID:  &productTypeID,
			SellerID:       &sellerID,
		}

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.Product{}, product.ErrNotFound,
		)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/products/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/products/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("update_conflict", func(t *testing.T) {

		description := "teste2"
		expirationRate := 2
		freezingRate := 2
		var height float32 = 2.2
		var length float32 = 2.2
		var netweight float32 = 2.2
		productCode := "teste2"
		var recomFreezTemp float32 = 2.2
		var width float32 = 2.2
		productTypeID := 2
		sellerID := 2

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateProductRequest := products.RequestUpdateProduct{
			Description:    &description,
			ExpirationRate: &expirationRate,
			FreezingRate:   &freezingRate,
			Height:         &height,
			Length:         &length,
			Netweight:      &netweight,
			ProductCode:    &productCode,
			RecomFreezTemp: &recomFreezTemp,
			Width:          &width,
			ProductTypeID:  &productTypeID,
			SellerID:       &sellerID,
		}

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.Product{}, product.ErrConflict,
		)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/products/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/products/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusConflict, res.Code)
	})

	t.Run("update_internal_server_error", func(t *testing.T) {

		description := "teste2"
		expirationRate := 2
		freezingRate := 2
		var height float32 = 2.2
		var length float32 = 2.2
		var netweight float32 = 2.2
		productCode := "teste2"
		var recomFreezTemp float32 = 2.2
		var width float32 = 2.2
		productTypeID := 2
		sellerID := 2

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateProductRequest := products.RequestUpdateProduct{
			Description:    &description,
			ExpirationRate: &expirationRate,
			FreezingRate:   &freezingRate,
			Height:         &height,
			Length:         &length,
			Netweight:      &netweight,
			ProductCode:    &productCode,
			RecomFreezTemp: &recomFreezTemp,
			Width:          &width,
			ProductTypeID:  &productTypeID,
			SellerID:       &sellerID,
		}

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.Product{}, errors.New("error"),
		)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/products/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/products/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})

	t.Run("update_id_conversion_error", func(t *testing.T) {

		description := "teste2"
		expirationRate := 2
		freezingRate := 2
		var height float32 = 2.2
		var length float32 = 2.2
		var netweight float32 = 2.2
		productCode := "teste2"
		var recomFreezTemp float32 = 2.2
		var width float32 = 2.2
		productTypeID := 2
		sellerID := 2

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateProductRequest := products.RequestUpdateProduct{
			Description:    &description,
			ExpirationRate: &expirationRate,
			FreezingRate:   &freezingRate,
			Height:         &height,
			Length:         &length,
			Netweight:      &netweight,
			ProductCode:    &productCode,
			RecomFreezTemp: &recomFreezTemp,
			Width:          &width,
			ProductTypeID:  &productTypeID,
			SellerID:       &sellerID,
		}

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.Product{}, errors.New("error"),
		)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/products/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/products/a", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("update_update_ok", func(t *testing.T) {

		//productID := 1

		description := "teste2"
		expirationRate := 2
		freezingRate := 2
		var height float32 = 2.2
		var length float32 = 2.2
		var netweight float32 = 2.2
		productCode := "teste2"
		var recomFreezTemp float32 = 2.2
		var width float32 = 2.2
		productTypeID := 2
		sellerID := 2

		//(Poderia utilizar dessa maneira também) -> experirationRate := func (i int) int{return i } (2)
		updateProductRequest := products.RequestUpdateProduct{
			Description:    &description,
			ExpirationRate: &expirationRate,
			FreezingRate:   &freezingRate,
			Height:         &height,
			Length:         &length,
			Netweight:      &netweight,
			ProductCode:    &productCode,
			RecomFreezTemp: &recomFreezTemp,
			Width:          &width,
			ProductTypeID:  &productTypeID,
			SellerID:       &sellerID,
		}
		updatedProduct := &domain.Product{
			ID:             1,
			Description:    "Teste2",
			ExpirationRate: 2,
			FreezingRate:   2,
			Height:         2.2,
			Length:         2.2,
			Netweight:      2.2,
			ProductCode:    "Teste2",
			RecomFreezTemp: 2.2,
			Width:          2.2,
			ProductTypeID:  2,
			SellerID:       2,
		}

		//Configurar o mock do service
		productServiceMock := new(product_mocks.ProductServiceMock)
		productServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			updatedProduct, nil,
		)
		handler := products.NewProduct(productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/products/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/products/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.Product `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseProduct := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *updatedProduct, *responseProduct)

	})

	t.Run("update_status_unprocessable_entity", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetProducts(t)
		mockService.On("Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(&domain.Product{}, errors.New("error"))
		server.PATCH("/api/v1/products/:id", handler.Update())

		//Definir request e response
		request := httptest.NewRequest(http.MethodPatch, "/api/v1/products/2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		//Parsear response
		bodyResponse, _ := io.ReadAll(response.Body)

		var responseProduct struct {
			Data *domain.Product `json:"data"`
		}
		json.Unmarshal(bodyResponse, &responseProduct)

		assert.Equal(t, http.StatusUnprocessableEntity, response.Code)
	})

}

func buildProductRequestDTO(description string, expirationRate int, freezinRate int, height float32, length float32, netweight float32, productCode string,
	recomFreezTemp float32, width float32, productTypeID int, sellerId int) products.RequestCreateProduct {
	return products.RequestCreateProduct{
		Description:    description,
		ExpirationRate: expirationRate,
		FreezingRate:   freezinRate,
		Height:         height,
		Length:         length,
		Netweight:      netweight,
		ProductCode:    productCode,
		RecomFreezTemp: recomFreezTemp,
		Width:          width,
		ProductTypeID:  productTypeID,
		SellerID:       sellerId,
	}
}

func InitServerWithGetProducts(t *testing.T) (*gin.Engine, *product_mocks.ProductServiceMock, *products.Product) {
	t.Helper()
	server := createServer()
	mockService := new(product_mocks.ProductServiceMock)
	handler := products.NewProduct(mockService)
	return server, mockService, handler
}
func createServer() *gin.Engine {
	//Configurar o servidor
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	return server
}
