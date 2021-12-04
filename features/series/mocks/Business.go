// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	articles "prog/features/articles"

	mock "github.com/stretchr/testify/mock"

	series "prog/features/series"
)

// Business is an autogenerated mock type for the Business type
type Business struct {
	mock.Mock
}

// AddArticleSeries provides a mock function with given fields: data
func (_m *Business) AddArticleSeries(data series.ArticlesSeriesCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(series.ArticlesSeriesCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateSeries provides a mock function with given fields: data
func (_m *Business) CreateSeries(data series.SeriesCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(series.SeriesCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteArticleSeries provides a mock function with given fields: data
func (_m *Business) DeleteArticleSeries(data series.ArticlesSeriesCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(series.ArticlesSeriesCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSeries provides a mock function with given fields: seriesId
func (_m *Business) DeleteSeries(seriesId int) error {
	ret := _m.Called(seriesId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(seriesId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllArticleSeries provides a mock function with given fields: seriesId
func (_m *Business) GetAllArticleSeries(seriesId int) ([]articles.ArticleCore, error) {
	ret := _m.Called(seriesId)

	var r0 []articles.ArticleCore
	if rf, ok := ret.Get(0).(func(int) []articles.ArticleCore); ok {
		r0 = rf(seriesId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]articles.ArticleCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(seriesId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllSeries provides a mock function with given fields:
func (_m *Business) GetAllSeries() ([]series.SeriesCore, error) {
	ret := _m.Called()

	var r0 []series.SeriesCore
	if rf, ok := ret.Get(0).(func() []series.SeriesCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]series.SeriesCore)
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

// GetSeriesById provides a mock function with given fields: seriesId
func (_m *Business) GetSeriesById(seriesId int) (series.SeriesCore, error) {
	ret := _m.Called(seriesId)

	var r0 series.SeriesCore
	if rf, ok := ret.Get(0).(func(int) series.SeriesCore); ok {
		r0 = rf(seriesId)
	} else {
		r0 = ret.Get(0).(series.SeriesCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(seriesId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSeriesById provides a mock function with given fields: seriesId, data
func (_m *Business) UpdateSeriesById(seriesId int, data series.SeriesCore) error {
	ret := _m.Called(seriesId, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, series.SeriesCore) error); ok {
		r0 = rf(seriesId, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifySeriesOwner provides a mock function with given fields: seriesId, userId
func (_m *Business) VerifySeriesOwner(seriesId int, userId int) error {
	ret := _m.Called(seriesId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(seriesId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
