// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	follows "prog/features/follows"

	mock "github.com/stretchr/testify/mock"

	users "prog/features/users"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// FollowUser provides a mock function with given fields: data
func (_m *Business) FollowUser(data follows.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(follows.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetFollowersUsers provides a mock function with given fields: followingUserId
func (_m *Business) GetFollowersUsers(followingUserId int) ([]users.Core, error) {
	ret := _m.Called(followingUserId)

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func(int) []users.Core); ok {
		r0 = rf(followingUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(followingUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFollowingUsers provides a mock function with given fields: followersUserId
func (_m *Business) GetFollowingUsers(followersUserId int) ([]users.Core, error) {
	ret := _m.Called(followersUserId)

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func(int) []users.Core); ok {
		r0 = rf(followersUserId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(followersUserId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnfollowUser provides a mock function with given fields: data
func (_m *Business) UnfollowUser(data follows.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(follows.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}