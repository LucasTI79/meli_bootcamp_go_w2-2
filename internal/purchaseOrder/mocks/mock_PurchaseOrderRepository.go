// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockPurchaseOrderRepository is an autogenerated mock type for the PurchaseOrderRepository type
type MockPurchaseOrderRepository struct {
	mock.Mock
}

// CountByBuyerID provides a mock function with given fields: ctx, buyerID
func (_m *MockPurchaseOrderRepository) CountByBuyerID(ctx context.Context, buyerID int) (int, error) {
	ret := _m.Called(ctx, buyerID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (int, error)); ok {
		return rf(ctx, buyerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) int); ok {
		r0 = rf(ctx, buyerID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, buyerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockPurchaseOrderRepository) Delete(ctx context.Context, id int) error {
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
func (_m *MockPurchaseOrderRepository) Exists(ctx context.Context, id int) bool {
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
func (_m *MockPurchaseOrderRepository) Get(ctx context.Context, id int) (domain.PurchaseOrder, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.PurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (domain.PurchaseOrder, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) domain.PurchaseOrder); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.PurchaseOrder)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockPurchaseOrderRepository) GetAll(ctx context.Context) ([]domain.PurchaseOrder, error) {
	ret := _m.Called(ctx)

	var r0 []domain.PurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.PurchaseOrder, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.PurchaseOrder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PurchaseOrder)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *MockPurchaseOrderRepository) Save(ctx context.Context, _a1 domain.PurchaseOrder) (int, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PurchaseOrder) (int, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.PurchaseOrder) int); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.PurchaseOrder) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *MockPurchaseOrderRepository) Update(ctx context.Context, _a1 domain.PurchaseOrder) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PurchaseOrder) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockPurchaseOrderRepository creates a new instance of MockPurchaseOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPurchaseOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPurchaseOrderRepository {
	mock := &MockPurchaseOrderRepository{}
	mock.Mock.Test(t)

	return mock
}
