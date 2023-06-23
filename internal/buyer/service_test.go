package buyer_test

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/buyer/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGet(t *testing.T) {

	expectedBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Test",
		LastName:     "Test",
	}

	tests := []struct {
		name string
		id   int
		//Mocking repository.Get
		expectedGetResult domain.Buyer
		expectedGetError  error
		expectedGetCalls  int
		//Asserting function
		expectedBuyer *domain.Buyer
		expectedError error
	}{
		{
			name:              "Successfully get buyer from db",
			id:                1,
			expectedGetResult: expectedBuyer,
			expectedGetError:  nil,
			expectedGetCalls:  1,
			expectedBuyer:     &expectedBuyer,
			expectedError:     nil,
		},
		{
			name:              "Buyer not found in db",
			id:                1,
			expectedGetResult: domain.Buyer{},
			expectedGetError:  sql.ErrNoRows,
			expectedGetCalls:  1,
			expectedBuyer:     &domain.Buyer{},
			expectedError:     buyer.ErrNotFound,
		},
		{
			name:              "Error connecting db",
			id:                1,
			expectedGetResult: domain.Buyer{},
			expectedGetError:  assert.AnError,
			expectedGetCalls:  1,
			expectedBuyer:     &domain.Buyer{},
			expectedError:     assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			buyerRepositoryMock := mocks.NewBuyerRepositoryMock()
			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)

			service := buyer.NewService(buyerRepositoryMock)
			buyerReceived, err := service.Get(&ctx, test.id)

			assert.Equal(t, *test.expectedBuyer, *buyerReceived)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
		})
	}
}
