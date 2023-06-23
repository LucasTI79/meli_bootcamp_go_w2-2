package mocks

import (
	"context"
	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type BuyerServiceMock struct {
	mock.Mock
}

func NewBuyerServiceMock() *BuyerServiceMock {
	return &BuyerServiceMock{}
}

func (service *BuyerServiceMock) Get(ctx *context.Context, id int) (*domain.Buyer, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.Buyer), args.Error(1)
}

func (service *BuyerServiceMock) GetAll(ctx *context.Context) (*[]domain.Buyer, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BuyerServiceMock) Create(ctx context.Context, createBuyerRequest *dtos.CreateBuyerRequestDTO) (*domain.Buyer, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BuyerServiceMock) Update(ctx context.Context, id int, updateBuyerRequest *dtos.UpdateBuyerRequestDTO) (*domain.Buyer, error) {
	//TODO implement me
	panic("implement me")
}

func (service *BuyerServiceMock) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
