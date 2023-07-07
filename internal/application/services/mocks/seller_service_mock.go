// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos"
	entities "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// SellerServiceMock is an autogenerated mock type for the SellerService type
type SellerServiceMock struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SellerServiceMock) Delete(ctx *context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *SellerServiceMock) Get(ctx *context.Context, id int) (*entities.Seller, error) {
	ret := _m.Called(ctx, id)

	var r0 *entities.Seller
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, int) (*entities.Seller, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, int) *entities.Seller); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *SellerServiceMock) GetAll(ctx *context.Context) (*[]entities.Seller, error) {
	ret := _m.Called(ctx)

	var r0 *[]entities.Seller
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context) (*[]entities.Seller, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(*context.Context) *[]entities.Seller); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entities.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, seller
func (_m *SellerServiceMock) Save(ctx *context.Context, seller entities.Seller) (*entities.Seller, error) {
	ret := _m.Called(ctx, seller)

	var r0 *entities.Seller
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, entities.Seller) (*entities.Seller, error)); ok {
		return rf(ctx, seller)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, entities.Seller) *entities.Seller); ok {
		r0 = rf(ctx, seller)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context, entities.Seller) error); ok {
		r1 = rf(ctx, seller)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, updateSellerRequest
func (_m *SellerServiceMock) Update(ctx *context.Context, id int, updateSellerRequest *dtos.UpdateSellerRequestDTO) (*entities.Seller, error) {
	ret := _m.Called(ctx, id, updateSellerRequest)

	var r0 *entities.Seller
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, int, *dtos.UpdateSellerRequestDTO) (*entities.Seller, error)); ok {
		return rf(ctx, id, updateSellerRequest)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, int, *dtos.UpdateSellerRequestDTO) *entities.Seller); ok {
		r0 = rf(ctx, id, updateSellerRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Seller)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context, int, *dtos.UpdateSellerRequestDTO) error); ok {
		r1 = rf(ctx, id, updateSellerRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSellerServiceMock creates a new instance of SellerServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSellerServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *SellerServiceMock {
	mock := &SellerServiceMock{}
	mock.Mock.Test(t)

	return mock
}
