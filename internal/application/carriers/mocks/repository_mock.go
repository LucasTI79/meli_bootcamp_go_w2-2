package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Carrier, error)
	Exists(ctx context.Context, cid string) bool
	Save(ctx context.Context, w domain.Carrier) (int, error)
}

type CarrierRepositoryMock struct {
	mock.Mock
}

func NewCarrierRepositoryMock() *CarrierRepositoryMock {
	return &CarrierRepositoryMock{}
}

func (repository *CarrierRepositoryMock) GetAll(ctx context.Context) ([]domain.Carrier, error) {
	args := repository.Called(ctx)

	return args.Get(0).([]domain.Carrier), args.Error(1)
}

func (repository *CarrierRepositoryMock) Exists(ctx context.Context, cid string) bool {
	args := repository.Called(ctx, cid)

	return args.Get(0).(bool)
}

func (repository *CarrierRepositoryMock) Save(ctx context.Context, carrier domain.Carrier) (int, error) {
	args := repository.Called(ctx, carrier)

	return args.Get(0).(int), args.Error(1)
}
