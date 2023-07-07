package services_test

import (
	"context"
	"encoding/json"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_sellerService_Create(t *testing.T) {

	ctx := context.TODO()
	t.Run("create_success", func(t *testing.T) {

		sellersSerialized, _ := os.ReadFile("../../../tests/resources/valid_seller.json")
		expectedSeller := &entities.Seller{}
		if err := json.Unmarshal(sellersSerialized, expectedSeller); err != nil {
			t.Fatal(err)
		}

		sellerRepositoryMock := mocks.NewSellerRepositoryMock(t)
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sellerRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Seller")).Return(1, nil)

		service := services.NewSellerService(sellerRepositoryMock)
		newSeller, err := service.Save(&ctx, *expectedSeller)

		assert.Equal(t, expectedSeller, newSeller)
		assert.Equal(t, nil, err)

	})

	t.Run("error_creating_duplicated_cid", func(t *testing.T) {

		sellersSerialized, _ := os.ReadFile("../../../tests/resources/valid_seller.json")
		var sellerToCreate entities.Seller
		if err := json.Unmarshal(sellersSerialized, &sellerToCreate); err != nil {
			t.Fatal(err)
		}

		sellerRepositoryMock := mocks.NewSellerRepositoryMock(t)
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)

		service := services.NewSellerService(sellerRepositoryMock)
		_, err := service.Save(&ctx, sellerToCreate)

		assert.Equal(t, services.ErrConflict, err)

	})

	t.Run("error_creating_seller", func(t *testing.T) {

		sellersSerialized, _ := os.ReadFile("../../../tests/resources/valid_seller.json")
		var sellerToCreate entities.Seller
		if err := json.Unmarshal(sellersSerialized, &sellerToCreate); err != nil {
			t.Fatal(err)
		}

		sellerRepositoryMock := mocks.NewSellerRepositoryMock(t)
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sellerRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Seller")).Return(1, assert.AnError)

		service := services.NewSellerService(sellerRepositoryMock)
		_, err := service.Save(&ctx, sellerToCreate)

		assert.Equal(t, assert.AnError, err)

	})

}

func Test_sellerService_Get(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		sellersSerialized, _ := os.ReadFile("../../../tests/resources/valid_seller.json")
		expectedSeller := &entities.Seller{}
		if err := json.Unmarshal(sellersSerialized, expectedSeller); err != nil {
			t.Fatal(err)
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(expectedSeller, nil)

		service := services.NewSellerService(sellerRepositoryMock)
		sellerReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedSeller, *sellerReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(&entities.Seller{}, services.ErrNotFound)

		service := services.NewSellerService(sellerRepositoryMock)
		sellerReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, sellerReceived)
		assert.Equal(t, services.ErrNotFound, err)
	})
}

func Test_sellerService_GetAll(t *testing.T) {
	t.Run("getAll_find_all", func(t *testing.T) {
		sellersSerialized, _ := os.ReadFile("../../../tests/resources/valid_sellers.json")
		var expectedSellers []entities.Seller
		if err := json.Unmarshal(sellersSerialized, &expectedSellers); err != nil {
			t.Fatal(err)
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("GetAll", ctx).Return(expectedSellers, nil)

		service := services.NewSellerService(sellerRepositoryMock)
		sellersReceived, err := service.GetAll(&ctx)

		assert.Equal(t, expectedSellers, *sellersReceived)
		assert.Equal(t, nil, err)
	})
}

func Test_sellerService_Delete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {

		sellerSerialized, _ := os.ReadFile("../../../tests/resources/valid_seller.json")
		var sellerToDelete entities.Seller
		if err := json.Unmarshal(sellerSerialized, &sellerToDelete); err != nil {
			t.Fatal(err)
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(&sellerToDelete, nil)
		sellerRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := services.NewSellerService(sellerRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		sellerRepositoryMock := mocks.NewSellerRepositoryMock(t)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(&entities.Seller{}, services.ErrNotFound)

		service := services.NewSellerService(sellerRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, services.ErrNotFound, err)
	})

	t.Run("delete_unexpected_error", func(t *testing.T) {

		sellerSerialized, _ := os.ReadFile("../../../tests/resources/valid_seller.json")
		var sellerToDelete entities.Seller
		if err := json.Unmarshal(sellerSerialized, &sellerToDelete); err != nil {
			t.Fatal(err)
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(&sellerToDelete, nil)
		sellerRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(assert.AnError)

		service := services.NewSellerService(sellerRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, assert.AnError, err)
	})

}

func Test_sellerService_Update(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		originalSeller := &entities.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		newCID := 2
		newCompanyName := "Test2"
		newAddress := "Test2"
		newTelephone := "67890"
		newLocalityID := "456"

		updateSellerRequest := &dtos.UpdateSellerRequestDTO{
			CID:         &newCID,
			CompanyName: &newCompanyName,
			Address:     &newAddress,
			Telephone:   &newTelephone,
			LocalityID:  &newLocalityID,
		}

		expectedSeller := &entities.Seller{
			ID:          1,
			CID:         2,
			CompanyName: "Test2",
			Address:     "Test2",
			Telephone:   "67890",
			LocalityID:  "456",
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(originalSeller, nil)
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sellerRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Seller")).Return(nil)

		service := services.NewSellerService(sellerRepositoryMock)
		updatedSeller, err := service.Update(&ctx, 1, updateSellerRequest)

		assert.Equal(t, *updatedSeller, *expectedSeller)
		assert.Nil(t, err)
	})

	t.Run("update_non_existing", func(t *testing.T) {

		newCID := 2
		newCompanyName := "Test2"
		newAddress := "Test2"
		newTelephone := "67890"
		newLocalityID := "456"

		updateSellerRequest := &dtos.UpdateSellerRequestDTO{
			CID:         &newCID,
			CompanyName: &newCompanyName,
			Address:     &newAddress,
			Telephone:   &newTelephone,
			LocalityID:  &newLocalityID,
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(nil, services.ErrNotFound)

		service := services.NewSellerService(sellerRepositoryMock)
		_, err := service.Update(&ctx, 1, updateSellerRequest)

		assert.Equal(t, services.ErrNotFound, err)
	})
}
