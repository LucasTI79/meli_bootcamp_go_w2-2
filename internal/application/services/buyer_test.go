package services_test

import (
	"context"
	"database/sql"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_buyerService_Get(t *testing.T) {

	expectedBuyer := entities.Buyer{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Test",
		LastName:     "Test",
	}

	tests := []struct {
		name string
		id   int
		//Mocking repository.Get
		expectedGetResult entities.Buyer
		expectedGetError  error
		expectedGetCalls  int
		//Asserting function
		expectedBuyer *entities.Buyer
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
			expectedGetResult: entities.Buyer{},
			expectedGetError:  sql.ErrNoRows,
			expectedGetCalls:  1,
			expectedBuyer:     &entities.Buyer{},
			expectedError:     services.ErrNotFound,
		},
		{
			name:              "Error connecting db",
			id:                1,
			expectedGetResult: entities.Buyer{},
			expectedGetError:  assert.AnError,
			expectedGetCalls:  1,
			expectedBuyer:     &entities.Buyer{},
			expectedError:     assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			buyerRepositoryMock := mocks.NewBuyerRepositoryMock(t)
			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)

			service := services.NewBuyerService(buyerRepositoryMock)
			buyerReceived, err := service.Get(&ctx, test.id)

			assert.Equal(t, *test.expectedBuyer, *buyerReceived)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
		})
	}
}

func Test_buyerService_GetAll(t *testing.T) {

	expectedBuyers := []entities.Buyer{
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
		expectedGetAllResult []entities.Buyer
		expectedGetAllError  error
		expectedGetAllCalls  int
		//Asserting function
		expectedBuyers *[]entities.Buyer
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
			expectedGetAllResult: []entities.Buyer{},
			expectedGetAllError:  assert.AnError,
			expectedGetAllCalls:  1,
			expectedBuyers:       &[]entities.Buyer{},
			expectedError:        assert.AnError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := context.TODO()

			buyerRepositoryMock := mocks.NewBuyerRepositoryMock(t)
			buyerRepositoryMock.On("GetAll", ctx).Return(test.expectedGetAllResult, test.expectedGetAllError)

			service := services.NewBuyerService(buyerRepositoryMock)
			buyerReceived, err := service.GetAll(&ctx)

			assert.Equal(t, *test.expectedBuyers, *buyerReceived)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "GetAll", test.expectedGetAllCalls)
		})
	}
}

func Test_buyerService_Create(t *testing.T) {

	createBuyerRequest := &dtos.CreateBuyerRequestDTO{
		CardNumberID: "123",
		FirstName:    "Test",
		LastName:     "Test",
	}

	expectedBuyer := &entities.Buyer{
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
		expectedBuyer *entities.Buyer
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
			expectedBuyer:        &entities.Buyer{},
			expectedError:        services.ErrConflict,
		},
		{
			name:                 "Error saving buyer",
			createBuyerRequest:   createBuyerRequest,
			expectedExistsResult: false,
			expectedExistsCalls:  1,
			expectedSaveError:    assert.AnError,
			expectedSaveCalls:    1,
			expectedBuyer:        &entities.Buyer{},
			expectedError:        assert.AnError,
		},
	}

	ctx := context.TODO()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerRepositoryMock := mocks.NewBuyerRepositoryMock(t)

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
				mock.AnythingOfType("entities.Buyer"),
			).Return(
				test.expectedSaveResult,
				test.expectedSaveError,
			)

			service := services.NewBuyerService(buyerRepositoryMock)
			createdBuyer, err := service.Create(&ctx, test.createBuyerRequest)

			assert.Equal(t, *test.expectedBuyer, *createdBuyer)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Exists", test.expectedExistsCalls)
			buyerRepositoryMock.AssertNumberOfCalls(t, "Save", test.expectedSaveCalls)
		})
	}
}

func Test_buyerService_Delete(t *testing.T) {
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
			buyerRepositoryMock := mocks.NewBuyerRepositoryMock(t)
			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Buyer{}, test.expectedGetError)
			buyerRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(test.expectedDeleteError)

			service := services.NewBuyerService(buyerRepositoryMock)
			err := service.Delete(&ctx, test.id)

			assert.Equal(t, test.expectedError, err)
			buyerRepositoryMock.On("Get", test.expectedGetCalls)
			buyerRepositoryMock.On("Delete", test.expectedDeleteCalls)
		})
	}
}

func Test_buyerService_Update(t *testing.T) {

	originalBuyer := entities.Buyer{
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

	updatedBuyer := &entities.Buyer{
		ID:           1,
		CardNumberID: "456",
		FirstName:    "Test2",
		LastName:     "Test2",
	}

	tests := []struct {
		name                 string
		id                   int
		updateBuyerRequest   *dtos.UpdateBuyerRequestDTO
		expectedGetResult    entities.Buyer
		expectedGetError     error
		expectedGetCalls     int
		expectedExistsResult bool
		expectedExistsCalls  int
		expectedUpdateError  error
		expectedUpdateCalls  int
		expectedBuyer        *entities.Buyer
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
			expectedGetResult:  entities.Buyer{},
			expectedGetError:   sql.ErrNoRows,
			expectedGetCalls:   1,
			//expectedExistsResult: false,
			//expectedExistsCalls:  0,
			//expectedUpdateError:  nil,
			//expectedUpdateCalls:  0,
			expectedBuyer: &entities.Buyer{},
			expectedError: services.ErrNotFound,
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
			expectedBuyer: &entities.Buyer{},
			expectedError: services.ErrConflict,
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
			expectedBuyer:        &entities.Buyer{},
			expectedError:        assert.AnError,
		},
	}

	ctx := context.TODO()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			buyerRepositoryMock := mocks.NewBuyerRepositoryMock(t)
			buyerService := services.NewBuyerService(buyerRepositoryMock)

			buyerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(test.expectedGetResult, test.expectedGetError)
			buyerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(test.expectedExistsResult)
			buyerRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Buyer")).Return(test.expectedUpdateError)

			newBuyer, err := buyerService.Update(&ctx, test.id, test.updateBuyerRequest)

			assert.Equal(t, test.expectedBuyer, newBuyer)
			assert.Equal(t, test.expectedError, err)

			buyerRepositoryMock.AssertNumberOfCalls(t, "Get", test.expectedGetCalls)
			buyerRepositoryMock.AssertNumberOfCalls(t, "Exists", test.expectedExistsCalls)
			buyerRepositoryMock.AssertNumberOfCalls(t, "Update", test.expectedUpdateCalls)
		})
	}
}
