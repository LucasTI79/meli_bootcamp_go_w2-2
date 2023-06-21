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
		name               string
		buyerId            int
		expectedGetCalls   int
		expectedGetError   error
		expectedGetResult  *domain.Buyer
		expectedStatusCode int
	}{
		{
			name:               "Success finding buyer",
			buyerId:            1,
			expectedGetCalls:   1,
			expectedGetResult:  &validBuyer,
			expectedGetError:   nil,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "Error finding buyer",
			buyerId:            999,
			expectedGetCalls:   1,
			expectedGetError:   buyer.ErrNotFound,
			expectedGetResult:  &domain.Buyer{},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "Error connecting db",
			buyerId:            1,
			expectedGetCalls:   1,
			expectedGetError:   assert.AnError,
			expectedGetResult:  &domain.Buyer{},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerServiceMock := mocks.NewBuyerServiceMock()
			buyerServiceMock.On("Get", mock.AnythingOfType("context.Context"), mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)

			w := httptest.NewRecorder()

			gin.SetMode(gin.TestMode)
			engine := gin.Default()

			handler := handlers.NewBuyerHandler(buyerServiceMock)
			engine.GET("/api/v1/buyers", handler.Get())

			request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%d", "api/v1/buyers", test.buyerId), nil)

			engine.ServeHTTP(w, request)

			assert.Equal(t, test.expectedStatusCode, w.Code)
			buyerServiceMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)

		})
	}

}
