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
	//TODO implement me
	panic("implement me")
}

func (repository *BuyerRepositoryMock) Get(ctx context.Context, id int) (domain.Buyer, error) {
	args := repository.Called(ctx, id)

	return args.Get(0).(domain.Buyer), args.Error(1)
}

func (repository *BuyerRepositoryMock) Exists(ctx context.Context, cardNumberID string) bool {
	//TODO implement me
	panic("implement me")
}

func (repository *BuyerRepositoryMock) Save(ctx context.Context, b domain.Buyer) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (repository *BuyerRepositoryMock) Update(ctx context.Context, b domain.Buyer) error {
	//TODO implement me
	panic("implement me")
}

func (repository *BuyerRepositoryMock) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
