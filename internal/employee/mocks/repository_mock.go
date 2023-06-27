package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type EmployeeRepositoryMock struct {
	mock.Mock
}

func NewEmployeeRepositoryMock() *EmployeeRepositoryMock {
	return &EmployeeRepositoryMock{}
}

func (repository *EmployeeRepositoryMock) GetAll(ctx context.Context) ([]domain.Employee, error) {
	args := repository.Called(ctx)

	return args.Get(0).([]domain.Employee), args.Error(1)
}

func (repository *EmployeeRepositoryMock) Get(ctx context.Context, id int) (domain.Employee, error) {
	args := repository.Called(ctx, id)

	return args.Get(0).(domain.Employee), args.Error(1)
}

func (repository *EmployeeRepositoryMock) Exists(ctx context.Context, employeeCode string) bool {
	args := repository.Called(ctx, employeeCode)

	return args.Get(0).(bool)
}

func (repository *EmployeeRepositoryMock) Save(ctx context.Context, employee domain.Employee) (int, error) {
	args := repository.Called(ctx, employee)

	return args.Get(0).(int), args.Error(1)
}

func (repository *EmployeeRepositoryMock) Update(ctx context.Context, updateEmployeeRequest domain.Employee) error {
	args := repository.Called(ctx, updateEmployeeRequest)

	return args.Error(0)
}

func (repository *EmployeeRepositoryMock) Delete(ctx context.Context, id int) error {
	args := repository.Called(ctx, id)

	return args.Error(0)

}
