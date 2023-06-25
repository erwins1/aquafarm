// Code generated by mockery v2.14.0. DO NOT EDIT.

package pond

import (
	http "net/http"

	httprouter "github.com/julienschmidt/httprouter"
	mock "github.com/stretchr/testify/mock"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

// HandlerAddPond provides a mock function with given fields: w, r, params
func (_m *MockHandler) HandlerAddPond(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_m.Called(w, r, params)
}

// HandlerDeletePond provides a mock function with given fields: w, r, params
func (_m *MockHandler) HandlerDeletePond(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_m.Called(w, r, params)
}

// HandlerGetPond provides a mock function with given fields: w, r, params
func (_m *MockHandler) HandlerGetPond(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_m.Called(w, r, params)
}

// HandlerGetPondByID provides a mock function with given fields: w, r, params
func (_m *MockHandler) HandlerGetPondByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_m.Called(w, r, params)
}

// HandlerUpsertPond provides a mock function with given fields: w, r, params
func (_m *MockHandler) HandlerUpsertPond(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_m.Called(w, r, params)
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
