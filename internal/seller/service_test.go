package seller_test

import (
	"context"
	"testing"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sellers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/seller/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {

	ctx := context.TODO()
	t.Run("create_success", func(t *testing.T) {

		sellerToCreate := &domain.Seller{
			// ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		expectedSeller := &domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		sellerRepositoryMock := mocks.NewSellerRepositoryMock()
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sellerRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Seller")).Return(1, nil)

		service := seller.NewService(sellerRepositoryMock)
		newSeller, err := service.Save(&ctx, *sellerToCreate)

		assert.Equal(t, *expectedSeller, *newSeller)
		assert.Equal(t, nil, err)

	})

	t.Run("error_creating_duplicated_cid", func(t *testing.T) {

		sellerToCreate := &domain.Seller{
			// ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		sellerRepositoryMock := mocks.NewSellerRepositoryMock()
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)

		service := seller.NewService(sellerRepositoryMock)
		_, err := service.Save(&ctx, *sellerToCreate)

		assert.Equal(t, seller.ErrConflict, err)

	})

	t.Run("error_creating_seller", func(t *testing.T) {

		sellerToCreate := &domain.Seller{
			// ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		sellerRepositoryMock := mocks.NewSellerRepositoryMock()
		sellerRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		sellerRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.Seller")).Return(1, assert.AnError)

		service := seller.NewService(sellerRepositoryMock)
		_, err := service.Save(&ctx, *sellerToCreate)

		assert.Equal(t, assert.AnError, err)

	})

}

func TestGet(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		expectedSeller := &domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "Test",
			LocalityID:  "123",
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(expectedSeller, nil)

		service := seller.NewService(sellerRepositoryMock)
		sellerReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedSeller, *sellerReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(&domain.Seller{}, seller.ErrNotFound)

		service := seller.NewService(sellerRepositoryMock)
		sellerReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, sellerReceived)
		assert.Equal(t, seller.ErrNotFound, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("getAll_find_all", func(t *testing.T) {
		expectedSellers := &[]domain.Seller{
			{
				ID:          1,
				CID:         1,
				CompanyName: "Test",
				Address:     "Test",
				Telephone:   "12345",
				LocalityID:  "123",
			},
			{
				CID:         1,
				CompanyName: "Test",
				Address:     "Test",
				Telephone:   "12345",
				LocalityID:  "123",
			},
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("GetAll", ctx).Return(*expectedSellers, nil)

		service := seller.NewService(sellerRepositoryMock)
		sellersReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedSellers, *sellersReceived)
		assert.Equal(t, nil, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {

		sellerToDelete := &domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(sellerToDelete, nil)
		sellerRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := seller.NewService(sellerRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		sellerRepositoryMock := mocks.NewSellerRepositoryMock()
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(&domain.Seller{}, seller.ErrNotFound)

		service := seller.NewService(sellerRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, seller.ErrNotFound, err)
	})

	t.Run("delete_unexpected_error", func(t *testing.T) {

		sellerToDelete := &domain.Seller{
			ID:          1,
			CID:         1,
			CompanyName: "Test",
			Address:     "Test",
			Telephone:   "12345",
			LocalityID:  "123",
		}

		ctx := context.TODO()

		sellerRepositoryMock := new(mocks.SellerRepositoryMock)
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(sellerToDelete, nil)
		sellerRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(assert.AnError)

		service := seller.NewService(sellerRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, assert.AnError, err)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("update_ok", func(t *testing.T) {
		originalSeller := &domain.Seller{
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

		expectedSeller := &domain.Seller{
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
		sellerRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.Seller")).Return(nil)

		service := seller.NewService(sellerRepositoryMock)
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
		sellerRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(nil, seller.ErrNotFound)

		service := seller.NewService(sellerRepositoryMock)
		_, err := service.Update(&ctx, 1, updateSellerRequest)

		assert.Equal(t, seller.ErrNotFound, err)
	})
}
