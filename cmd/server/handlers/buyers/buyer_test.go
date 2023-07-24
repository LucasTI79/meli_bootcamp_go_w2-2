package buyers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/buyers"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	servicesMocks "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/purchaseOrder/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGet(t *testing.T) {

	buyerSerialized, _ := os.ReadFile("../../../../test/resources/valid_buyer.json")
	var validBuyer domain.Buyer
	if err := json.Unmarshal(buyerSerialized, &validBuyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name             string
		id               string
		expectedGetCalls int
		expectedGetError error
		expectedBuyer    *domain.Buyer
		expectedCode     int
	}{
		{
			name:             "Success finding buyer",
			id:               "1",
			expectedGetCalls: 1,
			expectedBuyer:    &validBuyer,
			expectedGetError: nil,
			expectedCode:     http.StatusOK,
		},
		{
			name:             "Error finding buyer",
			id:               "999",
			expectedGetCalls: 1,
			expectedGetError: buyer.ErrNotFound,
			expectedBuyer:    &domain.Buyer{},
			expectedCode:     http.StatusNotFound,
		},
		{
			name:             "Error connecting db",
			id:               "1",
			expectedGetCalls: 1,
			expectedGetError: assert.AnError,
			expectedBuyer:    &domain.Buyer{},
			expectedCode:     http.StatusInternalServerError,
		},
		{
			name:         "Error invalid id",
			id:           "xyz",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedBuyer, test.expectedGetError)

			purchaseOrderServiceMock := servicesMocks.NewMockPurchaseOrderService(t)
			buyerHandler := buyers.NewBuyerHandler(buyerServiceMock, purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/buyers/:id", buyerHandler.Get())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", "/api/v1/buyers", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			buyerServiceMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			if test.expectedCode == http.StatusOK {
				//Parsear response
				var buyerResponse domain.Buyer
				body, _ := io.ReadAll(res.Body)
				json.Unmarshal(body, &buyerResponse)

				assert.Equal(t, *test.expectedBuyer, buyerResponse)
			}

		})
	}

}

func TestGetAll(t *testing.T) {

	expectedBuyers := &[]domain.Buyer{
		{
			ID:           1,
			CardNumberID: "123",
			FirstName:    "Test",
			LastName:     "Test",
		},
		{
			ID:           2,
			CardNumberID: "456",
			FirstName:    "Test2",
			LastName:     "Test2",
		},
	}

	tests := []struct {
		name                 string
		expectedGetAllResult *[]domain.Buyer
		expectedGetAllError  error
		expectedGetAllCalls  int
		expectedResponse     *[]domain.Buyer
		expectedCode         int
	}{
		{
			name:                 "Successfully get all expectedBuyers",
			expectedGetAllResult: expectedBuyers,
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedResponse:     expectedBuyers,
			expectedCode:         http.StatusOK,
		},
		{
			name:                 "Success empty database",
			expectedGetAllResult: &[]domain.Buyer{},
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedResponse:     &[]domain.Buyer{},
			expectedCode:         http.StatusNoContent,
		},
		{
			name:                 "Error getting all expectedBuyers",
			expectedGetAllResult: &[]domain.Buyer{},
			expectedGetAllError:  assert.AnError,
			expectedGetAllCalls:  1,
			expectedResponse:     &[]domain.Buyer{},
			expectedCode:         http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(test.expectedGetAllResult, test.expectedGetAllError)

			purchaseOrderServiceMock := servicesMocks.NewMockPurchaseOrderService(t)
			buyerHandler := buyers.NewBuyerHandler(buyerServiceMock, purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/expectedBuyers", buyerHandler.GetAll())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s", "/api/v1/expectedBuyers"), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			buyerServiceMock.AssertNumberOfCalls(t, "GetAll", test.expectedGetAllCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var responseDTO struct {
					Data *[]domain.Buyer `json:"data"`
				}
				json.Unmarshal(body, &responseDTO)

				buyerResponse := responseDTO.Data

				// Valida o response
				assert.Equal(t, *test.expectedResponse, *buyerResponse)
			}
		})
	}
}

func TestDelete(t *testing.T) {

	tests := []struct {
		name                string
		id                  string
		expectedDeleteError error
		expectedDeleteCalls int
		expectedCode        int
	}{
		{
			name:                "Successfully delete buyer",
			id:                  "1",
			expectedDeleteError: nil,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusNoContent,
		},
		{
			name:                "Error deleting inexistent buyer",
			id:                  "1",
			expectedDeleteError: buyer.ErrNotFound,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusNotFound,
		},
		{
			name:                "Error deleting buyer",
			id:                  "1",
			expectedDeleteError: assert.AnError,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusInternalServerError,
		},
		{
			name:                "Error invalid id",
			id:                  "xyz",
			expectedDeleteError: assert.AnError,
			expectedDeleteCalls: 0,
			expectedCode:        http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedDeleteError)

			purchaseOrderServiceMock := servicesMocks.NewMockPurchaseOrderService(t)
			buyerHandler := buyers.NewBuyerHandler(buyerServiceMock, purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.DELETE("/api/v1/buyers/:id", buyerHandler.Delete())

			//Definir request e response
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", "/api/v1/buyers", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			buyerServiceMock.AssertNumberOfCalls(t, "Delete", test.expectedDeleteCalls)
			assert.Equal(t, test.expectedCode, res.Code)

		})
	}
}

func TestCreate(t *testing.T) {

	createBuyerRequest := dtos.CreateBuyerRequestDTO{
		CardNumberID: "123",
		FirstName:    "Test",
		LastName:     "Test",
	}

	expectedBuyer := &domain.Buyer{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Test",
		LastName:     "Test",
	}

	tests := []struct {
		name                 string
		createBuyerRequest   dtos.CreateBuyerRequestDTO
		expectedCreateResult *domain.Buyer
		expectedCreateError  error
		expectedCreateCalls  int
		expectedResponse     *domain.Buyer
		expectedCode         int
	}{
		{
			name:                 "Successfully creating buyer",
			createBuyerRequest:   createBuyerRequest,
			expectedCreateResult: expectedBuyer,
			expectedCreateError:  nil,
			expectedCreateCalls:  1,
			expectedCode:         http.StatusCreated,
		},
		{
			name:                 "Error creating buyer with duplicated card number id",
			createBuyerRequest:   createBuyerRequest,
			expectedCreateResult: &domain.Buyer{},
			expectedCreateError:  buyer.ErrCardNumberDuplicated,
			expectedCreateCalls:  1,
			expectedCode:         http.StatusConflict,
		},
		{
			name:                 "Error creating buyer",
			createBuyerRequest:   createBuyerRequest,
			expectedCreateResult: &domain.Buyer{},
			expectedCreateError:  assert.AnError,
			expectedCreateCalls:  1,
			expectedCode:         http.StatusInternalServerError,
		},
		{
			name:               "Error invalid buyer",
			createBuyerRequest: dtos.CreateBuyerRequestDTO{},
			expectedCode:       http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("*dtos.CreateBuyerRequestDTO")).Return(test.expectedCreateResult, test.expectedCreateError)

			purchaseOrderServiceMock := servicesMocks.NewMockPurchaseOrderService(t)
			buyerHandler := buyers.NewBuyerHandler(buyerServiceMock, purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.POST("/api/v1/buyers", buyerHandler.Create())

			requestBody, _ := json.Marshal(test.createBuyerRequest)
			request := bytes.NewReader(requestBody)

			//Definir request e response
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s", "/api/v1/buyers"), request)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			buyerServiceMock.AssertNumberOfCalls(t, "Create", test.expectedCreateCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var buyerResponse *domain.Buyer
				json.Unmarshal(body, &buyerResponse)

				// Valida o response
				assert.Equal(t, *test.expectedResponse, *buyerResponse)
			}

		})
	}
}

func TestUpdate(t *testing.T) {

	updateBuyerRequest := dtos.UpdateBuyerRequestDTO{
		CardNumberID: func(s string) *string { return &s }("456"),
	}

	buyerUpdated := &domain.Buyer{
		ID:           1,
		CardNumberID: "456",
		FirstName:    "Test",
		LastName:     "Test",
	}

	tests := []struct {
		name                 string
		id                   string
		updateBuyerRequest   dtos.UpdateBuyerRequestDTO
		expectedUpdateResult *domain.Buyer
		expectedUpdateError  error
		expectedUpdateCalls  int
		expectedResponse     *domain.Buyer
		expectedCode         int
	}{
		{
			name:                 "Successfully updating buyer",
			id:                   "1",
			updateBuyerRequest:   updateBuyerRequest,
			expectedUpdateResult: buyerUpdated,
			expectedUpdateError:  nil,
			expectedUpdateCalls:  1,
			expectedResponse:     buyerUpdated,
			expectedCode:         http.StatusOK,
		},
		{
			name:                 "Error updating inexisting buyer",
			id:                   "1",
			updateBuyerRequest:   updateBuyerRequest,
			expectedUpdateResult: &domain.Buyer{},
			expectedUpdateError:  buyer.ErrNotFound,
			expectedUpdateCalls:  1,
			expectedCode:         http.StatusNotFound,
		},
		{
			name:                 "Error updating buyer with duplicated card number id",
			id:                   "1",
			updateBuyerRequest:   updateBuyerRequest,
			expectedUpdateResult: &domain.Buyer{},
			expectedUpdateError:  buyer.ErrCardNumberDuplicated,
			expectedUpdateCalls:  1,
			expectedCode:         http.StatusConflict,
		},
		{
			name:                 "Error updating buyer",
			id:                   "1",
			updateBuyerRequest:   updateBuyerRequest,
			expectedUpdateResult: &domain.Buyer{},
			expectedUpdateError:  assert.AnError,
			expectedUpdateCalls:  1,
			expectedCode:         http.StatusInternalServerError,
		},
		//{
		//	name:               "Error invalid buyer",
		//	id:                 "1",
		//	updateBuyerRequest: dtos.UpdateBuyerRequestDTO{},
		//	expectedCode:       http.StatusUnprocessableEntity,
		//},
		{
			name:               "Error invalid id",
			id:                 "xyz",
			updateBuyerRequest: dtos.UpdateBuyerRequestDTO{},
			expectedCode:       http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("*dtos.UpdateBuyerRequestDTO")).Return(test.expectedUpdateResult, test.expectedUpdateError)

			purchaseOrderServiceMock := servicesMocks.NewMockPurchaseOrderService(t)
			buyerHandler := buyers.NewBuyerHandler(buyerServiceMock, purchaseOrderServiceMock)

			//Configurar o servidor

			//// Create custom validation
			//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			//	v.RegisterStructValidation(dtos.UpdateBuyerRequestValidation, dtos.UpdateBuyerRequestDTO{})
			//}

			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.PATCH("/api/v1/buyers/:id", buyerHandler.Update())

			requestBody, _ := json.Marshal(test.updateBuyerRequest)
			request := bytes.NewReader(requestBody)

			//Definir request e response
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%s", "/api/v1/buyers", test.id), request)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			buyerServiceMock.AssertNumberOfCalls(t, "Update", test.expectedUpdateCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var buyerResponse *domain.Buyer
				json.Unmarshal(body, &buyerResponse)

				// Valida o response
				assert.Equal(t, *test.expectedResponse, *buyerResponse)
			}

		})
	}
}

func TestCountPurchaseOrders(t *testing.T) {

	var validBuyer domain.Buyer
	buyerSerialized, _ := os.ReadFile("../../../../test/resources/valid_buyer.json")
	if err := json.Unmarshal(buyerSerialized, &validBuyer); err != nil {
		t.Fatal(err)
	}

	expectedResponse := dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO{
		BuyerID:             1,
		PurchaseOrdersCount: 1,
	}

	tests := []struct {
		name                         string
		id                           string
		expectedCountByBuyerIDResult int
		expectedCountByBuyerIDError  error
		expectedCountByBuyerIDCalls  int
		expectedResponse             dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO
		expectedCode                 int
	}{
		{
			name:                         "Successfully counting purchase orders",
			id:                           "1",
			expectedCountByBuyerIDResult: 1,
			expectedCountByBuyerIDError:  nil,
			expectedCountByBuyerIDCalls:  1,
			expectedResponse:             expectedResponse,
			expectedCode:                 http.StatusOK,
		},
		{
			name:                         "Error buyer not found",
			id:                           "1",
			expectedCountByBuyerIDResult: 0,
			expectedCountByBuyerIDError:  buyer.ErrNotFound,
			expectedCountByBuyerIDCalls:  1,
			expectedResponse:             dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO{},
			expectedCode:                 http.StatusNotFound,
		},
		{
			name:                         "Error counting purchase orders",
			id:                           "1",
			expectedCountByBuyerIDResult: 0,
			expectedCountByBuyerIDError:  assert.AnError,
			expectedCountByBuyerIDCalls:  1,
			expectedResponse:             dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO{},
			expectedCode:                 http.StatusInternalServerError,
		},
		{
			name:         "Invalid ID",
			id:           "xyz",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			buyerServiceMock := mocks.NewBuyerServiceMock()

			purchaseOrderServiceMock := servicesMocks.NewMockPurchaseOrderService(t)
			purchaseOrderServiceMock.On("CountByBuyerID", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedCountByBuyerIDResult, test.expectedCountByBuyerIDError)

			buyerHandler := buyers.NewBuyerHandler(buyerServiceMock, purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/buyers/:id/report-purchase-orders", buyerHandler.CountPurchaseOrders())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/buyers/%s/report-purchase-orders", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			purchaseOrderServiceMock.AssertNumberOfCalls(t, "CountByBuyerID", test.expectedCountByBuyerIDCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {

				//Parsear response
				body, _ := io.ReadAll(res.Body)

				var response struct {
					Data dtos.GetNumberOfPurchaseOrdersByBuyerResponseDTO `json:"data"`
				}

				json.Unmarshal(body, &response)

				// Valida o response
				assert.Equal(t, test.expectedResponse, response.Data)
			}
		})
	}
}
