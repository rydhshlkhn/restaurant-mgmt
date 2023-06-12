// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/accalina/restaurant-mgmt/entity"
	mock "github.com/stretchr/testify/mock"

	model "github.com/accalina/restaurant-mgmt/model"
)

// OrderItemService is an autogenerated mock type for the OrderItemService type
type OrderItemService struct {
	mock.Mock
}

// CreateOrderItem provides a mock function with given fields: ctx, data
func (_m *OrderItemService) CreateOrderItem(ctx context.Context, data model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error) {
	ret := _m.Called(ctx, data)

	var r0 *entity.OrderItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.OrderItemCreateOrUpdateModel) *entity.OrderItem); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OrderItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.OrderItemCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOrderItem provides a mock function with given fields: ctx, id
func (_m *OrderItemService) DeleteOrderItem(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllOrderItem provides a mock function with given fields: ctx, filter
func (_m *OrderItemService) GetAllOrderItem(ctx context.Context, filter *model.OrderItemFilter) ([]model.OrderItemResponse, model.Meta, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.OrderItemResponse
	var r1 model.Meta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderItemFilter) ([]model.OrderItemResponse, model.Meta, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.OrderItemFilter) []model.OrderItemResponse); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.OrderItemResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.OrderItemFilter) model.Meta); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Get(1).(model.Meta)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *model.OrderItemFilter) error); ok {
		r2 = rf(ctx, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetDetailOrderItem provides a mock function with given fields: ctx, id
func (_m *OrderItemService) GetDetailOrderItem(ctx context.Context, id string) (model.OrderItemResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 model.OrderItemResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.OrderItemResponse, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.OrderItemResponse); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.OrderItemResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrderItem provides a mock function with given fields: ctx, data
func (_m *OrderItemService) UpdateOrderItem(ctx context.Context, data model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error) {
	ret := _m.Called(ctx, data)

	var r0 *entity.OrderItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.OrderItemCreateOrUpdateModel) (*entity.OrderItem, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.OrderItemCreateOrUpdateModel) *entity.OrderItem); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.OrderItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.OrderItemCreateOrUpdateModel) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOrderItemService interface {
	mock.TestingT
	Cleanup(func())
}

// NewOrderItemService creates a new instance of OrderItemService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrderItemService(t mockConstructorTestingTNewOrderItemService) *OrderItemService {
	mock := &OrderItemService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}