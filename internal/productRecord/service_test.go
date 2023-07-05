package productRecord_test

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/cmd/server/handlers/productsRecords"

	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/productRecord/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		expectedProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Teste2",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedProductRecord, nil)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedProductRecord, *productRecordReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.ProductRecord{}, sql.ErrNoRows)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, productRecordReceived)
		assert.Equal(t, productRecord.ErrNotFound, err)
	})

	t.Run("get_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.ProductRecord{}, errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, productRecordReceived)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("getAll_find_all", func(t *testing.T) {
		expectedProductsRecords := &[]domain.ProductRecord{
			{
				ID:             1,
				LastUpdateRate: "Test",
				PurchasePrice:  1,
				SalePrice:      1,
				ProductId:      1,
			},
			{
				LastUpdateRate: "Test",
				PurchasePrice:  1,
				SalePrice:      1,
				ProductId:      1,
			},
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("GetAll", ctx).Return(*expectedProductsRecords, nil)

		service := productRecord.NewService(productRecordRepositoryMock)
		productsRecordsReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedProductsRecords, *productsRecordsReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("getAll_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("GetAll", ctx).Return([]domain.ProductRecord{}, errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		productsRecordsReceived, err := service.GetAll(&ctx)

		assert.Nil(t, productsRecordsReceived)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := productRecord.NewService(productRecordRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(sql.ErrNoRows)

		service := productRecord.NewService(productRecordRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, productRecord.ErrNotFound, err)
	})
	t.Run("delete_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, errors.New("error"), err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {

		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordSaved, err := service.Save(&ctx, createProductRecordRequestDTO.LastUpdateRate, createProductRecordRequestDTO.PurchasePrice, createProductRecordRequestDTO.SalePrice,
			createProductRecordRequestDTO.ProductId)

		assert.Equal(t, productRecord.ErrConflict, err)
		assert.Nil(t, productRecordSaved)

	})

	t.Run("create_error", func(t *testing.T) {

		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		productRecordRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductRecord")).Return(0, errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordSaved, err := service.Save(&ctx, createProductRecordRequestDTO.LastUpdateRate, createProductRecordRequestDTO.PurchasePrice, createProductRecordRequestDTO.SalePrice,
			createProductRecordRequestDTO.ProductId)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productRecordSaved)

	})

	t.Run("create_error_get_product_record", func(t *testing.T) {

		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		productRecordRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductRecord")).Return(1, nil)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.ProductRecord{}, errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordSaved, err := service.Save(&ctx, createProductRecordRequestDTO.LastUpdateRate, createProductRecordRequestDTO.PurchasePrice, createProductRecordRequestDTO.SalePrice,
			createProductRecordRequestDTO.ProductId)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productRecordSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}
		createProductRecordRequestDTO := productsRecords.RequestCreateProductRecord{
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		productRecordRepositoryMock.On("Save", ctx, mock.AnythingOfType("domain.ProductRecord")).Return(1, nil)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedProductRecord, nil)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordSaved, err := service.Save(&ctx, createProductRecordRequestDTO.LastUpdateRate, createProductRecordRequestDTO.PurchasePrice, createProductRecordRequestDTO.SalePrice,
			createProductRecordRequestDTO.ProductId)

		assert.Equal(t, productRecordSaved, expectedProductRecord)
		assert.Nil(t, err)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		originalProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}
		updateProductRecordRequestDTO := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &originalProductRecord.LastUpdateRate,
			PurchasePrice:  &originalProductRecord.PurchasePrice,
			SalePrice:      &originalProductRecord.SalePrice,
			ProductId:      &originalProductRecord.ProductId,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProductRecord, nil)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		productRecordRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.ProductRecord")).Return(nil)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordUpdate, err := service.Update(&ctx, updateProductRecordRequestDTO.LastUpdateRate, updateProductRecordRequestDTO.PurchasePrice, updateProductRecordRequestDTO.SalePrice,
			updateProductRecordRequestDTO.ProductId, 1)

		assert.Equal(t, productRecordUpdate, originalProductRecord)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		originalProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}
		updateProductRecordRequestDTO := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &originalProductRecord.LastUpdateRate,
			PurchasePrice:  &originalProductRecord.PurchasePrice,
			SalePrice:      &originalProductRecord.SalePrice,
			ProductId:      &originalProductRecord.ProductId,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProductRecord, nil)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		productRecordRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.ProductRecord")).Return(sql.ErrNoRows)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordUpdate, err := service.Update(&ctx, updateProductRecordRequestDTO.LastUpdateRate, updateProductRecordRequestDTO.PurchasePrice, updateProductRecordRequestDTO.SalePrice,
			updateProductRecordRequestDTO.ProductId, 1)

		assert.Equal(t, productRecord.ErrNotFound, err)
		assert.Nil(t, productRecordUpdate)
	})

	t.Run("update_unexpected_error", func(t *testing.T) {
		originalProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}
		updateProductRecordRequestDTO := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &originalProductRecord.LastUpdateRate,
			PurchasePrice:  &originalProductRecord.PurchasePrice,
			SalePrice:      &originalProductRecord.SalePrice,
			ProductId:      &originalProductRecord.ProductId,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProductRecord, nil)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(false)
		productRecordRepositoryMock.On("Update", ctx, mock.AnythingOfType("domain.ProductRecord")).Return(errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordUpdate, err := service.Update(&ctx, updateProductRecordRequestDTO.LastUpdateRate, updateProductRecordRequestDTO.PurchasePrice, updateProductRecordRequestDTO.SalePrice,
			updateProductRecordRequestDTO.ProductId, 1)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productRecordUpdate)
	})

	t.Run("update_get_error", func(t *testing.T) {
		originalProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}
		updateProductRecordRequestDTO := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &originalProductRecord.LastUpdateRate,
			PurchasePrice:  &originalProductRecord.PurchasePrice,
			SalePrice:      &originalProductRecord.SalePrice,
			ProductId:      &originalProductRecord.ProductId,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.ProductRecord{}, errors.New("error"))

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordUpdate, err := service.Update(&ctx, updateProductRecordRequestDTO.LastUpdateRate, updateProductRecordRequestDTO.PurchasePrice, updateProductRecordRequestDTO.SalePrice,
			updateProductRecordRequestDTO.ProductId, 1)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productRecordUpdate)
	})

	t.Run("update_different_product_id", func(t *testing.T) {
		originalProductRecord := &domain.ProductRecord{
			ID:             1,
			LastUpdateRate: "Test",
			PurchasePrice:  1,
			SalePrice:      1,
			ProductId:      1,
		}
		productId := 2
		updateProductRecordRequestDTO := productsRecords.RequestUpdateProductRecord{
			LastUpdateRate: &originalProductRecord.LastUpdateRate,
			PurchasePrice:  &originalProductRecord.PurchasePrice,
			SalePrice:      &originalProductRecord.SalePrice,
			ProductId:      &productId,
		}

		ctx := context.TODO()

		productRecordRepositoryMock := new(mocks.ProductRecordRepositoryMock)
		productRecordRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProductRecord, nil)
		productRecordRepositoryMock.On("Exists", ctx, mock.AnythingOfType("int")).Return(true)

		service := productRecord.NewService(productRecordRepositoryMock)
		productRecordUpdate, err := service.Update(&ctx, updateProductRecordRequestDTO.LastUpdateRate, updateProductRecordRequestDTO.PurchasePrice, updateProductRecordRequestDTO.SalePrice,
			updateProductRecordRequestDTO.ProductId, 1)

		assert.Equal(t, productRecord.ErrConflict, err)
		assert.Nil(t, productRecordUpdate)
	})
}
