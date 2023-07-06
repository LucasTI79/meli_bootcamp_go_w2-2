package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type InboundOrdersServiceMock struct {
	mock.Mock
}

func NewInboundOrdersServiceMock() *InboundOrdersServiceMock {
	return &InboundOrdersServiceMock{}
}

func (service *InboundOrdersServiceMock) Get(ctx *context.Context, id int) (*domain.InboundOrders, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.InboundOrders), args.Error(1)
}

func (service *InboundOrdersServiceMock) GetAll(ctx *context.Context) (*[]domain.InboundOrders, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.InboundOrders), args.Error(1)
}

func (service *InboundOrdersServiceMock) Save(ctx *context.Context, inboundOrders domain.InboundOrders) (*domain.InboundOrders, error) {
	args := service.Called(ctx, inboundOrders)
	return args.Get(0).(*domain.InboundOrders), args.Error(1)
}

func (service *InboundOrdersServiceMock) Update(ctx *context.Context, id int, reqUpdateInboundOrders *domain.RequestUpdateInboundOrders) (*domain.InboundOrders, error) {
	args := service.Called(ctx, id, reqUpdateInboundOrders)
	return args.Get(0).(*domain.InboundOrders), args.Error(1)
}

func (service *InboundOrdersServiceMock) Delete(ctx *context.Context, id int) error {
	args := service.Called(ctx, id)
	return args.Error(0)
}
