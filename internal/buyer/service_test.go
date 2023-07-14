package buyer_test

import (
	"context"
	"database/sql"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
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

func TestGetAll(t *testing.T) {

	expectedBuyers := []domain.Buyer{
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
		name string
		//Mocking repository.GetAll
		expectedGetAllResult []domain.Buyer
		expectedGetAllError  error
		expectedGetAllCalls  int
		//Asserting function
		expectedBuyers *[]domain.Buyer
		expectedError  error
	}{
		{
			name:                 "Successfully get all buyers from db",
			expectedGetAllResult: expectedBuyers,
			expectedGetAllError:  nil,
			expectedGetAllCalls:  1,
			expectedBuyers:       &expectedBuyers,
			expectedError:        nil,
		},
		{
			name:                 "Error connecting db",
			expectedGetAllResult: []domain.Buyer{},
			expectedGetAllError:  assert.AnError,
			expectedGetAllCalls:  1,
			expectedBuyers:       &[]domain.Buyer{},
			expectedError:        assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			buyerRepositoryMock := mocks.NewBuyerRepositoryMock()
			buyerRepositoryMock.On("GetAll", ctx).Return(test.expectedGetAllResult, test.expectedGetAllError)

			service := buyer.NewService(buyerRepositoryMock)
			buyerReceived, err := service.GetAll(&ctx)

			assert.Equal(t, *test.expectedBuyers, *buyerReceived)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "GetAll", test.expectedGetAllCalls)
		})
	}
}

func TestCreate(t *testing.T) {

	createBuyerRequest := &dtos.CreateBuyerRequestDTO{
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
		name               string
		createBuyerRequest *dtos.CreateBuyerRequestDTO
		// Mocking repository.Exists
		expectedExistsResult bool
		expectedExistsCalls  int
		// Mocking repository.Save
		expectedSaveResult int
		expectedSaveError  error
		expectedSaveCalls  int
		// Asserting
		expectedBuyer *domain.Buyer
		expectedError error
	}{
		{
			name:                 "Successfully create buyer",
			createBuyerRequest:   createBuyerRequest,
			expectedExistsResult: false,
			expectedExistsCalls:  1,
			expectedSaveResult:   1,
			expectedSaveError:    nil,
			expectedSaveCalls:    1,
			expectedBuyer:        expectedBuyer,
			expectedError:        nil,
		},
		{
			name:                 "Error duplicated card number",
			createBuyerRequest:   createBuyerRequest,
			expectedExistsResult: true,
			expectedExistsCalls:  1,
			expectedSaveResult:   0,
			expectedSaveError:    nil,
			expectedSaveCalls:    0,
			expectedBuyer:        &domain.Buyer{},
			expectedError:        buyer.ErrCardNumberDuplicated,
		},
		{
			name:                 "Error saving buyer",
			createBuyerRequest:   createBuyerRequest,
			expectedExistsResult: false,
			expectedExistsCalls:  1,
			expectedSaveError:    assert.AnError,
			expectedSaveCalls:    1,
			expectedBuyer:        &domain.Buyer{},
			expectedError:        assert.AnError,
		},
	}

	ctx := context.TODO()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerRepositoryMock := mocks.NewBuyerRepositoryMock()

			buyerRepositoryMock.On(
				"Exists",
				ctx,
				mock.AnythingOfType("string"),
			).Return(
				test.expectedExistsResult,
			)

			buyerRepositoryMock.On(
				"Save",
				ctx,
				mock.AnythingOfType("domain.Buyer"),
			).Return(
				test.expectedSaveResult,
				test.expectedSaveError,
			)

			service := buyer.NewService(buyerRepositoryMock)
			createdBuyer, err := service.Create(&ctx, test.createBuyerRequest)

			assert.Equal(t, *test.expectedBuyer, *createdBuyer)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Exists", test.expectedExistsCalls)
			buyerRepositoryMock.AssertNumberOfCalls(t, "Save", test.expectedSaveCalls)
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name                string
		id                  int
		expectedGetError    error
		expectedGetCalls    int
		expectedDeleteError error
		expectedDeleteCalls int
		expectedError       error
	}{
		{
			name:                "Successfully deleting buyer",
			id:                  1,
			expectedGetError:    nil,
			expectedGetCalls:    1,
			expectedDeleteError: nil,
			expectedDeleteCalls: 1,
			expectedError:       nil,
		},
		{
			name:                "Error getting buyer",
			id:                  1,
			expectedGetError:    assert.AnError,
			expectedGetCalls:    1,
			expectedDeleteError: nil,
			expectedDeleteCalls: 0,
			expectedError:       assert.AnError,
		},
		{
			name:                "Error deleting buyer",
			id:                  1,
			expectedGetError:    nil,
			expectedGetCalls:    1,
			expectedDeleteError: assert.AnError,
			expectedDeleteCalls: 0,
			expectedError:       assert.AnError,
		},
	}

	ctx := context.TODO()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerRepositoryMock := mocks.NewBuyerRepositoryMock()
			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Buyer{}, test.expectedGetError)
			buyerRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(test.expectedDeleteError)

			service := buyer.NewService(buyerRepositoryMock)
			err := service.Delete(&ctx, test.id)

			assert.Equal(t, test.expectedError, err)
			buyerRepositoryMock.On("Get", test.expectedGetCalls)
			buyerRepositoryMock.On("Delete", test.expectedDeleteCalls)
		})
	}
}

func TestUpdate(t *testing.T) {

	originalBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Test",
		LastName:     "Test",
	}

	updateBuyerRequest := &dtos.UpdateBuyerRequestDTO{
		CardNumberID: func(s string) *string { return &s }("456"),
		FirstName:    func(s string) *string { return &s }("Test2"),
		LastName:     func(s string) *string { return &s }("Test2"),
	}

	updatedBuyer := &domain.Buyer{
		ID:           1,
		CardNumberID: "456",
		FirstName:    "Test2",
		LastName:     "Test2",
	}

	tests := []struct {
		name                 string
		id                   int
		updateBuyerRequest   *dtos.UpdateBuyerRequestDTO
		expectedGetResult    domain.Buyer
		expectedGetError     error
		expectedGetCalls     int
		expectedExistsResult bool
		expectedExistsCalls  int
		expectedUpdateError  error
		expectedUpdateCalls  int
		expectedBuyer        *domain.Buyer
		expectedError        error
	}{
		{
			name:                 "Successfully updating all fields",
			id:                   originalBuyer.ID,
			updateBuyerRequest:   updateBuyerRequest,
			expectedGetResult:    originalBuyer,
			expectedGetError:     nil,
			expectedGetCalls:     1,
			expectedExistsResult: false,
			expectedExistsCalls:  1,
			expectedUpdateError:  nil,
			expectedUpdateCalls:  1,
			expectedBuyer:        updatedBuyer,
			expectedError:        nil,
		},
		{
			name:               "Error buyer doesn't exists",
			id:                 originalBuyer.ID,
			updateBuyerRequest: updateBuyerRequest,
			expectedGetResult:  domain.Buyer{},
			expectedGetError:   sql.ErrNoRows,
			expectedGetCalls:   1,
			//expectedExistsResult: false,
			//expectedExistsCalls:  0,
			//expectedUpdateError:  nil,
			//expectedUpdateCalls:  0,
			expectedBuyer: &domain.Buyer{},
			expectedError: buyer.ErrNotFound,
		},
		{
			name:                 "Error duplicated card number",
			id:                   originalBuyer.ID,
			updateBuyerRequest:   updateBuyerRequest,
			expectedGetResult:    originalBuyer,
			expectedGetError:     nil,
			expectedGetCalls:     1,
			expectedExistsResult: true,
			expectedExistsCalls:  1,
			//expectedUpdateError:  nil,
			//expectedUpdateCalls:  1,
			expectedBuyer: &domain.Buyer{},
			expectedError: buyer.ErrCardNumberDuplicated,
		},
		{
			name:                 "Error updating buyer",
			id:                   originalBuyer.ID,
			updateBuyerRequest:   updateBuyerRequest,
			expectedGetResult:    originalBuyer,
			expectedGetError:     nil,
			expectedGetCalls:     1,
			expectedExistsResult: false,
			expectedExistsCalls:  1,
			expectedUpdateError:  assert.AnError,
			expectedUpdateCalls:  1,
			expectedBuyer:        &domain.Buyer{},
			expectedError:        assert.AnError,
		},
	}

	ctx := context.TODO()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerRepositoryMock := mocks.NewBuyerRepositoryMock()
			buyerService := buyer.NewService(buyerRepositoryMock)

			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)
			buyerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(test.expectedExistsResult)
			buyerRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Buyer")).Return(test.expectedUpdateError)

			newBuyer, err := buyerService.Update(&ctx, test.id, test.updateBuyerRequest)

			assert.Equal(t, test.expectedBuyer, newBuyer)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
			buyerRepositoryMock.AssertNumberOfCalls(t, "Exists", test.expectedExistsCalls)
			buyerRepositoryMock.AssertNumberOfCalls(t, "Update", test.expectedUpdateCalls)
		})
	}
}
