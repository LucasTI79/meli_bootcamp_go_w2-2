package mocks

import (
	"context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type SectionRepositoryMock struct {
	mock.Mock
}

func NewSectionRepositoryMock() *SectionRepositoryMock {
	return &SectionRepositoryMock{}
}

func (r *SectionRepositoryMock) GetAll(ctx context.Context) ([]domain.Section, error) {
	args := r.Called(ctx)

	return args.Get(0).([]domain.Section), args.Error(1)
}

func (r *SectionRepositoryMock) Get(ctx context.Context, id int) (domain.Section, error) {
	args := r.Called(ctx, id)

	return args.Get(0).(domain.Section), args.Error(1)
}

func (r *SectionRepositoryMock) Exists(ctx context.Context, cardNumberID int) bool {
	args := r.Called(ctx, cardNumberID)

	return args.Get(0).(bool)
}

func (r *SectionRepositoryMock) Save(ctx context.Context, section domain.Section) (int, error) {
	args := r.Called(ctx, section)

	return args.Get(0).(int), args.Error(1)
}

func (r *SectionRepositoryMock) Update(ctx context.Context, b domain.Section) error {
	args := r.Called(ctx, b)

	return args.Error(0)
}

func (r *SectionRepositoryMock) Delete(ctx context.Context, id int) error {
	args := r.Called(ctx, id)

	return args.Error(0)
}
