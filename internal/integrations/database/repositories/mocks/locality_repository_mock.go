package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	"github.com/stretchr/testify/mock"
)

type LocalityRepositoryMock struct {
	mock.Mock
}

func NewLocalityRepositoryMock() *LocalityRepositoryMock {
	return &LocalityRepositoryMock{}
}

func (repository *LocalityRepositoryMock) GetAll(ctx context.Context) ([]entities.Locality, error) {
	args := repository.Called(ctx)

	return args.Get(0).([]entities.Locality), args.Error(1)
}

func (repository *LocalityRepositoryMock) Get(ctx context.Context, id int) (entities.Locality, error) {
	args := repository.Called(ctx, id)

	return args.Get(0).(entities.Locality), args.Error(1)
}

func (repository *LocalityRepositoryMock) Exists(ctx context.Context, id int) bool {
	args := repository.Called(ctx, id)

	return args.Get(0).(bool)
}

func (repository *LocalityRepositoryMock) Save(ctx context.Context, buyer entities.Locality) (int, error) {
	args := repository.Called(ctx, buyer)

	return args.Get(0).(int), args.Error(1)
}

func (repository *LocalityRepositoryMock) Update(ctx context.Context, b entities.Locality) error {
	args := repository.Called(ctx, b)

	return args.Error(0)
}

func (repository *LocalityRepositoryMock) Delete(ctx context.Context, id int) error {
	args := repository.Called(ctx, id)

	return args.Error(0)
}

func (repository *LocalityRepositoryMock) GetNumberOfSellers(ctx context.Context, id int) (int, error) {
	args := repository.Called(ctx, id)

	return args.Get(0).(int), args.Error(1)
}
