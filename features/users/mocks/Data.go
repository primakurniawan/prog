// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	users "prog/features/users"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: data
func (_m *Data) CreateUser(data users.Core) (int, error) {
	ret := _m.Called(data)

	var r0 int
	if rf, ok := ret.Get(0).(func(users.Core) int); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(users.Core) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserById provides a mock function with given fields: userId
func (_m *Data) DeleteUserById(userId int) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields:
func (_m *Data) GetAllUsers() ([]users.Core, error) {
	ret := _m.Called()

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func() []users.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
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

// GetUserById provides a mock function with given fields: userId
func (_m *Data) GetUserById(userId int) (users.Core, error) {
	ret := _m.Called(userId)

	var r0 users.Core
	if rf, ok := ret.Get(0).(func(int) users.Core); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(users.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUserById provides a mock function with given fields: userId, data
func (_m *Data) UpdateUserById(userId int, data users.Core) error {
	ret := _m.Called(userId, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, users.Core) error); ok {
		r0 = rf(userId, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
