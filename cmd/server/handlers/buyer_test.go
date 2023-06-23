package handlers_test

import (
	"encoding/json"
	"fmt"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
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

	buyerSerialized, _ := os.ReadFile("../../../test/resources/valid_buyer.json")
	var validBuyer domain.Buyer
	if err := json.Unmarshal(buyerSerialized, &validBuyer); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name             string
		buyerId          int
		expectedGetCalls int
		expectedGetError error
		expectedBuyer    *domain.Buyer
		expectedCode     int
	}{
		{
			name:             "Success finding buyer",
			buyerId:          1,
			expectedGetCalls: 1,
			expectedBuyer:    &validBuyer,
			expectedGetError: nil,
			expectedCode:     http.StatusOK,
		},
		{
			name:             "Error finding buyer",
			buyerId:          999,
			expectedGetCalls: 1,
			expectedGetError: buyer.ErrNotFound,
			expectedBuyer:    &domain.Buyer{},
			expectedCode:     http.StatusNotFound,
		},
		{
			name:             "Error connecting db",
			buyerId:          1,
			expectedGetCalls: 1,
			expectedGetError: assert.AnError,
			expectedBuyer:    &domain.Buyer{},
			expectedCode:     http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("Get", mock.AnythingOfType("*context.Context"), mock.AnythingOfType("int")).Return(test.expectedBuyer, test.expectedGetError)

			buyerHandler := handlers.NewBuyerHandler(buyerServiceMock)

			//Configurar o servidor
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.GET("/api/v1/buyers/:id", buyerHandler.Get())

			//Definir request e response
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", "/api/v1/buyers", test.buyerId), nil)
			res := httptest.NewRecorder()

			//Executar request
			r.ServeHTTP(res, req)

			//Parsear response
			var buyerResponse domain.Buyer
			body, _ := io.ReadAll(res.Body)
			json.Unmarshal(body, &buyerResponse)

			//Validar resultado
			assert.Equal(t, test.expectedCode, res.Code)
			assert.Equal(t, *test.expectedBuyer, buyerResponse)

		})
	}

}
