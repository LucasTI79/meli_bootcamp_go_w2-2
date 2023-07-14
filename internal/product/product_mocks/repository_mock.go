package product_mocks

import (
	context "context"

	domain "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type ProductRepositoryMock struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ProductRepositoryMock) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, productCode
func (_m *ProductRepositoryMock) Exists(ctx context.Context, productCode string) bool {
	ret := _m.Called(ctx, productCode)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, productCode)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

func (_m *ProductRepositoryMock) ExistsByID(ctx context.Context, productID int) bool {
	ret := _m.Called(ctx, productID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *ProductRepositoryMock) Get(ctx context.Context, id int) (domain.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (domain.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *ProductRepositoryMock) GetAll(ctx context.Context) ([]domain.Product, error) {
	ret := _m.Called(ctx)

	var r0 []domain.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, p
func (_m *ProductRepositoryMock) Save(ctx context.Context, p domain.Product) (int, error) {
	ret := _m.Called(ctx, p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) (int, error)); ok {
		return rf(ctx, p)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) int); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Product) error); ok {
		r1 = rf(ctx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, p
func (_m *ProductRepositoryMock) Update(ctx context.Context, p domain.Product) error {
	ret := _m.Called(ctx, p)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductRepositoryMock {
	mock := &ProductRepositoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}