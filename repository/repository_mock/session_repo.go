// Code generated by mockery v2.39.1. DO NOT EDIT.

package repository_mock

import (
	context "context"

	entity "gibhub.com/raytr/simple-bank/models/entity"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// SessionRepo is an autogenerated mock type for the SessionRepo type
type SessionRepo struct {
	mock.Mock
}

// CreateWithTx provides a mock function with given fields: ctx, session, tx
func (_m *SessionRepo) CreateWithTx(ctx context.Context, session *entity.Session, tx *gorm.DB) error {
	ret := _m.Called(ctx, session, tx)

	if len(ret) == 0 {
		panic("no return value specified for CreateWithTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Session, *gorm.DB) error); ok {
		r0 = rf(ctx, session, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SessionRepo) Delete(ctx context.Context, id uint64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindById provides a mock function with given fields: ctx, id
func (_m *SessionRepo) FindById(ctx context.Context, id uuid.UUID) (*entity.Session, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 *entity.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entity.Session, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entity.Session); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewSessionRepo creates a new instance of SessionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSessionRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *SessionRepo {
	mock := &SessionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
