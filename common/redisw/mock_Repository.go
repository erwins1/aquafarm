// Code generated by mockery v2.14.0. DO NOT EDIT.

package redisw

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// HGet provides a mock function with given fields: ctx, key, field
func (_m *MockRepository) HGet(ctx context.Context, key string, field string) (string, error) {
	ret := _m.Called(ctx, key, field)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, key, field)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, field)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HGetAll provides a mock function with given fields: ctx, key
func (_m *MockRepository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	ret := _m.Called(ctx, key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HIncrBy provides a mock function with given fields: ctx, key, kv
func (_m *MockRepository) HIncrBy(ctx context.Context, key string, kv map[string]interface{}) error {
	ret := _m.Called(ctx, key, kv)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) error); ok {
		r0 = rf(ctx, key, kv)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HMSet provides a mock function with given fields: ctx, key, kv
func (_m *MockRepository) HMSet(ctx context.Context, key string, kv map[string]interface{}) (string, error) {
	ret := _m.Called(ctx, key, kv)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) string); ok {
		r0 = rf(ctx, key, kv)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]interface{}) error); ok {
		r1 = rf(ctx, key, kv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SAdd provides a mock function with given fields: ctx, key, members
func (_m *MockRepository) SAdd(ctx context.Context, key string, members []string) (int64, error) {
	ret := _m.Called(ctx, key, members)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, []string) int64); ok {
		r0 = rf(ctx, key, members)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string) error); ok {
		r1 = rf(ctx, key, members)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SIsMember provides a mock function with given fields: ctx, key, member
func (_m *MockRepository) SIsMember(ctx context.Context, key string, member string) (bool, error) {
	ret := _m.Called(ctx, key, member)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, key, member)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, member)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepository(t mockConstructorTestingTNewMockRepository) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
