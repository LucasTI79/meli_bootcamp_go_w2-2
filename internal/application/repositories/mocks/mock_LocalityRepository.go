// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// MockLocalityRepository is an autogenerated mock type for the LocalityRepository type
type MockLocalityRepository struct {
	mock.Mock
}

// CountSellers provides a mock function with given fields: ctx, id
func (_m *MockLocalityRepository) CountSellers(ctx context.Context, id int) (int, error) {
	ret := _m.Called(ctx, id)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (int, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockLocalityRepository) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, id
func (_m *MockLocalityRepository) Exists(ctx context.Context, id int) bool {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *MockLocalityRepository) Get(ctx context.Context, id int) (entities.Locality, error) {
	ret := _m.Called(ctx, id)

	var r0 entities.Locality
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (entities.Locality, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) entities.Locality); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entities.Locality)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockLocalityRepository) GetAll(ctx context.Context) ([]entities.Locality, error) {
	ret := _m.Called(ctx)

	var r0 []entities.Locality
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.Locality, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.Locality); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Locality)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, locality
func (_m *MockLocalityRepository) Save(ctx context.Context, locality entities.Locality) (int, error) {
	ret := _m.Called(ctx, locality)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Locality) (int, error)); ok {
		return rf(ctx, locality)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Locality) int); ok {
		r0 = rf(ctx, locality)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Locality) error); ok {
		r1 = rf(ctx, locality)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, locality
func (_m *MockLocalityRepository) Update(ctx context.Context, locality entities.Locality) error {
	ret := _m.Called(ctx, locality)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Locality) error); ok {
		r0 = rf(ctx, locality)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockLocalityRepository creates a new instance of MockLocalityRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLocalityRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLocalityRepository {
	mock := &MockLocalityRepository{}
	mock.Mock.Test(t)

	return mock
}
