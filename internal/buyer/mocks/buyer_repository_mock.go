package mocks

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type BuyerRepositoryMock struct {
	mock.Mock
}

func NewBuyerRepositoryMock() *BuyerRepositoryMock {
	return &BuyerRepositoryMock{}
}

func (repository *BuyerRepositoryMock) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	args := repository.Called(ctx)

	return args.Get(0).([]domain.Buyer), args.Error(1)
}

func (repository *BuyerRepositoryMock) Get(ctx context.Context, id int) (domain.Buyer, error) {
	args := repository.Called(ctx, id)

	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (repository *BuyerRepositoryMock) CardNumberExists(ctx context.Context, cardNumberID string) bool {
	args := repository.Called(ctx, cardNumberID)

	return args.Get(0).(bool)
}

func (repository *BuyerRepositoryMock) Save(ctx context.Context, buyer domain.Buyer) (int, error) {
	args := repository.Called(ctx, buyer)

	return args.Get(0).(int), args.Error(1)
}

func (repository *BuyerRepositoryMock) Update(ctx context.Context, b domain.Buyer) error {
	args := repository.Called(ctx, b)

	return args.Error(0)
}

func (repository *BuyerRepositoryMock) Delete(ctx context.Context, id int) error {
	args := repository.Called(ctx, id)

	return args.Error(0)
}
