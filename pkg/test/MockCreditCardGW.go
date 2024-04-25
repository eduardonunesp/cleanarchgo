// Code generated by mockery v2.42.2. DO NOT EDIT.

package test

import mock "github.com/stretchr/testify/mock"

// MockCreditCardGW is an autogenerated mock type for the CreditCardGW type
type MockCreditCardGW struct {
	mock.Mock
}

type MockCreditCardGW_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreditCardGW) EXPECT() *MockCreditCardGW_Expecter {
	return &MockCreditCardGW_Expecter{mock: &_m.Mock}
}

// ProcessPayment provides a mock function with given fields: token, amount
func (_m *MockCreditCardGW) ProcessPayment(token string, amount string) error {
	ret := _m.Called(token, amount)

	if len(ret) == 0 {
		panic("no return value specified for ProcessPayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(token, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCreditCardGW_ProcessPayment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessPayment'
type MockCreditCardGW_ProcessPayment_Call struct {
	*mock.Call
}

// ProcessPayment is a helper method to define mock.On call
//   - token string
//   - amount string
func (_e *MockCreditCardGW_Expecter) ProcessPayment(token interface{}, amount interface{}) *MockCreditCardGW_ProcessPayment_Call {
	return &MockCreditCardGW_ProcessPayment_Call{Call: _e.mock.On("ProcessPayment", token, amount)}
}

func (_c *MockCreditCardGW_ProcessPayment_Call) Run(run func(token string, amount string)) *MockCreditCardGW_ProcessPayment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockCreditCardGW_ProcessPayment_Call) Return(_a0 error) *MockCreditCardGW_ProcessPayment_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCreditCardGW_ProcessPayment_Call) RunAndReturn(run func(string, string) error) *MockCreditCardGW_ProcessPayment_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreditCardGW creates a new instance of MockCreditCardGW. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreditCardGW(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreditCardGW {
	mock := &MockCreditCardGW{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}