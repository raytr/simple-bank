// Code generated by mockery v2.14.1. DO NOT EDIT.

package repository_mock

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// BaseRepository is an autogenerated mock type for the BaseRepository type
type BaseRepository struct {
	mock.Mock
}

// GetBegin provides a mock function with given fields:
func (_m *BaseRepository) GetBegin() *gorm.DB {
	ret := _m.Called()

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

type mockConstructorTestingTNewBaseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBaseRepository creates a new instance of BaseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBaseRepository(t mockConstructorTestingTNewBaseRepository) *BaseRepository {
	mock := &BaseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
