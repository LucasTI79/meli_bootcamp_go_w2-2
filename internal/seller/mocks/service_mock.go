package mocks

import (
	"context"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/sellers"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SellerServiceMock struct {
	mock.Mock
}

func (service *SellerServiceMock) Get(ctx *context.Context, id int) (*domain.Seller, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (service *SellerServiceMock) GetAll(ctx *context.Context) (*[]domain.Seller, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.Seller), args.Error(1)
}

func (service *SellerServiceMock) Save(ctx *context.Context, seller domain.Seller) (*domain.Seller, error) {
	args := service.Called(ctx, seller)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (service *SellerServiceMock) Update(ctx *context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*domain.Seller, error) {
	args := service.Called(ctx, id, updateSellerRequest)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (service *SellerServiceMock) Delete(ctx *context.Context, id int) error {
	args := service.Called(ctx, id)

	return args.Error(0)
}
