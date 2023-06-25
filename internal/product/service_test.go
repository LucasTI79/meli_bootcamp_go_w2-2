package product_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/product/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet(t *testing.T) {

	t.Run("get_find_by_id_existent", func(t *testing.T) {
		expectedProduct := &domain.Product{
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

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(*expectedProduct, nil)

		service := product.NewService(productRepositoryMock)
		productReceived, err := service.Get(&ctx, 1)

		assert.Equal(t, *expectedProduct, *productReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("get_find_by_id_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Product{}, sql.ErrNoRows)

		service := product.NewService(productRepositoryMock)
		productReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, productReceived)
		assert.Equal(t, product.ErrNotFound, err)
	})

	t.Run("get_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("Get", ctx, mock.AnythingOfType("int")).Return(domain.Product{}, errors.New("error"))

		service := product.NewService(productRepositoryMock)
		productReceived, err := service.Get(&ctx, 1)

		assert.Nil(t, productReceived)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("getAll_find_all", func(t *testing.T) {
		expectedProducts := &[]domain.Product{
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

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("GetAll", ctx).Return(*expectedProducts, nil)

		service := product.NewService(productRepositoryMock)
		productsReceived, err := service.GetAll(&ctx)

		assert.Equal(t, *expectedProducts, *productsReceived)
		assert.Equal(t, nil, err)
	})

	t.Run("getAll_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("GetAll", ctx).Return([]domain.Product{}, errors.New("error"))

		service := product.NewService(productRepositoryMock)
		productsReceived, err := service.GetAll(&ctx)

		assert.Nil(t, productsReceived)
		assert.Equal(t, errors.New("error"), err)
	})

}

func TestDelete(t *testing.T) {
	t.Run("delete_ok", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(nil)

		service := product.NewService(productRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(sql.ErrNoRows)

		service := product.NewService(productRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, product.ErrNotFound, err)
	})

	t.Run("delete_unexpected_error", func(t *testing.T) {

		ctx := context.TODO()

		productRepositoryMock := new(mocks.ProductRepositoryMock)
		productRepositoryMock.On("Delete", ctx, mock.AnythingOfType("int")).Return(errors.New("error"))

		service := product.NewService(productRepositoryMock)
		err := service.Delete(&ctx, 1)

		assert.Equal(t, errors.New("error"), err)
	})
}
