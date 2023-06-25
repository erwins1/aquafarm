// Code generated by mockery v2.14.0. DO NOT EDIT.

package log

import mock "github.com/stretchr/testify/mock"

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

type mockConstructorTestingTNewMockHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockHandler(t mockConstructorTestingTNewMockHandler) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}