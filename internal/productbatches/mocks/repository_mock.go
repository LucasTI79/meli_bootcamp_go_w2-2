package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type ProductBatchesRepositoryMock struct {
	mock.Mock
}

func NewProductBachesRepositoryMock() *ProductBatchesRepositoryMock {
	return &ProductBatchesRepositoryMock{}
}

func (p *ProductBatchesRepositoryMock) Get(ctx context.Context, id int) (domain.ProductBatches, error) {
	args := p.Called(ctx, id)

	return args.Get(0).(domain.ProductBatches), args.Error(1)
}

func (p *ProductBatchesRepositoryMock) ExistsProductBatch(ctx context.Context, batchNumber int) bool {
	args := p.Called(ctx, batchNumber)

	return args.Get(0).(bool)
}

func (p *ProductBatchesRepositoryMock) Save(ctx context.Context, product domain.ProductBatches) (int, error) {
	args := p.Called(ctx, product)

	return args.Get(0).(int), args.Error(1)
}
