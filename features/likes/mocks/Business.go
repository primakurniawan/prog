// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	articles "prog/features/articles"

	mock "github.com/stretchr/testify/mock"

	users "prog/features/users"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// GetLikedArticles provides a mock function with given fields: userId
func (_m *Business) GetLikedArticles(userId int) ([]articles.ArticleCore, error) {
	ret := _m.Called(userId)

	var r0 []articles.ArticleCore
	if rf, ok := ret.Get(0).(func(int) []articles.ArticleCore); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]articles.ArticleCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLikingUsers provides a mock function with given fields: articleId
func (_m *Business) GetLikingUsers(articleId int) ([]users.Core, error) {
	ret := _m.Called(articleId)

	var r0 []users.Core
	if rf, ok := ret.Get(0).(func(int) []users.Core); ok {
		r0 = rf(articleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]users.Core)
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

// LikeArticle provides a mock function with given fields: articleId, userId
func (_m *Business) LikeArticle(articleId int, userId int) error {
	ret := _m.Called(articleId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(articleId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnlikeArticle provides a mock function with given fields: articleId, userId
func (_m *Business) UnlikeArticle(articleId int, userId int) error {
	ret := _m.Called(articleId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(articleId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
