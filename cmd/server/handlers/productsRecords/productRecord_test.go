package productsRecords_test

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
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/productsRecords"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	mocks2 "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		// Definir resultado da consulta
		productRecordFound := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(productRecordFound, nil)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)

		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseProductRecord := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *productRecordFound, responseProductRecord)
	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrNotFound)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("get_invalid_id", func(t *testing.T) {

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrNotFound)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords/xyz", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("get_internal_server_error", func(t *testing.T) {

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Get", mock.AnythingOfType("*context.Context"),
			mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, assert.AnError)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords/:id", handler.Get())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords/1", nil)
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
		expectedProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}
		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"), mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(expectedProductRecord, nil)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body2, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body2, &responseDTO)

		actualProductRecord := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, *expectedProductRecord, *actualProductRecord)
	})

	t.Run("create_fail", func(t *testing.T) {
		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Create", mock.AnythingOfType("*context.Context")).Return(createProductRecordRequestDTO, productRecord.ErrConflict)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_last_update_rate", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRecordRequestDTO := buildProductRecordRequestDTO("", 2.2, 2.2, 2)

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrConflict)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_purchase_price_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRecordRequestDTO := buildProductRecordRequestDTO("teste", 0, 2.2, 2)

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrConflict)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_fail_sale_price_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRecordRequestDTO := buildProductRecordRequestDTO("teste", 2.2, 0, 2)

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrConflict)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_product_record_id_nil", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRecordRequestDTO := buildProductRecordRequestDTO("teste", 2.2, 2.2, 0)

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrConflict)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})

	t.Run("create_conflict", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, productRecord.ErrConflict)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusConflict, res.Code)

	})

	t.Run("create_internal_server_error", func(t *testing.T) {
		// Definir resultado da consulta
		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1.1,
			SalePrice:      1.1,
			ProductId:      1,
		}

		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("string"), mock.AnythingOfType("float32"),
			mock.AnythingOfType("float32"), mock.AnythingOfType("int")).Return(&domain.ProductRecord{}, errors.New("error"))
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.POST("/api/v1/productsRecords", handler.Create())

		requestBody, _ := json.Marshal(createProductRecordRequestDTO)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPost, "/api/v1/productsRecords", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)
		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)

	})
}

func TestGetAll(t *testing.T) {

	t.Run("getAll_find_all", func(t *testing.T) {
		productsRecordsFounds := &[]domain.ProductRecord{
			{
				ID:             1,
				LastUpdateRate: "Test",
				PurchasePrice:  1.1,
				SalePrice:      1.1,
				ProductId:      1,
			},
			{
				LastUpdateRate: "Test",
				PurchasePrice:  1.1,
				SalePrice:      1.1,
				ProductId:      1,
			},
		}
		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(productsRecordsFounds, nil)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data []domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)
		responseProductsRecords := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *productsRecordsFounds, responseProductsRecords)
	})

	t.Run("getAll_empty_database", func(t *testing.T) {
		productsRecordsFounds := &[]domain.ProductRecord{}
		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(productsRecordsFounds, nil)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("getAll_internal_error", func(t *testing.T) {
		productsRecordsFounds := &[]domain.ProductRecord{}
		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(productsRecordsFounds, assert.AnError)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.GET("/api/v1/productsRecords", handler.GetAll())

		//Definir request e response
		req := httptest.NewRequest(http.MethodGet, "/api/v1/productsRecords", nil)
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
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(productRecord.ErrNotFound)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/productsRecords/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/productsRecords/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("delete_delete_ok", func(t *testing.T) {

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/productsRecords/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/productsRecords/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusNoContent, res.Code)
	})

	t.Run("delete_error_parsing_id", func(t *testing.T) {

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(nil)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/productsRecords/:id", handler.Delete())

		//Definir request e response
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/productsRecords/xyz", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("delete_error_deleting_product_record", func(t *testing.T) {

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(assert.AnError)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.DELETE("/api/v1/productsRecords/:id", handler.Delete())

		//Definir request e response'
		req := httptest.NewRequest(http.MethodDelete, "/api/v1/productsRecords/1", nil)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update_non_existent", func(t *testing.T) {
		lastUpdaterate := "teste2"
		var purchasePrice float32 = 2.2
		var salePrice float32 = 2.2
		productID := 2

		updateProductRecordRequest := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &lastUpdaterate,
			PurchasePrice:  &purchasePrice,
			SalePrice:      &salePrice,
			ProductId:      &productID,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.ProductRecord{}, productRecord.ErrNotFound,
		)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/productsRecords/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRecordRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/productsRecords/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusNotFound, res.Code)
	})

	t.Run("update_conflict", func(t *testing.T) {

		lastUpdaterate := "teste2"
		var purchasePrice float32 = 2.2
		var salePrice float32 = 2.2
		productID := 2

		updateProductRecordRequest := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &lastUpdaterate,
			PurchasePrice:  &purchasePrice,
			SalePrice:      &salePrice,
			ProductId:      &productID,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.ProductRecord{}, productRecord.ErrConflict,
		)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/productsRecords/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRecordRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/productsRecords/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusConflict, res.Code)
	})

	t.Run("update_internal_server_error", func(t *testing.T) {

		lastUpdaterate := "teste2"
		var purchasePrice float32 = 2.2
		var salePrice float32 = 2.2
		productID := 2

		updateProductRecordRequest := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &lastUpdaterate,
			PurchasePrice:  &purchasePrice,
			SalePrice:      &salePrice,
			ProductId:      &productID,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.ProductRecord{}, errors.New("error"),
		)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/productsRecords/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRecordRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/productsRecords/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusInternalServerError, res.Code)
	})
	t.Run("update_id_conversion_error", func(t *testing.T) {

		lastUpdaterate := "teste2"
		var purchasePrice float32 = 2.2
		var salePrice float32 = 2.2
		productID := 2

		updateProductRecordRequest := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &lastUpdaterate,
			PurchasePrice:  &purchasePrice,
			SalePrice:      &salePrice,
			ProductId:      &productID,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			&domain.ProductRecord{}, errors.New("error"),
		)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/productsRecords/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRecordRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/productsRecords/a", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		//Validar resultado
		assert.Equal(t, http.StatusBadRequest, res.Code)
	})

	t.Run("update_update_ok", func(t *testing.T) {

		//productID := 1

		lastUpdaterate := "teste2"
		var purchasePrice float32 = 2.2
		var salePrice float32 = 2.2
		productID := 2

		updateProductRecordRequest := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &lastUpdaterate,
			PurchasePrice:  &purchasePrice,
			SalePrice:      &salePrice,
			ProductId:      &productID,
		}
		updatedProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Teste2",
			PurchasePrice:  2,
			SalePrice:      2,
			ProductId:      2,
		}

		//Configurar o mock do service
		productRecordServiceMock := new(mocks.ProductRecordServiceMock)
		productRecordServiceMock.On(
			"Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(
			updatedProductRecord, nil,
		)
		productServiceMock := new(mocks2.ProductServiceMock)
		handler := productsRecords.NewProductRecord(productRecordServiceMock, productServiceMock)
		//Configurar o servidor
		gin.SetMode(gin.TestMode)
		r := gin.Default()
		r.PATCH("/api/v1/productsRecords/:id", handler.Update())

		requestBody, _ := json.Marshal(updateProductRecordRequest)
		request := bytes.NewReader(requestBody)

		//Definir request e response
		req := httptest.NewRequest(http.MethodPatch, "/api/v1/productsRecords/1", request)
		res := httptest.NewRecorder()

		//Executar request
		r.ServeHTTP(res, req)

		//Parsear response
		body, _ := ioutil.ReadAll(res.Body)

		var responseDTO struct {
			Data *domain.ProductRecord `json:"data"`
		}

		json.Unmarshal(body, &responseDTO)

		responseProductRecord := responseDTO.Data

		//Validar resultado
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, *updatedProductRecord, *responseProductRecord)

	})

	t.Run("update_status_unprocessable_entity", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetProductsRecords(t)
		mockService.On("Update",
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(&domain.ProductRecord{}, errors.New("error"))
		server.PATCH("/api/v1/productsRecords/:id", handler.Update())

		//Definir request e response
		request := httptest.NewRequest(http.MethodPatch, "/api/v1/productsRecords/2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		//Parsear response
		bodyResponse, _ := io.ReadAll(response.Body)

		var responseProductRecord struct {
			Data *domain.ProductRecord `json:"data"`
		}
		json.Unmarshal(bodyResponse, &responseProductRecord)

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

func buildProductRecordRequestDTO(lastUpdaterate string, purchasePrice float32, salePrice float32, productID int) productsRecords.RequestCreateProductRecord {
	return productsRecords.RequestCreateProductRecord{
		LastUpdateRate: lastUpdaterate,
		PurchasePrice:  purchasePrice,
		SalePrice:      salePrice,
		ProductId:      productID,
	}
}

func InitServerWithGetProductsRecords(t *testing.T) (*gin.Engine, *mocks.ProductRecordServiceMock, *productsRecords.ProductRecord) {
	t.Helper()
	server := createServer()
	mockProductRecordService := new(mocks.ProductRecordServiceMock)
	mockProductService := new(mocks2.ProductServiceMock)
	handler := productsRecords.NewProductRecord(mockProductRecordService, mockProductService)
	return server, mockProductRecordService, handler
}
func createServer() *gin.Engine {
	//Configurar o servidor
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	return server
}
