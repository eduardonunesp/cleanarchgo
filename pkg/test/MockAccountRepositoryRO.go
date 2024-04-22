// Code generated by mockery v2.42.2. DO NOT EDIT.

package test

import (
	domain "github.com/eduardonunesp/cleanarchgo/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockAccountRepositoryRO is an autogenerated mock type for the AccountRepositoryRO type
type MockAccountRepositoryRO struct {
	mock.Mock
}

type MockAccountRepositoryRO_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountRepositoryRO) EXPECT() *MockAccountRepositoryRO_Expecter {
	return &MockAccountRepositoryRO_Expecter{mock: &_m.Mock}
}

// GetAccountByID provides a mock function with given fields: id
func (_m *MockAccountRepositoryRO) GetAccountByID(id string) (*domain.Account, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountByID")
	}

	var r0 *domain.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Account, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepositoryRO_GetAccountByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountByID'
type MockAccountRepositoryRO_GetAccountByID_Call struct {
	*mock.Call
}

// GetAccountByID is a helper method to define mock.On call
//   - id string
func (_e *MockAccountRepositoryRO_Expecter) GetAccountByID(id interface{}) *MockAccountRepositoryRO_GetAccountByID_Call {
	return &MockAccountRepositoryRO_GetAccountByID_Call{Call: _e.mock.On("GetAccountByID", id)}
}

func (_c *MockAccountRepositoryRO_GetAccountByID_Call) Run(run func(id string)) *MockAccountRepositoryRO_GetAccountByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAccountRepositoryRO_GetAccountByID_Call) Return(_a0 *domain.Account, _a1 error) *MockAccountRepositoryRO_GetAccountByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepositoryRO_GetAccountByID_Call) RunAndReturn(run func(string) (*domain.Account, error)) *MockAccountRepositoryRO_GetAccountByID_Call {
	_c.Call.Return(run)
	return _c
}

// HasAccountByEmail provides a mock function with given fields: email
func (_m *MockAccountRepositoryRO) HasAccountByEmail(email string) (bool, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for HasAccountByEmail")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepositoryRO_HasAccountByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasAccountByEmail'
type MockAccountRepositoryRO_HasAccountByEmail_Call struct {
	*mock.Call
}

// HasAccountByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockAccountRepositoryRO_Expecter) HasAccountByEmail(email interface{}) *MockAccountRepositoryRO_HasAccountByEmail_Call {
	return &MockAccountRepositoryRO_HasAccountByEmail_Call{Call: _e.mock.On("HasAccountByEmail", email)}
}

func (_c *MockAccountRepositoryRO_HasAccountByEmail_Call) Run(run func(email string)) *MockAccountRepositoryRO_HasAccountByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAccountRepositoryRO_HasAccountByEmail_Call) Return(_a0 bool, _a1 error) *MockAccountRepositoryRO_HasAccountByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepositoryRO_HasAccountByEmail_Call) RunAndReturn(run func(string) (bool, error)) *MockAccountRepositoryRO_HasAccountByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// IsDriverFreeByDriverID provides a mock function with given fields: driverID
func (_m *MockAccountRepositoryRO) IsDriverFreeByDriverID(driverID string) (bool, error) {
	ret := _m.Called(driverID)

	if len(ret) == 0 {
		panic("no return value specified for IsDriverFreeByDriverID")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(driverID)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(driverID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(driverID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountRepositoryRO_IsDriverFreeByDriverID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDriverFreeByDriverID'
type MockAccountRepositoryRO_IsDriverFreeByDriverID_Call struct {
	*mock.Call
}

// IsDriverFreeByDriverID is a helper method to define mock.On call
//   - driverID string
func (_e *MockAccountRepositoryRO_Expecter) IsDriverFreeByDriverID(driverID interface{}) *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call {
	return &MockAccountRepositoryRO_IsDriverFreeByDriverID_Call{Call: _e.mock.On("IsDriverFreeByDriverID", driverID)}
}

func (_c *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call) Run(run func(driverID string)) *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call) Return(_a0 bool, _a1 error) *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call) RunAndReturn(run func(string) (bool, error)) *MockAccountRepositoryRO_IsDriverFreeByDriverID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAccountRepositoryRO creates a new instance of MockAccountRepositoryRO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAccountRepositoryRO(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountRepositoryRO {
	mock := &MockAccountRepositoryRO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
