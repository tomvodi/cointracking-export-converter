// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Initializer is an autogenerated mock type for the Initializer type
type Initializer struct {
	mock.Mock
}

type Initializer_Expecter struct {
	mock *mock.Mock
}

func (_m *Initializer) EXPECT() *Initializer_Expecter {
	return &Initializer_Expecter{mock: &_m.Mock}
}

// Init provides a mock function with given fields:
func (_m *Initializer) Init() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Initializer_Init_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Init'
type Initializer_Init_Call struct {
	*mock.Call
}

// Init is a helper method to define mock.On call
func (_e *Initializer_Expecter) Init() *Initializer_Init_Call {
	return &Initializer_Init_Call{Call: _e.mock.On("Init")}
}

func (_c *Initializer_Init_Call) Run(run func()) *Initializer_Init_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Initializer_Init_Call) Return(_a0 error) *Initializer_Init_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Initializer_Init_Call) RunAndReturn(run func() error) *Initializer_Init_Call {
	_c.Call.Return(run)
	return _c
}

// NewInitializer creates a new instance of Initializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInitializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Initializer {
	mock := &Initializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
