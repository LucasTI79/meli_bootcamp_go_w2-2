package sellers_test

import (
	"encoding/json"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/sellers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

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

	t.Run("find_by_id_non_existent", func(t *testing.T) {
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

	t.Run("invalid_id", func(t *testing.T) {
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

	t.Run("internal_server_error", func(t *testing.T) {
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
