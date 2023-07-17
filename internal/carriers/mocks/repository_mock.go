package mocks

import (
	"context"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

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

func (repository *CarrierRepositoryMock) GetLocalityById(ctx context.Context, localityId int) (domain.Locality, error) {
	args := repository.Called(ctx, localityId)

	return args.Get(0).(domain.Locality), args.Error(1)
}

func (repository *CarrierRepositoryMock) GetCountCarriersByLocalityId(ctx context.Context, localityId int) (int, error) {
	args := repository.Called(ctx, localityId)

	return args.Get(0).(int), args.Error(1)
}

func (repository *CarrierRepositoryMock) GetCountAndDataByLocality(ctx context.Context) ([]dtos.DataLocalityAndCarrier, error) {
	args := repository.Called(ctx)

	return args.Get(0).([]dtos.DataLocalityAndCarrier), args.Error(1)
}
