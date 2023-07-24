package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type InboundOrdersRepositoryMock struct {
	mock.Mock
}

func NewInboundOrdersRepositoryMock() *InboundOrdersRepositoryMock {
	return &InboundOrdersRepositoryMock{}
}

func (repository *InboundOrdersRepositoryMock) GetAll(ctx context.Context) ([]domain.InboundOrders, error) {
	args := repository.Called(ctx)

	return args.Get(0).([]domain.InboundOrders), args.Error(1)
}

func (repository *InboundOrdersRepositoryMock) Get(ctx context.Context, id int) (domain.InboundOrders, error) {
	args := repository.Called(ctx, id)

	return args.Get(0).(domain.InboundOrders), args.Error(1)
}

func (repository *InboundOrdersRepositoryMock) Exists(ctx context.Context, inboundOrdersCode string) bool {
	args := repository.Called(ctx, inboundOrdersCode)

	return args.Get(0).(bool)
}

func (repository *InboundOrdersRepositoryMock) Save(ctx context.Context, inboundOrders domain.InboundOrders) (int, error) {
	args := repository.Called(ctx, inboundOrders)

	return args.Get(0).(int), args.Error(1)
}

func (repository *InboundOrdersRepositoryMock) Update(ctx context.Context, updateInboundOrdersRequest domain.InboundOrders) error {
	args := repository.Called(ctx, updateInboundOrdersRequest)

	return args.Error(0)
}

func (repository *InboundOrdersRepositoryMock) Delete(ctx context.Context, id int) error {
	args := repository.Called(ctx, id)

	return args.Error(0)

}
