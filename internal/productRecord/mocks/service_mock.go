package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRecordServiceMock struct {
	mock.Mock
}

func (service *ProductRecordServiceMock) Get(ctx *context.Context, id int) (*domain.ProductRecord, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.ProductRecord), args.Error(1)
}

func (service *ProductRecordServiceMock) GetAll(ctx *context.Context) (*[]domain.ProductRecord, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.ProductRecord), args.Error(1)
}

func (service *ProductRecordServiceMock) Save(ctx *context.Context, lastUpdateRate string, purchasePrice, salePrice float32, productId int) (*domain.ProductRecord, error) {
	args := service.Called(ctx, lastUpdateRate, purchasePrice, salePrice, productId)
	return args.Get(0).(*domain.ProductRecord), args.Error(1)
}

func (service *ProductRecordServiceMock) Update(ctx *context.Context, lastUpdateRate *string, purchasePrice, salePrice *float32, productId *int, id int) (*domain.ProductRecord, error) {
	args := service.Called(ctx, lastUpdateRate, purchasePrice, salePrice, productId)
	return args.Get(0).(*domain.ProductRecord), args.Error(1)
}

func (service *ProductRecordServiceMock) Delete(ctx *context.Context, id int) error {
	args := service.Called(ctx, id)

	return args.Error(0)
}
