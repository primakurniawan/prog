// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	comments "prog/features/comments"

	mock "github.com/stretchr/testify/mock"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: data
func (_m *Business) AddComment(data comments.Core) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(comments.Core) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: commentId
func (_m *Business) DeleteComment(commentId int) error {
	ret := _m.Called(commentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(commentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetArticleComments provides a mock function with given fields: articleId
func (_m *Business) GetArticleComments(articleId int) ([]comments.Core, error) {
	ret := _m.Called(articleId)

	var r0 []comments.Core
	if rf, ok := ret.Get(0).(func(int) []comments.Core); ok {
		r0 = rf(articleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comments.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(articleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateComment provides a mock function with given fields: commentId, data
func (_m *Business) UpdateComment(commentId int, data comments.Core) error {
	ret := _m.Called(commentId, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, comments.Core) error); ok {
		r0 = rf(commentId, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyCommentOwner provides a mock function with given fields: commentId, userId
func (_m *Business) VerifyCommentOwner(commentId int, userId int) error {
	ret := _m.Called(commentId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(commentId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
