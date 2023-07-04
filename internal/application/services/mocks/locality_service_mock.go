package mocks

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type LocalityServiceMock struct {
	mock.Mock
}

func NewLocalityServiceMock() *LocalityServiceMock {
	return &LocalityServiceMock{}
}

func (service *LocalityServiceMock) Get(ctx *context.Context, id string) (entities.Locality, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(entities.Locality), args.Error(1)
}

func (service *LocalityServiceMock) GetAll(ctx *context.Context) ([]entities.Locality, error) {
	args := service.Called(ctx)

	return args.Get(0).([]entities.Locality), args.Error(1)
}

func (service *LocalityServiceMock) Create(ctx *context.Context, locality entities.Locality) (entities.Locality, error) {
	args := service.Called(ctx, locality)

	return args.Get(0).(entities.Locality), args.Error(1)
}

func (service *LocalityServiceMock) Update(ctx *context.Context, id string, updateLocalityRequest dtos.UpdateLocalityRequestDTO) (entities.Locality, error) {
	args := service.Called(ctx, id, updateLocalityRequest)

	return args.Get(0).(entities.Locality), args.Error(1)
}

func (service *LocalityServiceMock) Delete(ctx *context.Context, id string) error {
	args := service.Called(ctx, id)

	return args.Error(0)
}

func (service *LocalityServiceMock) CountSellers(ctx *context.Context, id string) (int, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(int), args.Error(1)
}
