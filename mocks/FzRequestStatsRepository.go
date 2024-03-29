// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	models "fizzbuzzlbc/database/models"

	mock "github.com/stretchr/testify/mock"
)

// FzRequestStatsRepository is an autogenerated mock type for the FzRequestStatsRepository type
type FzRequestStatsRepository struct {
	mock.Mock
}

// CreateRequestStats provides a mock function with given fields: _a0
func (_m *FzRequestStatsRepository) CreateRequestStats(_a0 models.FzRequestStat) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateRequestStats")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(models.FzRequestStat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRequestStatsByUid provides a mock function with given fields: uid
func (_m *FzRequestStatsRepository) GetRequestStatsByUid(uid string) (*models.FzRequestStat, error) {
	ret := _m.Called(uid)

	if len(ret) == 0 {
		panic("no return value specified for GetRequestStatsByUid")
	}

	var r0 *models.FzRequestStat
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.FzRequestStat, error)); ok {
		return rf(uid)
	}
	if rf, ok := ret.Get(0).(func(string) *models.FzRequestStat); ok {
		r0 = rf(uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FzRequestStat)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRequestStatsMostCount provides a mock function with given fields:
func (_m *FzRequestStatsRepository) GetRequestStatsMostCount() (*models.FzRequestStat, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRequestStatsMostCount")
	}

	var r0 *models.FzRequestStat
	var r1 error
	if rf, ok := ret.Get(0).(func() (*models.FzRequestStat, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *models.FzRequestStat); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.FzRequestStat)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRequestStats provides a mock function with given fields: _a0
func (_m *FzRequestStatsRepository) UpdateRequestStats(_a0 *models.FzRequestStat) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRequestStats")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.FzRequestStat) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewFzRequestStatsRepository creates a new instance of FzRequestStatsRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFzRequestStatsRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FzRequestStatsRepository {
	mock := &FzRequestStatsRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
