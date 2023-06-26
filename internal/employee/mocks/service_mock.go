package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	"github.com/stretchr/testify/mock"
)

type EmployeeServiceMock struct {
	mock.Mock
}

func (service *EmployeeServiceMock) Get(ctx *context.Context, id int) (*domain.Employee, error) {
	args := service.Called(ctx, id)

	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (service *EmployeeServiceMock) GetAll(ctx *context.Context) (*[]domain.Employee, error) {
	args := service.Called(ctx)

	return args.Get(0).(*[]domain.Employee), args.Error(1)
}

func (service *EmployeeServiceMock) Save(ctx *context.Context, employee domain.Employee) (*domain.Employee, error) {
	args := service.Called(ctx, employee)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (service *EmployeeServiceMock) Update(ctx *context.Context, id int, reqUpdateEmployee *domain.RequestUpdateEmployee) (*domain.Employee, error) {
	args := service.Called(ctx, id, reqUpdateEmployee)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (service *EmployeeServiceMock) Delete(ctx *context.Context, id int) error {
	args := service.Called(ctx, id)
	return args.Error(0)
}
