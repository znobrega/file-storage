// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "github.com/znobrega/file-storage/pkg/domain/entities"
	dto "github.com/znobrega/file-storage/pkg/dto"
)

// UsersRepository is an autogenerated mock type for the UsersRepository type
type UsersRepository struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *UsersRepository) FindByEmail(email string) (*entities.User, error) {
	ret := _m.Called(email)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(string) *entities.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: userID
func (_m *UsersRepository) FindById(userID int) (*entities.User, error) {
	ret := _m.Called(userID)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(int) *entities.User); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAll provides a mock function with given fields:
func (_m *UsersRepository) ListAll() ([]dto.User, error) {
	ret := _m.Called()

	var r0 []dto.User
	if rf, ok := ret.Get(0).(func() []dto.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: users
func (_m *UsersRepository) Store(users *entities.User) error {
	ret := _m.Called(users)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(users)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: user
func (_m *UsersRepository) UpdateUser(user *entities.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entities.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
