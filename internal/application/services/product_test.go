package services_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/repositories/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/services"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/integrations/http/handlers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_productService_Get(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		expectedProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedProduct, nil)

		service := services.NewProductService(productRepositoryMock)
		productReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedProduct, *productReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Product{}, sql.ErrNoRows)

		service := services.NewProductService(productRepositoryMock)
		productReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, productReceived)
		assert.Equal(t, services.ErrNotFound, err)
	})

	t.Run("get_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Product{}, errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		productReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, productReceived)
		assert.Equal(t, errors.New("error"), err)
	})
}

func Test_productService_GetAll(t *testing.T) {
	t.Run("getAll_find_all", func(t *testing.T) {
		expectedProducts := &[]entities.Product{
			{
				ID:             1,
				Description:    "Test",
				ExpirationRate: 1,
				FreezingRate:   1,
				Height:         1.1,
				Length:         1.1,
				Netweight:      1.1,
				ProductCode:    "Teste",
				RecomFreezTemp: 1.1,
				Width:          1.1,
				ProductTypeID:  1,
				SellerID:       1,
			},
			{
				Description:    "Teste",
				ExpirationRate: 1,
				FreezingRate:   1,
				Height:         1.1,
				Length:         1.1,
				Netweight:      1.1,
				ProductCode:    "Teste",
				RecomFreezTemp: 1.1,
				Width:          1.1,
				ProductTypeID:  1,
				SellerID:       1,
			},
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("GetAll", ctx).Return(*expectedProducts, nil)

		service := services.NewProductService(productRepositoryMock)
		productsReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedProducts, *productsReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("getAll_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("GetAll", ctx).Return([]entities.Product{}, errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		productsReceived, err := service.GetAll(&ctx)

		assert.Nil(t, productsReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func Test_productService_Delete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := services.NewProductService(productRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(sql.ErrNoRows)

		service := services.NewProductService(productRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, services.ErrNotFound, err)
	})

	t.Run("delete_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, errors.New("error"), err)
	})
}

func Test_productService_Create(t *testing.T) {
	t.Run("create_conflict", func(t *testing.T) {

		createProductRequestDTO := handlers.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewProductService(productRepositoryMock)
		productSaved, err := service.Save(&ctx, createProductRequestDTO.Description, createProductRequestDTO.ExpirationRate, createProductRequestDTO.FreezingRate,
			createProductRequestDTO.Height, createProductRequestDTO.Length, createProductRequestDTO.Netweight, createProductRequestDTO.ProductCode,
			createProductRequestDTO.RecomFreezTemp, createProductRequestDTO.Width, createProductRequestDTO.ProductTypeID, createProductRequestDTO.SellerID)

		assert.Equal(t, services.ErrConflict, err)
		assert.Nil(t, productSaved)

	})

	t.Run("create_error", func(t *testing.T) {

		createProductRequestDTO := handlers.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		productRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Product")).Return(0, errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		productSaved, err := service.Save(&ctx, createProductRequestDTO.Description, createProductRequestDTO.ExpirationRate, createProductRequestDTO.FreezingRate,
			createProductRequestDTO.Height, createProductRequestDTO.Length, createProductRequestDTO.Netweight, createProductRequestDTO.ProductCode,
			createProductRequestDTO.RecomFreezTemp, createProductRequestDTO.Width, createProductRequestDTO.ProductTypeID, createProductRequestDTO.SellerID)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productSaved)

	})

	t.Run("create_error_get_product", func(t *testing.T) {

		createProductRequestDTO := handlers.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		productRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Product")).Return(1, nil)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Product{}, errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		productSaved, err := service.Save(&ctx, createProductRequestDTO.Description, createProductRequestDTO.ExpirationRate, createProductRequestDTO.FreezingRate,
			createProductRequestDTO.Height, createProductRequestDTO.Length, createProductRequestDTO.Netweight, createProductRequestDTO.ProductCode,
			createProductRequestDTO.RecomFreezTemp, createProductRequestDTO.Width, createProductRequestDTO.ProductTypeID, createProductRequestDTO.SellerID)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productSaved)

	})

	t.Run("create_ok", func(t *testing.T) {
		expectedProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		createProductRequestDTO := handlers.RequestCreateProduct{
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		productRepositoryMock.On("Save", ctx, mock.AnythingOfType("entities.Product")).Return(1, nil)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedProduct, nil)

		service := services.NewProductService(productRepositoryMock)
		productSaved, err := service.Save(&ctx, createProductRequestDTO.Description, createProductRequestDTO.ExpirationRate, createProductRequestDTO.FreezingRate,
			createProductRequestDTO.Height, createProductRequestDTO.Length, createProductRequestDTO.Netweight, createProductRequestDTO.ProductCode,
			createProductRequestDTO.RecomFreezTemp, createProductRequestDTO.Width, createProductRequestDTO.ProductTypeID, createProductRequestDTO.SellerID)

		assert.Equal(t, productSaved, expectedProduct)
		assert.Nil(t, err)

	})
}

func Test_productService_Update(t *testing.T) {
	t.Run("update_existent", func(t *testing.T) {
		originalProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		updateProductRequestDTO := handlers.RequestUpdateProduct{
			Description:    &originalProduct.Description,
			ExpirationRate: &originalProduct.ExpirationRate,
			FreezingRate:   &originalProduct.FreezingRate,
			Height:         &originalProduct.Height,
			Length:         &originalProduct.Length,
			Netweight:      &originalProduct.Netweight,
			ProductCode:    &originalProduct.ProductCode,
			RecomFreezTemp: &originalProduct.RecomFreezTemp,
			Width:          &originalProduct.Width,
			ProductTypeID:  &originalProduct.ProductTypeID,
			SellerID:       &originalProduct.SellerID,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProduct, nil)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		productRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Product")).Return(nil)

		service := services.NewProductService(productRepositoryMock)
		productUpdate, err := service.Update(&ctx, updateProductRequestDTO.Description, updateProductRequestDTO.ExpirationRate, updateProductRequestDTO.FreezingRate,
			updateProductRequestDTO.Height, updateProductRequestDTO.Length, updateProductRequestDTO.Netweight, updateProductRequestDTO.ProductCode,
			updateProductRequestDTO.RecomFreezTemp, updateProductRequestDTO.Width, updateProductRequestDTO.ProductTypeID, updateProductRequestDTO.SellerID, 1)

		assert.Equal(t, productUpdate, originalProduct)
		assert.Nil(t, err)
	})

	t.Run("update_non_existent", func(t *testing.T) {
		originalProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		updateProductRequestDTO := handlers.RequestUpdateProduct{
			Description:    &originalProduct.Description,
			ExpirationRate: &originalProduct.ExpirationRate,
			FreezingRate:   &originalProduct.FreezingRate,
			Height:         &originalProduct.Height,
			Length:         &originalProduct.Length,
			Netweight:      &originalProduct.Netweight,
			ProductCode:    &originalProduct.ProductCode,
			RecomFreezTemp: &originalProduct.RecomFreezTemp,
			Width:          &originalProduct.Width,
			ProductTypeID:  &originalProduct.ProductTypeID,
			SellerID:       &originalProduct.SellerID,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProduct, nil)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		productRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Product")).Return(sql.ErrNoRows)

		service := services.NewProductService(productRepositoryMock)
		productUpdate, err := service.Update(&ctx, updateProductRequestDTO.Description, updateProductRequestDTO.ExpirationRate, updateProductRequestDTO.FreezingRate,
			updateProductRequestDTO.Height, updateProductRequestDTO.Length, updateProductRequestDTO.Netweight, updateProductRequestDTO.ProductCode,
			updateProductRequestDTO.RecomFreezTemp, updateProductRequestDTO.Width, updateProductRequestDTO.ProductTypeID, updateProductRequestDTO.SellerID, 1)

		assert.Equal(t, services.ErrNotFound, err)
		assert.Nil(t, productUpdate)
	})

	t.Run("update_unexpected_error", func(t *testing.T) {
		originalProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		updateProductRequestDTO := handlers.RequestUpdateProduct{
			Description:    &originalProduct.Description,
			ExpirationRate: &originalProduct.ExpirationRate,
			FreezingRate:   &originalProduct.FreezingRate,
			Height:         &originalProduct.Height,
			Length:         &originalProduct.Length,
			Netweight:      &originalProduct.Netweight,
			ProductCode:    &originalProduct.ProductCode,
			RecomFreezTemp: &originalProduct.RecomFreezTemp,
			Width:          &originalProduct.Width,
			ProductTypeID:  &originalProduct.ProductTypeID,
			SellerID:       &originalProduct.SellerID,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProduct, nil)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(false)
		productRepositoryMock.On("Update", ctx, mock.AnythingOfType("entities.Product")).Return(errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		productUpdate, err := service.Update(&ctx, updateProductRequestDTO.Description, updateProductRequestDTO.ExpirationRate, updateProductRequestDTO.FreezingRate,
			updateProductRequestDTO.Height, updateProductRequestDTO.Length, updateProductRequestDTO.Netweight, updateProductRequestDTO.ProductCode,
			updateProductRequestDTO.RecomFreezTemp, updateProductRequestDTO.Width, updateProductRequestDTO.ProductTypeID, updateProductRequestDTO.SellerID, 1)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productUpdate)
	})

	t.Run("update_get_error", func(t *testing.T) {
		originalProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		updateProductRequestDTO := handlers.RequestUpdateProduct{
			Description:    &originalProduct.Description,
			ExpirationRate: &originalProduct.ExpirationRate,
			FreezingRate:   &originalProduct.FreezingRate,
			Height:         &originalProduct.Height,
			Length:         &originalProduct.Length,
			Netweight:      &originalProduct.Netweight,
			ProductCode:    &originalProduct.ProductCode,
			RecomFreezTemp: &originalProduct.RecomFreezTemp,
			Width:          &originalProduct.Width,
			ProductTypeID:  &originalProduct.ProductTypeID,
			SellerID:       &originalProduct.SellerID,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(entities.Product{}, errors.New("error"))

		service := services.NewProductService(productRepositoryMock)
		productUpdate, err := service.Update(&ctx, updateProductRequestDTO.Description, updateProductRequestDTO.ExpirationRate, updateProductRequestDTO.FreezingRate,
			updateProductRequestDTO.Height, updateProductRequestDTO.Length, updateProductRequestDTO.Netweight, updateProductRequestDTO.ProductCode,
			updateProductRequestDTO.RecomFreezTemp, updateProductRequestDTO.Width, updateProductRequestDTO.ProductTypeID, updateProductRequestDTO.SellerID, 1)

		assert.Equal(t, errors.New("error"), err)
		assert.Nil(t, productUpdate)
	})

	t.Run("update_different_product_code", func(t *testing.T) {
		originalProduct := &entities.Product{
			ID:             1,
			Description:    "Test",
			ExpirationRate: 1,
			FreezingRate:   1,
			Height:         1.1,
			Length:         1.1,
			Netweight:      1.1,
			ProductCode:    "Test",
			RecomFreezTemp: 1.1,
			Width:          1.1,
			ProductTypeID:  1,
			SellerID:       1,
		}
		productCode := "Test1"
		updateProductRequestDTO := handlers.RequestUpdateProduct{
			Description:    &originalProduct.Description,
			ExpirationRate: &originalProduct.ExpirationRate,
			FreezingRate:   &originalProduct.FreezingRate,
			Height:         &originalProduct.Height,
			Length:         &originalProduct.Length,
			Netweight:      &originalProduct.Netweight,
			ProductCode:    &productCode,
			RecomFreezTemp: &originalProduct.RecomFreezTemp,
			Width:          &originalProduct.Width,
			ProductTypeID:  &originalProduct.ProductTypeID,
			SellerID:       &originalProduct.SellerID,
		}

		ctx := context.TODO()

		productRepositoryMock := mocks.NewProductRepositoryMock(t)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*originalProduct, nil)
		productRepositoryMock.On("Exists", ctx, mock.AnythingOfType("string")).Return(true)

		service := services.NewProductService(productRepositoryMock)
		productUpdate, err := service.Update(&ctx, updateProductRequestDTO.Description, updateProductRequestDTO.ExpirationRate, updateProductRequestDTO.FreezingRate,
			updateProductRequestDTO.Height, updateProductRequestDTO.Length, updateProductRequestDTO.Netweight, updateProductRequestDTO.ProductCode,
			updateProductRequestDTO.RecomFreezTemp, updateProductRequestDTO.Width, updateProductRequestDTO.ProductTypeID, updateProductRequestDTO.SellerID, 1)

		assert.Equal(t, services.ErrConflict, err)
		assert.Nil(t, productUpdate)
	})
}
