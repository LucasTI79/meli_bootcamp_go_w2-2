package localities_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/localities"
	handlers "github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/localities"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/locality/mocks"
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

	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var validLocality domain.Locality
	if err := json.Unmarshal(localitySerialized, &validLocality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name             string
		id               string
		expectedGetCalls int
		expectedGetError error
		expectedLocality domain.Locality
		expectedCode     int
	}{
		{
			name:             "Success finding locality",
			id:               "1",
			expectedGetCalls: 1,
			expectedLocality: validLocality,
			expectedGetError: nil,
			expectedCode:     http.StatusOK,
		},
		{
			name:             "Error finding locality",
			id:               "999",
			expectedGetCalls: 1,
			expectedGetError: errors.ErrNotFound,
			expectedLocality: domain.Locality{},
			expectedCode:     http.StatusNotFound,
		},
		{
			name:             "Error connecting db",
			id:               "1",
			expectedGetCalls: 1,
			expectedGetError: assert.AnError,
			expectedLocality: domain.Locality{},
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
			localityServiceMock := mocks.NewMockLocalityService(t)
			localityServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedLocality, test.expectedGetError)

			localityHandler := localities.NewLocalityHandler(localityServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/localities/:id", localityHandler.Get())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", "/api/v1/localities", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			localityServiceMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {

				//Parsear response
				body, _ := io.ReadAll(res.Body)

				var response domain.Locality

				json.Unmarshal(body, &response)

				// Valida o response
				assert.Equal(t, test.expectedLocality, response)
			}

		})
	}

}

func TestGetAll(t *testing.T) {

	localitiesSerialized, _ := os.ReadFile("../../../../test/resources/valid_localities.json")
	var validLocalities []domain.Locality
	if err := json.Unmarshal(localitiesSerialized, &validLocalities); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                 string
		expectedGetAllResult []domain.Locality
		expectedGetAllError  error
		expectedGetAllCalls  int
		expectedResponse     []domain.Locality
		expectedCode         int
	}{
		{
			name:                 "Successfully get all expectedLocalities",
			expectedGetAllResult: validLocalities,
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedResponse:     validLocalities,
			expectedCode:         http.StatusOK,
		},
		{
			name:                 "Success empty database",
			expectedGetAllResult: []domain.Locality{},
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedResponse:     []domain.Locality{},
			expectedCode:         http.StatusNoContent,
		},
		{
			name:                 "Error getting all expectedLocalities",
			expectedGetAllResult: []domain.Locality{},
			expectedGetAllError:  assert.AnError,
			expectedGetAllCalls:  1,
			expectedResponse:     []domain.Locality{},
			expectedCode:         http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			localityServiceMock := mocks.NewMockLocalityService(t)
			localityServiceMock.On("GetAll", mock.AnythingOfType("*context.Context")).Return(test.expectedGetAllResult, test.expectedGetAllError)

			localityHandler := localities.NewLocalityHandler(localityServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/expectedLocalities", localityHandler.GetAll())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s", "/api/v1/expectedLocalities"), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			localityServiceMock.AssertNumberOfCalls(t, "GetAll", test.expectedGetAllCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var responseDTO struct {
					Data []domain.Locality `json:"data"`
				}
				json.Unmarshal(body, &responseDTO)

				localityResponse := responseDTO.Data

				// Valida o response
				assert.Equal(t, test.expectedResponse, localityResponse)
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
			name:                "Successfully delete locality",
			id:                  "1",
			expectedDeleteError: nil,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusNoContent,
		},
		{
			name:                "Error deleting inexistent locality",
			id:                  "1",
			expectedDeleteError: errors.ErrNotFound,
			expectedDeleteCalls: 1,
			expectedCode:        http.StatusNotFound,
		},
		{
			name:                "Error deleting locality",
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
			localityServiceMock := mocks.NewMockLocalityService(t)
			localityServiceMock.On("Delete", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedDeleteError)

			localityHandler := localities.NewLocalityHandler(localityServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.DELETE("/api/v1/localities/:id", localityHandler.Delete())

			//Definir request e response
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", "/api/v1/localities", test.id), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			localityServiceMock.AssertNumberOfCalls(t, "Delete", test.expectedDeleteCalls)
			assert.Equal(t, test.expectedCode, res.Code)

		})
	}
}

func TestCreate(t *testing.T) {

	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var locality domain.Locality
	if err := json.Unmarshal(localitySerialized, &locality); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name                  string
		createLocalityRequest domain.Locality
		expectedCreateResult  domain.Locality
		expectedCreateError   error
		expectedCreateCalls   int
		expectedResponse      domain.Locality
		expectedCode          int
	}{
		{
			name:                  "Successfully creating locality",
			createLocalityRequest: locality,
			expectedCreateResult:  locality,
			expectedCreateError:   nil,
			expectedCreateCalls:   1,
			expectedResponse:      locality,
			expectedCode:          http.StatusCreated,
		},
		{
			name:                  "Error creating locality with duplicated id",
			createLocalityRequest: locality,
			expectedCreateResult:  domain.Locality{},
			expectedCreateError:   errors.ErrConflict,
			expectedCreateCalls:   1,
			expectedResponse:      domain.Locality{},
			expectedCode:          http.StatusConflict,
		},
		{
			name:                  "Error creating locality",
			createLocalityRequest: locality,
			expectedCreateResult:  domain.Locality{},
			expectedCreateError:   assert.AnError,
			expectedCreateCalls:   1,
			expectedResponse:      domain.Locality{},
			expectedCode:          http.StatusInternalServerError,
		},
		{
			name:                  "Error invalid locality",
			createLocalityRequest: domain.Locality{},
			expectedResponse:      domain.Locality{},
			expectedCode:          http.StatusUnprocessableEntity,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			localityServiceMock := mocks.NewMockLocalityService(t)
			localityServiceMock.On("Create", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("domain.Locality")).Return(test.expectedCreateResult, test.expectedCreateError)

			localityHandler := handlers.NewLocalityHandler(localityServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.POST("/api/v1/localities", localityHandler.Create())

			requestBody, _ := json.Marshal(test.createLocalityRequest)
			request := bytes.NewReader(requestBody)

			//Definir request e response
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("%s", "/api/v1/localities"), request)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			localityServiceMock.AssertNumberOfCalls(t, "Create", test.expectedCreateCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {

				//Parsear response
				body, _ := io.ReadAll(res.Body)

				var response struct {
					Data domain.Locality `json:"data"`
				}

				json.Unmarshal(body, &response)

				// Valida o response
				assert.Equal(t, test.expectedResponse, response.Data)
			}

		})
	}
}

func TestUpdate(t *testing.T) {

	updateLocalityRequest := dtos.UpdateLocalityRequestDTO{
		ProvinceName: func(s string) *string { return &s }("São Paulo"),
		LocalityName: func(s string) *string { return &s }("São Paulo"),
	}

	localityUpdated := domain.Locality{
		ID:           1,
		CountryName:  "Brasil",
		ProvinceName: "São Paulo",
		LocalityName: "São Paulo",
	}

	tests := []struct {
		name                  string
		id                    string
		updateLocalityRequest interface{}
		expectedUpdateResult  domain.Locality
		expectedUpdateError   error
		expectedUpdateCalls   int
		expectedResponse      domain.Locality
		expectedCode          int
	}{
		{
			name:                  "Successfully updating locality",
			id:                    "1",
			updateLocalityRequest: updateLocalityRequest,
			expectedUpdateResult:  localityUpdated,
			expectedUpdateError:   nil,
			expectedUpdateCalls:   1,
			expectedResponse:      localityUpdated,
			expectedCode:          http.StatusOK,
		},
		{
			name:                  "Error updating inexisting locality",
			id:                    "1",
			updateLocalityRequest: updateLocalityRequest,
			expectedUpdateResult:  domain.Locality{},
			expectedUpdateError:   errors.ErrNotFound,
			expectedUpdateCalls:   1,
			expectedCode:          http.StatusNotFound,
		},
		{
			name:                  "Error updating locality with duplicated card number id",
			id:                    "1",
			updateLocalityRequest: updateLocalityRequest,
			expectedUpdateResult:  domain.Locality{},
			expectedUpdateError:   errors.ErrConflict,
			expectedUpdateCalls:   1,
			expectedCode:          http.StatusConflict,
		},
		{
			name:                  "Error updating locality",
			id:                    "1",
			updateLocalityRequest: updateLocalityRequest,
			expectedUpdateResult:  domain.Locality{},
			expectedUpdateError:   assert.AnError,
			expectedUpdateCalls:   1,
			expectedCode:          http.StatusInternalServerError,
		},
		{
			name:                  "Error invalid locality update request",
			id:                    "1",
			updateLocalityRequest: interface{}(""),
			expectedCode:          http.StatusUnprocessableEntity,
		},
		{
			name:                  "Error invalid id",
			id:                    "xyz",
			updateLocalityRequest: dtos.UpdateLocalityRequestDTO{},
			expectedCode:          http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			localityServiceMock := mocks.NewMockLocalityService(t)
			localityServiceMock.On("Update", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int"), mock.AnythingOfType("dtos.UpdateLocalityRequestDTO")).Return(test.expectedUpdateResult, test.expectedUpdateError)

			localityHandler := localities.NewLocalityHandler(localityServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.PATCH("/api/v1/localities/:id", localityHandler.Update())

			requestBody, _ := json.Marshal(test.updateLocalityRequest)
			request := bytes.NewReader(requestBody)

			//Definir request e response
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("%s/%s", "/api/v1/localities", test.id), request)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			localityServiceMock.AssertNumberOfCalls(t, "Update", test.expectedUpdateCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {
				//Parsear response
				body, _ := io.ReadAll(res.Body)
				var localityResponse domain.Locality
				json.Unmarshal(body, &localityResponse)

				// Valida o response
				assert.Equal(t, test.expectedResponse, localityResponse)
			}

		})
	}
}

func TestGetNumberOfSellers(t *testing.T) {
	localitySerialized, _ := os.ReadFile("../../../../test/resources/valid_locality.json")
	var validLocality domain.Locality
	if err := json.Unmarshal(localitySerialized, &validLocality); err != nil {
		t.Fatal(err)
	}

	expectedResponse := dtos.GetNumberOfSellersResponseDTO{
		LocalityID:   validLocality.ID,
		LocalityName: validLocality.LocalityName,
		SellersCount: 1,
	}

	tests := []struct {
		name                       string
		id                         int
		expectedGetResult          domain.Locality
		expectedGetError           error
		expectedGetCalls           int
		expectedCountSellersResult int
		expectedCountSellersError  error
		expectedCountSellersCalls  int
		expectedResponse           dtos.GetNumberOfSellersResponseDTO
		expectedCode               int
	}{
		{
			name:                       "Valid Locality",
			id:                         validLocality.ID,
			expectedGetResult:          validLocality,
			expectedGetError:           nil,
			expectedGetCalls:           1,
			expectedCountSellersResult: 1,
			expectedCountSellersError:  nil,
			expectedCountSellersCalls:  1,
			expectedResponse:           expectedResponse,
			expectedCode:               http.StatusOK,
		},
		{
			name:                       "Error Locality not found",
			id:                         validLocality.ID,
			expectedGetResult:          domain.Locality{},
			expectedGetError:           errors.ErrNotFound,
			expectedGetCalls:           1,
			expectedCountSellersResult: 0,
			expectedCountSellersError:  nil,
			expectedCountSellersCalls:  0,
			expectedResponse:           dtos.GetNumberOfSellersResponseDTO{},
			expectedCode:               http.StatusNotFound,
		},
		{
			name:                       "Internal error finding Locality",
			id:                         validLocality.ID,
			expectedGetResult:          domain.Locality{},
			expectedGetError:           assert.AnError,
			expectedGetCalls:           1,
			expectedCountSellersResult: 0,
			expectedCountSellersError:  nil,
			expectedCountSellersCalls:  0,
			expectedResponse:           dtos.GetNumberOfSellersResponseDTO{},
			expectedCode:               http.StatusInternalServerError,
		},
		{
			name:                       "Error counting sellers",
			id:                         validLocality.ID,
			expectedGetResult:          validLocality,
			expectedGetError:           nil,
			expectedGetCalls:           1,
			expectedCountSellersResult: 0,
			expectedCountSellersError:  assert.AnError,
			expectedCountSellersCalls:  1,
			expectedResponse:           expectedResponse,
			expectedCode:               http.StatusInternalServerError,
		},
		//{
		//	name:         "Invalid ID",
		//	id:           "xyz",
		//	expectedCode: http.StatusBadRequest,
		//},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			localityServiceMock := mocks.NewMockLocalityService(t)
			localityServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)
			localityServiceMock.On("CountSellers", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedCountSellersResult, test.expectedCountSellersError)

			localityHandler := localities.NewLocalityHandler(localityServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/localities/:id/reportSellers", localityHandler.CountSellers())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d/%s", "/api/v1/localities", test.id, "reportSellers"), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Validar resultado
			localityServiceMock.AssertNumberOfCalls(t, "CountSellers", test.expectedCountSellersCalls)
			assert.Equal(t, test.expectedCode, res.Code)

			// Só testa o body em caso de sucesso
			if test.expectedCode == http.StatusOK {

				//Parsear response
				body, _ := io.ReadAll(res.Body)

				var response struct {
					Data dtos.GetNumberOfSellersResponseDTO `json:"data"`
				}

				json.Unmarshal(body, &response)

				// Valida o response
				assert.Equal(t, test.expectedResponse, response.Data)
			}
		})
	}
}
