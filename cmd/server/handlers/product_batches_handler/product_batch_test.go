package productbatcheshandler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	productbatcheshandler "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/product_batches_handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/productbatchesdto"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productbatches/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	expectedProductBatch = domain.ProductBatches{
		ID:                 1,
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}
	payload = productbatchesdto.CreateProductBatchesDTO{
		BatchNumber:        123,
		CurrentQuantity:    10,
		CurrentTemperature: 25.5,
		DueDate:            "2023-07-10",
		InitialQuantity:    100,
		ManufacturingDate:  "2023-07-01",
		ManufacturingHour:  8,
		MinimumTemperature: 20.0,
		ProductID:          456,
		SectionID:          789,
	}
)

func TestCreate(t *testing.T) {
	t.Run("CREATE - StatusUnprocessableEntity", func(t *testing.T) {
		r, mockService, handler := InitServerWithGetSections(t)

		mockService.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.ProductBatches")).Return(&domain.ProductBatches{}, errors.New("error"))
		r.POST("/api/v1/product-batch", handler.Create())

		req := httptest.NewRequest(http.MethodPost, "/api/v1/product-batch", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
	})
	t.Run("CREATE - OK -  When data entry is successful, a 201 code will be returned along with the inserted object", func(t *testing.T) {
		r, mockService, handler := InitServerWithGetSections(t)
		mockService.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.ProductBatches")).Return(&expectedProductBatch, nil)

		r.POST("/api/v1/product-batch", handler.Create())

		requestBody, _ := json.Marshal(payload)
		req := bytes.NewReader(requestBody)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/product-batch", req)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		bodyResponse, _ := io.ReadAll(response.Body)

		var responseProductBatch struct {
			Data *domain.ProductBatches `json:"data"`
		}
		json.Unmarshal(bodyResponse, &responseProductBatch)
		actual := responseProductBatch.Data
		//Validar resultado
		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Equal(t, expectedProductBatch, *actual)
	})
	t.Run("CREATE - Create_Conflict", func(t *testing.T) {
		r, mockService, handler := InitServerWithGetSections(t)
		mockService.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.ProductBatches")).Return(&domain.ProductBatches{}, productbatches.ErrConflict)
		r.POST("/api/v1/product-batch", handler.Create())

		requestBody, _ := json.Marshal(payload)
		req := bytes.NewReader(requestBody)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/product-batch", req)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("CREATE - Create_Internal_Server_Error -  return status code 500", func(t *testing.T) {
		r, mockService, handler := InitServerWithGetSections(t)
		mockService.On("Save", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.ProductBatches")).Return(&domain.ProductBatches{}, errors.New("error"))
		r.POST("/api/v1/product-batch", handler.Create())

		requestBody, _ := json.Marshal(payload)
		req := bytes.NewReader(requestBody)

		request := httptest.NewRequest(http.MethodPost, "/api/v1/product-batch", req)
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

}

func TestGet(t *testing.T) {
	t.Run("GET - SectionProductsReports - StatusInternalServerError", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("SectionProductsReports").Return([]domain.ProductBySection{}, assert.AnError)
		server.GET("/api/v1/product-batches/sections/report-products", handler.Get())

		request := httptest.NewRequest(http.MethodGet, "/api/v1/product-batches/sections/report-products", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	t.Run("GET - SectionProductsReports - Status OK", func(t *testing.T) {
		expectedReportProducts := []domain.ProductBySection{
			{
				SectionID:     1,
				SectionNumber: "10",
				ProductsCount: 100,
			},
			{
				SectionID:     2,
				SectionNumber: "20",
				ProductsCount: 200,
			},
		}
		server, mockService, handler := InitServerWithGetSections(t)
		mockService.On("SectionProductsReports", mock.Anything).Return(expectedReportProducts, nil)

		server.GET("/api/v1/product-batches/sections/report-products", handler.Get())

		request := httptest.NewRequest(http.MethodGet, "/api/v1/product-batches/sections/report-products", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)

	})
	t.Run("GET - StatusBadRequest - Invalid ID", func(t *testing.T) {
		server, _, handler := InitServerWithGetSections(t)
		server.GET("/api/v1/product-batches/sections/report-products/:id", handler.Get())

		request := httptest.NewRequest(http.MethodGet, "/api/v1/product-batches/sections/report-products/$", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)

	})
	t.Run("GET - Not Found", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)
		server.GET("/api/v1/product-batches/sections/report-products/:id", handler.Get())
		mockService.On("SectionProductsReportsBySection", mock.AnythingOfType("int")).Return([]domain.ProductBySection{}, productbatches.ErrNotFoundSection)

		request := httptest.NewRequest(http.MethodGet, "/api/v1/product-batches/sections/report-products/10", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusNotFound, response.Code)

	})
	t.Run("READ - Internal Server Error SectionProductsReportsBySection", func(t *testing.T) {
		server, mockService, handler := InitServerWithGetSections(t)

		mockService.On("SectionProductsReportsBySection", mock.AnythingOfType("int")).Return([]domain.ProductBySection{}, assert.AnError)
		server.GET("/api/v1/product-batches/sections/report-products/:id", handler.Get())

		request := httptest.NewRequest(http.MethodGet, "/api/v1/product-batches/sections/report-products/2", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	t.Run("GET - SectionProductsReportsBySection - Status OK", func(t *testing.T) {
		expectedReportProductsBySection := []domain.ProductBySection{
			{
				SectionID:     1,
				SectionNumber: "10",
				ProductsCount: 100,
			},
			}
	
		server, mockService, handler := InitServerWithGetSections(t)
		mockService.On("SectionProductsReportsBySection", mock.AnythingOfType("int")).Return(expectedReportProductsBySection, nil)

		server.GET("/api/v1/product-batches/sections/report-products/:id", handler.Get())

		request := httptest.NewRequest(http.MethodGet, "/api/v1/product-batches/sections/report-products/1", nil)
		response := httptest.NewRecorder()

		body, _ := io.ReadAll(response.Body)
		responseResult := &productbatchesdto.ProductBatchResponse{}
		json.Unmarshal(body, responseResult)

		server.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func InitServerWithGetSections(t *testing.T) (*gin.Engine, *mocks.ProductBatchServiceMock, *productbatcheshandler.ProductBatches) {
	t.Helper()
	server := createServer()
	mockService := new(mocks.ProductBatchServiceMock)
	handler := productbatcheshandler.NewProductBatches(mockService)
	return server, mockService, handler
}
func createServer() *gin.Engine {
	//Configurar o servidor
	gin.SetMode(gin.TestMode)
	server := gin.Default()
	return server
}
