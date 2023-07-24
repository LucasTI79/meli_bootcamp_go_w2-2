// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	context "context"
	"github.com/extmatperez/meli_bootcamp_go_w2-2/internal/domain"

	dtos "github.com/extmatperez/meli_bootcamp_go_w2-2/internal/application/dtos/purchase_order"

	mock "github.com/stretchr/testify/mock"
)

// MockPurchaseOrderService is an autogenerated mock type for the PurchaseOrderService type
type MockPurchaseOrderService struct {
	mock.Mock
}

// CountByBuyerID provides a mock function with given fields: ctx, buyerID
func (_m *MockPurchaseOrderService) CountByBuyerID(ctx *context.Context, buyerID int) (int, error) {
	ret := _m.Called(ctx, buyerID)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, int) (int, error)); ok {
		return rf(ctx, buyerID)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, int) int); ok {
		r0 = rf(ctx, buyerID)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*context.Context, int) error); ok {
		r1 = rf(ctx, buyerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *MockPurchaseOrderService) Create(ctx *context.Context, _a1 domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	ret := _m.Called(ctx, _a1)

	var r0 domain.PurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, domain.PurchaseOrder) (domain.PurchaseOrder, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, domain.PurchaseOrder) domain.PurchaseOrder); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(domain.PurchaseOrder)
	}

	if rf, ok := ret.Get(1).(func(*context.Context, domain.PurchaseOrder) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockPurchaseOrderService) Delete(ctx *context.Context, id int) error {
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
func (_m *MockPurchaseOrderService) Get(ctx *context.Context, id int) (domain.PurchaseOrder, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.PurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, int) (domain.PurchaseOrder, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, int) domain.PurchaseOrder); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.PurchaseOrder)
	}

	if rf, ok := ret.Get(1).(func(*context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: ctx
func (_m *MockPurchaseOrderService) GetAll(ctx *context.Context) ([]domain.PurchaseOrder, error) {
	ret := _m.Called(ctx)

	var r0 []domain.PurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context) ([]domain.PurchaseOrder, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(*context.Context) []domain.PurchaseOrder); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PurchaseOrder)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, updatePurchaseOrderRequest
func (_m *MockPurchaseOrderService) Update(ctx *context.Context, id int, updatePurchaseOrderRequest dtos.UpdatePurchaseOrderRequestDTO) (domain.PurchaseOrder, error) {
	ret := _m.Called(ctx, id, updatePurchaseOrderRequest)

	var r0 domain.PurchaseOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, int, dtos.UpdatePurchaseOrderRequestDTO) (domain.PurchaseOrder, error)); ok {
		return rf(ctx, id, updatePurchaseOrderRequest)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, int, dtos.UpdatePurchaseOrderRequestDTO) domain.PurchaseOrder); ok {
		r0 = rf(ctx, id, updatePurchaseOrderRequest)
	} else {
		r0 = ret.Get(0).(domain.PurchaseOrder)
	}

	if rf, ok := ret.Get(1).(func(*context.Context, int, dtos.UpdatePurchaseOrderRequestDTO) error); ok {
		r1 = rf(ctx, id, updatePurchaseOrderRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockPurchaseOrderService creates a new instance of MockPurchaseOrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPurchaseOrderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPurchaseOrderService {
	mock := &MockPurchaseOrderService{}
	mock.Mock.Test(t)

	return mock
}
