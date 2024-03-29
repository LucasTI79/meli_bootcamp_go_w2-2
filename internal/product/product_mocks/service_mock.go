package product_mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (service *ProductServiceMock) Get(ctx *context.Context, id int) (*domain.Product, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.Product), args.Error(1)
}

func (service *ProductServiceMock) GetAll(ctx *context.Context) (*[]domain.Product, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.Product), args.Error(1)
}

func (service *ProductServiceMock) Save(ctx *context.Context, description string, expiration_rate, freezing_rate int, height, length, netweight float32, product_code string,
	recommended_freezing_temperature, width float32, product_type_id, seller_id int) (*domain.Product, error) {
	args := service.Called(ctx, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (service *ProductServiceMock) Update(ctx *context.Context, description *string, expiration_rate, freezing_rate *int, height, length, netweight *float32, product_code *string,
	recommended_freezing_temperature, width *float32, product_type_id, seller_id *int, id int) (*domain.Product, error) {
	args := service.Called(ctx, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (service *ProductServiceMock) Delete(ctx *context.Context, id int) error {
	args := service.Called(ctx, id)

	return args.Error(0)
}
