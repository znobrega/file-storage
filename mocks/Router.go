// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	chi "github.com/go-chi/chi/v5"
	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

// Routes provides a mock function with given fields: router
func (_m *Router) Routes(router chi.Router) {
	_m.Called(router)
}