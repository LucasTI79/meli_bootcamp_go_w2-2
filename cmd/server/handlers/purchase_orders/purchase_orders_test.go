package purchase_orders_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/purchase_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/purchase_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
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

	purchaseOrderSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_order.json")
	var validPurchaseOrder entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrderSerialized, &validPurchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                  string
		id                    string
		expectedGetCalls      int
		expectedGetError      error
		expectedPurchaseOrder entities.PurchaseOrder
		expectedCode          int
	}{
		{
			name:                  "Success finding purchaseOrder",
			id:                    "1",
			expectedGetCalls:      1,
			expectedPurchaseOrder: validPurchaseOrder,
			expectedGetError:      nil,
			expectedCode:          http.StatusOK,
		},
		{
			name:                  "Error finding purchaseOrder",
			id:                    "999",
			expectedGetCalls:      1,
			expectedGetError:      services.ErrNotFound,
			expectedPurchaseOrder: entities.PurchaseOrder{},
			expectedCode:          http.StatusNotFound,
		},
		{
			name:                  "Error connecting db",
			id:                    "1",
			expectedGetCalls:      1,
			expectedGetError:      assert.AnError,
			expectedPurchaseOrder: entities.PurchaseOrder{},
			expectedCode:          http.StatusInternalServerError,
		},
		{
			name:         "Error invalid id",
			id:           "xyz",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			purchaseOrderServiceMock := mocks.NewMockPurchaseOrderService(t)
			purchaseOrderServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedPurchaseOrder, test.expectedGetError)

			purchaseOrderHandler := purchase_orders.NewPurchaseOrderHandler(purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/puchase-orders/:id", purchaseOrderHandler.Get())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", "/api/v1/puchase-orders", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			purchaseOrderServiceMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {

				//Parsear response
				body, _ := io.ReadAll(res.Body)

				var response entities.PurchaseOrder

				json.Unmarshal(body, &response)

				// Valida o response
				assert.Equal(t, test.expectedPurchaseOrder, response)
			}

		})
	}

}

func TestGetAll(t *testing.T) {

	purchaseOrdersSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_orders.json")
	var validPurchaseOrders []entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrdersSerialized, &validPurchaseOrders); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                 string
		expectedGetAllResult []entities.PurchaseOrder
		expectedGetAllError  error
		expectedGetAllCalls  int
		expectedResponse     []entities.PurchaseOrder
		expectedCode         int
	}{
		{
			name:                 "Successfully get all",
			expectedGetAllResult: validPurchaseOrders,
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedResponse:     validPurchaseOrders,
			expectedCode:         http.StatusOK,
		},
		{
			name:                 "Success empty database",
			expectedGetAllResult: []entities.PurchaseOrder{},
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedResponse:     []entities.PurchaseOrder{},
			expectedCode:         http.StatusNoContent,
		},
		{
			name:                 "Error getting all",
			expectedGetAllResult: []entities.PurchaseOrder{},
			expectedGetAllError:  assert.AnError,
			expectedGetAllCalls:  1,
			expectedResponse:     []entities.PurchaseOrder{},
			expectedCode:         http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			purchaseOrderServiceMock := mocks.NewMockPurchaseOrderService(t)
			purchaseOrderServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(test.expectedGetAllResult, test.expectedGetAllError)

			purchaseOrderHandler := purchase_orders.NewPurchaseOrderHandler(purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/purchase-orders", purchaseOrderHandler.GetAll())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s", "/api/v1/purchase-orders"), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			purchaseOrderServiceMock.AssertNumberOfCalls(t, "GetAll", test.expectedGetAllCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var responseDTO struct {
					Data []entities.PurchaseOrder `json:"data"`
				}
				json.Unmarshal(body, &responseDTO)

				purchaseOrderResponse := responseDTO.Data

				// Valida o response
				assert.Equal(t, test.expectedResponse, purchaseOrderResponse)
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
			name:                "Successfully delete purchaseOrder",
			id:                  "1",
			expectedDeleteError: nil,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusNoContent,
		},
		{
			name:                "Error deleting inexistent purchaseOrder",
			id:                  "1",
			expectedDeleteError: services.ErrNotFound,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusNotFound,
		},
		{
			name:                "Error deleting purchaseOrder",
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
			purchaseOrderServiceMock := mocks.NewMockPurchaseOrderService(t)
			purchaseOrderServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedDeleteError)

			purchaseOrderHandler := purchase_orders.NewPurchaseOrderHandler(purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.DELETE("/api/v1/puchase-orders/:id", purchaseOrderHandler.Delete())

			//Definir request e response
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", "/api/v1/puchase-orders", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			purchaseOrderServiceMock.AssertNumberOfCalls(t, "Delete", test.expectedDeleteCalls)
			assert.Equal(t, test.expectedCode, res.Code)

		})
	}
}

func TestCreate(t *testing.T) {

	purchaseOrderSerialized, _ := os.ReadFile("../../../../test/resources/valid_purchase_order.json")
	var purchaseOrder entities.PurchaseOrder
	if err := json.Unmarshal(purchaseOrderSerialized, &purchaseOrder); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                       string
		createPurchaseOrderRequest entities.PurchaseOrder
		expectedCreateResult       entities.PurchaseOrder
		expectedCreateError        error
		expectedCreateCalls        int
		expectedResponse           entities.PurchaseOrder
		expectedCode               int
	}{
		{
			name:                       "Successfully creating purchaseOrder",
			createPurchaseOrderRequest: purchaseOrder,
			expectedCreateResult:       purchaseOrder,
			expectedCreateError:        nil,
			expectedCreateCalls:        1,
			expectedResponse:           purchaseOrder,
			expectedCode:               http.StatusCreated,
		},
		{
			name:                       "Error creating purchaseOrder with duplicated id",
			createPurchaseOrderRequest: purchaseOrder,
			expectedCreateResult:       entities.PurchaseOrder{},
			expectedCreateError:        services.ErrConflict,
			expectedCreateCalls:        1,
			expectedResponse:           entities.PurchaseOrder{},
			expectedCode:               http.StatusConflict,
		},
		{
			name:                       "Error creating purchaseOrder",
			createPurchaseOrderRequest: purchaseOrder,
			expectedCreateResult:       entities.PurchaseOrder{},
			expectedCreateError:        assert.AnError,
			expectedCreateCalls:        1,
			expectedResponse:           entities.PurchaseOrder{},
			expectedCode:               http.StatusInternalServerError,
		},
		{
			name:                       "Error invalid purchaseOrder",
			createPurchaseOrderRequest: entities.PurchaseOrder{},
			expectedResponse:           entities.PurchaseOrder{},
			expectedCode:               http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			purchaseOrderServiceMock := mocks.NewMockPurchaseOrderService(t)
			purchaseOrderServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("entities.PurchaseOrder")).Return(test.expectedCreateResult, test.expectedCreateError)

			purchaseOrderHandler := purchase_orders.NewPurchaseOrderHandler(purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.POST("/api/v1/puchase-orders", purchaseOrderHandler.Create())

			requestBody, _ := json.Marshal(test.createPurchaseOrderRequest)
			request := bytes.NewReader(requestBody)

			//Definir request e response
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s", "/api/v1/puchase-orders"), request)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			purchaseOrderServiceMock.AssertNumberOfCalls(t, "Create", test.expectedCreateCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {

				//Parsear response
				body, _ := io.ReadAll(res.Body)

				var response struct {
					Data entities.PurchaseOrder `json:"data"`
				}

				json.Unmarshal(body, &response)

				// Valida o response
				assert.Equal(t, test.expectedResponse, response.Data)
			}

		})
	}
}

func TestUpdate(t *testing.T) {

	newOrderNumber := "1"
	newOrderDate := "2023-07-07"
	newTrackingCode := "123"
	newBuyerID := 1
	newCarrierID := 1
	newOrderStatusID := 1
	newWarehouseID := 1
	newProductRecordID := 1

	updatePurchaseOrderRequest := dtos.UpdatePurchaseOrderRequestDTO{
		OrderNumber:     &newOrderNumber,
		OrderDate:       &newOrderDate,
		TrackingCode:    &newTrackingCode,
		BuyerID:         &newBuyerID,
		CarrierID:       &newCarrierID,
		OrderStatusID:   &newOrderStatusID,
		WarehouseID:     &newWarehouseID,
		ProductRecordID: &newProductRecordID,
	}

	purchaseOrderUpdated := entities.PurchaseOrder{
		ID:              1,
		OrderNumber:     newOrderNumber,
		OrderDate:       newOrderDate,
		TrackingCode:    newTrackingCode,
		BuyerID:         newBuyerID,
		CarrierID:       newCarrierID,
		OrderStatusID:   newOrderStatusID,
		WarehouseID:     newWarehouseID,
		ProductRecordID: newProductRecordID,
	}

	tests := []struct {
		name                       string
		id                         string
		updatePurchaseOrderRequest interface{}
		expectedUpdateResult       entities.PurchaseOrder
		expectedUpdateError        error
		expectedUpdateCalls        int
		expectedResponse           entities.PurchaseOrder
		expectedCode               int
	}{
		{
			name:                       "Successfully updating purchaseOrder",
			id:                         "1",
			updatePurchaseOrderRequest: updatePurchaseOrderRequest,
			expectedUpdateResult:       purchaseOrderUpdated,
			expectedUpdateError:        nil,
			expectedUpdateCalls:        1,
			expectedResponse:           purchaseOrderUpdated,
			expectedCode:               http.StatusOK,
		},
		{
			name:                       "Error updating inexisting purchaseOrder",
			id:                         "1",
			updatePurchaseOrderRequest: updatePurchaseOrderRequest,
			expectedUpdateResult:       entities.PurchaseOrder{},
			expectedUpdateError:        services.ErrNotFound,
			expectedUpdateCalls:        1,
			expectedCode:               http.StatusNotFound,
		},
		{
			name:                       "Error updating purchaseOrder with duplicated card number id",
			id:                         "1",
			updatePurchaseOrderRequest: updatePurchaseOrderRequest,
			expectedUpdateResult:       entities.PurchaseOrder{},
			expectedUpdateError:        services.ErrConflict,
			expectedUpdateCalls:        1,
			expectedCode:               http.StatusConflict,
		},
		{
			name:                       "Error updating purchaseOrder",
			id:                         "1",
			updatePurchaseOrderRequest: updatePurchaseOrderRequest,
			expectedUpdateResult:       entities.PurchaseOrder{},
			expectedUpdateError:        assert.AnError,
			expectedUpdateCalls:        1,
			expectedCode:               http.StatusInternalServerError,
		},
		{
			name:                       "Error invalid purchaseOrder update request",
			id:                         "1",
			updatePurchaseOrderRequest: interface{}(""),
			expectedCode:               http.StatusUnprocessableEntity,
		},
		{
			name:                       "Error invalid id",
			id:                         "xyz",
			updatePurchaseOrderRequest: dtos.UpdatePurchaseOrderRequestDTO{},
			expectedCode:               http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			purchaseOrderServiceMock := mocks.NewMockPurchaseOrderService(t)
			purchaseOrderServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("dtos.UpdatePurchaseOrderRequestDTO")).Return(test.expectedUpdateResult, test.expectedUpdateError)

			purchaseOrderHandler := purchase_orders.NewPurchaseOrderHandler(purchaseOrderServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.PATCH("/api/v1/puchase-orders/:id", purchaseOrderHandler.Update())

			requestBody, _ := json.Marshal(test.updatePurchaseOrderRequest)
			request := bytes.NewReader(requestBody)

			//Definir request e response
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%s", "/api/v1/puchase-orders", test.id), request)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			purchaseOrderServiceMock.AssertNumberOfCalls(t, "Update", test.expectedUpdateCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var purchaseOrderResponse entities.PurchaseOrder
				json.Unmarshal(body, &purchaseOrderResponse)

				// Valida o response
				assert.Equal(t, test.expectedResponse, purchaseOrderResponse)
			}

		})
	}
}
