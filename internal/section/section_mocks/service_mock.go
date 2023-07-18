package section_mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SectionServiceMock struct {
	mock.Mock
}

func (s *SectionServiceMock) Get(ctx *context.Context, id int) (*domain.Section, error) {
	args := s.Called(ctx, id)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *SectionServiceMock) GetAll(ctx *context.Context) (*[]domain.Section, error) {
	args := s.Called(ctx)
	return args.Get(0).(*[]domain.Section), args.Error(1)
}

func (s *SectionServiceMock) Save(ctx *context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,
	maximumCapacity, warehouseID, productTypeID int) (*domain.Section, error) {
	args := s.Called(ctx, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity,
		maximumCapacity, warehouseID, productTypeID)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *SectionServiceMock) Update(ctx context.Context, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
	warehouseID, productTypeID *int, id int) (*domain.Section, error) {
	args := s.Called(ctx, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity,
		warehouseID, productTypeID, id)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *SectionServiceMock) Delete(ctx *context.Context, id int) error {
	args := s.Called(ctx, id)
	return args.Error(0)
}
