// Code generated by mockery v2.34.2. DO NOT EDIT.

package repo

import (
	context "context"
	dto "userAuthApp/model/dto"

	entity "userAuthApp/model/entity"

	mock "github.com/stretchr/testify/mock"

	sql "database/sql"
)

// InterfaceUserRepository is an autogenerated mock type for the InterfaceUserRepository type
type InterfaceUserRepository struct {
	mock.Mock
}

// GetByAccountId provides a mock function with given fields: ctx, tx, accountId
func (_m *InterfaceUserRepository) GetByAccountId(ctx context.Context, tx *sql.Tx, accountId int) (*entity.User, error) {
	ret := _m.Called(ctx, tx, accountId)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, int) (*entity.User, error)); ok {
		return rf(ctx, tx, accountId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, int) *entity.User); ok {
		r0 = rf(ctx, tx, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sql.Tx, int) error); ok {
		r1 = rf(ctx, tx, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserDetail provides a mock function with given fields: ctx, tx, accountId
func (_m *InterfaceUserRepository) GetUserDetail(ctx context.Context, tx *sql.Tx, accountId int) (*dto.UserDetailResponse, error) {
	ret := _m.Called(ctx, tx, accountId)

	var r0 *dto.UserDetailResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, int) (*dto.UserDetailResponse, error)); ok {
		return rf(ctx, tx, accountId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, int) *dto.UserDetailResponse); ok {
		r0 = rf(ctx, tx, accountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserDetailResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sql.Tx, int) error); ok {
		r1 = rf(ctx, tx, accountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertUser provides a mock function with given fields: ctx, tx, _a2
func (_m *InterfaceUserRepository) InsertUser(ctx context.Context, tx *sql.Tx, _a2 *entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, tx, _a2)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, *entity.User) (*entity.User, error)); ok {
		return rf(ctx, tx, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sql.Tx, *entity.User) *entity.User); ok {
		r0 = rf(ctx, tx, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sql.Tx, *entity.User) error); ok {
		r1 = rf(ctx, tx, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInterfaceUserRepository creates a new instance of InterfaceUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInterfaceUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *InterfaceUserRepository {
	mock := &InterfaceUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
