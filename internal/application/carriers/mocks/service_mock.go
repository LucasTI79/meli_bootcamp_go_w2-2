package mocks

import (
	"context"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type CarrierServiceMock struct {
	mock.Mock
}

func (service *CarrierServiceMock) Exists(ctx *context.Context, cid string) bool {
	args := service.Called(ctx, cid)

	return args.Get(0).(bool)
}

func (service *CarrierServiceMock) GetAll(ctx *context.Context) (*[]domain.Carrier, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.Carrier), args.Error(1)
}

func (service *CarrierServiceMock) Create(ctx *context.Context, carrier dtos.CarrierRequestDTO) (*domain.Carrier, error) {
	args := service.Called(ctx, carrier)

	return args.Get(0).(*domain.Carrier), args.Error(1)
}
