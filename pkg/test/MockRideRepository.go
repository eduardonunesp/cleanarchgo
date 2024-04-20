// Code generated by mockery v2.42.2. DO NOT EDIT.

package test

import (
	domain "github.com/eduardonunesp/cleanarchgo/pkg/domain"
	mock "github.com/stretchr/testify/mock"
)

// MockRideRepository is an autogenerated mock type for the RideRepository type
type MockRideRepository struct {
	mock.Mock
}

type MockRideRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRideRepository) EXPECT() *MockRideRepository_Expecter {
	return &MockRideRepository_Expecter{mock: &_m.Mock}
}

// GetRideByID provides a mock function with given fields: rideID
func (_m *MockRideRepository) GetRideByID(rideID string) (*domain.Ride, error) {
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

// MockRideRepository_GetRideByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRideByID'
type MockRideRepository_GetRideByID_Call struct {
	*mock.Call
}

// GetRideByID is a helper method to define mock.On call
//   - rideID string
func (_e *MockRideRepository_Expecter) GetRideByID(rideID interface{}) *MockRideRepository_GetRideByID_Call {
	return &MockRideRepository_GetRideByID_Call{Call: _e.mock.On("GetRideByID", rideID)}
}

func (_c *MockRideRepository_GetRideByID_Call) Run(run func(rideID string)) *MockRideRepository_GetRideByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRideRepository_GetRideByID_Call) Return(_a0 *domain.Ride, _a1 error) *MockRideRepository_GetRideByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRideRepository_GetRideByID_Call) RunAndReturn(run func(string) (*domain.Ride, error)) *MockRideRepository_GetRideByID_Call {
	_c.Call.Return(run)
	return _c
}

// HasActiveRideByPassengerID provides a mock function with given fields: passengerID
func (_m *MockRideRepository) HasActiveRideByPassengerID(passengerID string) bool {
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

// MockRideRepository_HasActiveRideByPassengerID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasActiveRideByPassengerID'
type MockRideRepository_HasActiveRideByPassengerID_Call struct {
	*mock.Call
}

// HasActiveRideByPassengerID is a helper method to define mock.On call
//   - passengerID string
func (_e *MockRideRepository_Expecter) HasActiveRideByPassengerID(passengerID interface{}) *MockRideRepository_HasActiveRideByPassengerID_Call {
	return &MockRideRepository_HasActiveRideByPassengerID_Call{Call: _e.mock.On("HasActiveRideByPassengerID", passengerID)}
}

func (_c *MockRideRepository_HasActiveRideByPassengerID_Call) Run(run func(passengerID string)) *MockRideRepository_HasActiveRideByPassengerID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRideRepository_HasActiveRideByPassengerID_Call) Return(_a0 bool) *MockRideRepository_HasActiveRideByPassengerID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRideRepository_HasActiveRideByPassengerID_Call) RunAndReturn(run func(string) bool) *MockRideRepository_HasActiveRideByPassengerID_Call {
	_c.Call.Return(run)
	return _c
}

// SaveRide provides a mock function with given fields: ride
func (_m *MockRideRepository) SaveRide(ride domain.Ride) {
	_m.Called(ride)
}

// MockRideRepository_SaveRide_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveRide'
type MockRideRepository_SaveRide_Call struct {
	*mock.Call
}

// SaveRide is a helper method to define mock.On call
//   - ride domain.Ride
func (_e *MockRideRepository_Expecter) SaveRide(ride interface{}) *MockRideRepository_SaveRide_Call {
	return &MockRideRepository_SaveRide_Call{Call: _e.mock.On("SaveRide", ride)}
}

func (_c *MockRideRepository_SaveRide_Call) Run(run func(ride domain.Ride)) *MockRideRepository_SaveRide_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.Ride))
	})
	return _c
}

func (_c *MockRideRepository_SaveRide_Call) Return() *MockRideRepository_SaveRide_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockRideRepository_SaveRide_Call) RunAndReturn(run func(domain.Ride)) *MockRideRepository_SaveRide_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRideRepository creates a new instance of MockRideRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRideRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRideRepository {
	mock := &MockRideRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}