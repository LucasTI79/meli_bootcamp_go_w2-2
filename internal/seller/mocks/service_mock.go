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

func (service *SellerServiceMock) Get(c *context.Context, id int) (*domain.Seller, error) {
	args := service.Called(c, id)

	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (service *SellerServiceMock) GetAll(ctx context.Context) (*[]domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (service *SellerServiceMock) Save(ctx context.Context, seller domain.Seller) (*domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (service *SellerServiceMock) Update(ctx context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*domain.Seller, error) {
	//TODO implement me
	panic("implement me")
}

func (service *SellerServiceMock) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
