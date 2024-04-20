// Code generated by mockery v2.42.2. DO NOT EDIT.

package test

import (
	domain "github.com/eduardonunesp/cleanarchgo/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockRideRepositoryRO is an autogenerated mock type for the RideRepositoryRO type
type MockRideRepositoryRO struct {
	mock.Mock
}

type MockRideRepositoryRO_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRideRepositoryRO) EXPECT() *MockRideRepositoryRO_Expecter {
	return &MockRideRepositoryRO_Expecter{mock: &_m.Mock}
}

// GetRideByID provides a mock function with given fields: rideID
func (_m *MockRideRepositoryRO) GetRideByID(rideID string) (*domain.Ride, error) {
	ret := _m.Called(rideID)

	if len(ret) == 0 {
		panic("no return value specified for GetRideByID")
	}

	var r0 *domain.Ride
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Ride, error)); ok {
		return rf(rideID)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Ride); ok {
		r0 = rf(rideID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Ride)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(rideID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRideRepositoryRO_GetRideByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRideByID'
type MockRideRepositoryRO_GetRideByID_Call struct {
	*mock.Call
}

// GetRideByID is a helper method to define mock.On call
//   - rideID string
func (_e *MockRideRepositoryRO_Expecter) GetRideByID(rideID interface{}) *MockRideRepositoryRO_GetRideByID_Call {
	return &MockRideRepositoryRO_GetRideByID_Call{Call: _e.mock.On("GetRideByID", rideID)}
}

func (_c *MockRideRepositoryRO_GetRideByID_Call) Run(run func(rideID string)) *MockRideRepositoryRO_GetRideByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRideRepositoryRO_GetRideByID_Call) Return(_a0 *domain.Ride, _a1 error) *MockRideRepositoryRO_GetRideByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRideRepositoryRO_GetRideByID_Call) RunAndReturn(run func(string) (*domain.Ride, error)) *MockRideRepositoryRO_GetRideByID_Call {
	_c.Call.Return(run)
	return _c
}

// HasActiveRideByPassengerID provides a mock function with given fields: passengerID
func (_m *MockRideRepositoryRO) HasActiveRideByPassengerID(passengerID string) bool {
	ret := _m.Called(passengerID)

	if len(ret) == 0 {
		panic("no return value specified for HasActiveRideByPassengerID")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(passengerID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockRideRepositoryRO_HasActiveRideByPassengerID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasActiveRideByPassengerID'
type MockRideRepositoryRO_HasActiveRideByPassengerID_Call struct {
	*mock.Call
}

// HasActiveRideByPassengerID is a helper method to define mock.On call
//   - passengerID string
func (_e *MockRideRepositoryRO_Expecter) HasActiveRideByPassengerID(passengerID interface{}) *MockRideRepositoryRO_HasActiveRideByPassengerID_Call {
	return &MockRideRepositoryRO_HasActiveRideByPassengerID_Call{Call: _e.mock.On("HasActiveRideByPassengerID", passengerID)}
}

func (_c *MockRideRepositoryRO_HasActiveRideByPassengerID_Call) Run(run func(passengerID string)) *MockRideRepositoryRO_HasActiveRideByPassengerID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRideRepositoryRO_HasActiveRideByPassengerID_Call) Return(_a0 bool) *MockRideRepositoryRO_HasActiveRideByPassengerID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRideRepositoryRO_HasActiveRideByPassengerID_Call) RunAndReturn(run func(string) bool) *MockRideRepositoryRO_HasActiveRideByPassengerID_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRideRepositoryRO creates a new instance of MockRideRepositoryRO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRideRepositoryRO(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRideRepositoryRO {
	mock := &MockRideRepositoryRO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
