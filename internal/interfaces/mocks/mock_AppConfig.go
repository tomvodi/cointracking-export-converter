// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	common "github.com/tomvodi/cointracking-export-converter/internal/common"

	mock "github.com/stretchr/testify/mock"
)

// AppConfig is an autogenerated mock type for the AppConfig type
type AppConfig struct {
	mock.Mock
}

type AppConfig_Expecter struct {
	mock *mock.Mock
}

func (_m *AppConfig) EXPECT() *AppConfig_Expecter {
	return &AppConfig_Expecter{mock: &_m.Mock}
}

// BlockpitTxTypes provides a mock function with given fields:
func (_m *AppConfig) BlockpitTxTypes() ([]common.TxDisplayName, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BlockpitTxTypes")
	}

	var r0 []common.TxDisplayName
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]common.TxDisplayName, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []common.TxDisplayName); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.TxDisplayName)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppConfig_BlockpitTxTypes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockpitTxTypes'
type AppConfig_BlockpitTxTypes_Call struct {
	*mock.Call
}

// BlockpitTxTypes is a helper method to define mock.On call
func (_e *AppConfig_Expecter) BlockpitTxTypes() *AppConfig_BlockpitTxTypes_Call {
	return &AppConfig_BlockpitTxTypes_Call{Call: _e.mock.On("BlockpitTxTypes")}
}

func (_c *AppConfig_BlockpitTxTypes_Call) Run(run func()) *AppConfig_BlockpitTxTypes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppConfig_BlockpitTxTypes_Call) Return(_a0 []common.TxDisplayName, _a1 error) *AppConfig_BlockpitTxTypes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppConfig_BlockpitTxTypes_Call) RunAndReturn(run func() ([]common.TxDisplayName, error)) *AppConfig_BlockpitTxTypes_Call {
	_c.Call.Return(run)
	return _c
}

// SetCointracking2BlockpitMapping provides a mock function with given fields: ctTxType, bpTxType
func (_m *AppConfig) SetCointracking2BlockpitMapping(ctTxType string, bpTxType string) error {
	ret := _m.Called(ctTxType, bpTxType)

	if len(ret) == 0 {
		panic("no return value specified for SetCointracking2BlockpitMapping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(ctTxType, bpTxType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppConfig_SetCointracking2BlockpitMapping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetCointracking2BlockpitMapping'
type AppConfig_SetCointracking2BlockpitMapping_Call struct {
	*mock.Call
}

// SetCointracking2BlockpitMapping is a helper method to define mock.On call
//   - ctTxType string
//   - bpTxType string
func (_e *AppConfig_Expecter) SetCointracking2BlockpitMapping(ctTxType interface{}, bpTxType interface{}) *AppConfig_SetCointracking2BlockpitMapping_Call {
	return &AppConfig_SetCointracking2BlockpitMapping_Call{Call: _e.mock.On("SetCointracking2BlockpitMapping", ctTxType, bpTxType)}
}

func (_c *AppConfig_SetCointracking2BlockpitMapping_Call) Run(run func(ctTxType string, bpTxType string)) *AppConfig_SetCointracking2BlockpitMapping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *AppConfig_SetCointracking2BlockpitMapping_Call) Return(_a0 error) *AppConfig_SetCointracking2BlockpitMapping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppConfig_SetCointracking2BlockpitMapping_Call) RunAndReturn(run func(string, string) error) *AppConfig_SetCointracking2BlockpitMapping_Call {
	_c.Call.Return(run)
	return _c
}

// SetSwapHandling provides a mock function with given fields: handling
func (_m *AppConfig) SetSwapHandling(handling string) error {
	ret := _m.Called(handling)

	if len(ret) == 0 {
		panic("no return value specified for SetSwapHandling")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(handling)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppConfig_SetSwapHandling_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSwapHandling'
type AppConfig_SetSwapHandling_Call struct {
	*mock.Call
}

// SetSwapHandling is a helper method to define mock.On call
//   - handling string
func (_e *AppConfig_Expecter) SetSwapHandling(handling interface{}) *AppConfig_SetSwapHandling_Call {
	return &AppConfig_SetSwapHandling_Call{Call: _e.mock.On("SetSwapHandling", handling)}
}

func (_c *AppConfig_SetSwapHandling_Call) Run(run func(handling string)) *AppConfig_SetSwapHandling_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppConfig_SetSwapHandling_Call) Return(_a0 error) *AppConfig_SetSwapHandling_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppConfig_SetSwapHandling_Call) RunAndReturn(run func(string) error) *AppConfig_SetSwapHandling_Call {
	_c.Call.Return(run)
	return _c
}

// SetTimezone provides a mock function with given fields: tz
func (_m *AppConfig) SetTimezone(tz string) error {
	ret := _m.Called(tz)

	if len(ret) == 0 {
		panic("no return value specified for SetTimezone")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tz)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AppConfig_SetTimezone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetTimezone'
type AppConfig_SetTimezone_Call struct {
	*mock.Call
}

// SetTimezone is a helper method to define mock.On call
//   - tz string
func (_e *AppConfig_Expecter) SetTimezone(tz interface{}) *AppConfig_SetTimezone_Call {
	return &AppConfig_SetTimezone_Call{Call: _e.mock.On("SetTimezone", tz)}
}

func (_c *AppConfig_SetTimezone_Call) Run(run func(tz string)) *AppConfig_SetTimezone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AppConfig_SetTimezone_Call) Return(_a0 error) *AppConfig_SetTimezone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppConfig_SetTimezone_Call) RunAndReturn(run func(string) error) *AppConfig_SetTimezone_Call {
	_c.Call.Return(run)
	return _c
}

// SwapHandling provides a mock function with given fields:
func (_m *AppConfig) SwapHandling() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SwapHandling")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AppConfig_SwapHandling_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SwapHandling'
type AppConfig_SwapHandling_Call struct {
	*mock.Call
}

// SwapHandling is a helper method to define mock.On call
func (_e *AppConfig_Expecter) SwapHandling() *AppConfig_SwapHandling_Call {
	return &AppConfig_SwapHandling_Call{Call: _e.mock.On("SwapHandling")}
}

func (_c *AppConfig_SwapHandling_Call) Run(run func()) *AppConfig_SwapHandling_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppConfig_SwapHandling_Call) Return(_a0 string) *AppConfig_SwapHandling_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppConfig_SwapHandling_Call) RunAndReturn(run func() string) *AppConfig_SwapHandling_Call {
	_c.Call.Return(run)
	return _c
}

// Timezone provides a mock function with given fields:
func (_m *AppConfig) Timezone() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Timezone")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AppConfig_Timezone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Timezone'
type AppConfig_Timezone_Call struct {
	*mock.Call
}

// Timezone is a helper method to define mock.On call
func (_e *AppConfig_Expecter) Timezone() *AppConfig_Timezone_Call {
	return &AppConfig_Timezone_Call{Call: _e.mock.On("Timezone")}
}

func (_c *AppConfig_Timezone_Call) Run(run func()) *AppConfig_Timezone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppConfig_Timezone_Call) Return(_a0 string) *AppConfig_Timezone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AppConfig_Timezone_Call) RunAndReturn(run func() string) *AppConfig_Timezone_Call {
	_c.Call.Return(run)
	return _c
}

// TxTypeMappings provides a mock function with given fields:
func (_m *AppConfig) TxTypeMappings() ([]common.Ct2BpTxMapping, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for TxTypeMappings")
	}

	var r0 []common.Ct2BpTxMapping
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]common.Ct2BpTxMapping, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []common.Ct2BpTxMapping); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]common.Ct2BpTxMapping)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AppConfig_TxTypeMappings_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxTypeMappings'
type AppConfig_TxTypeMappings_Call struct {
	*mock.Call
}

// TxTypeMappings is a helper method to define mock.On call
func (_e *AppConfig_Expecter) TxTypeMappings() *AppConfig_TxTypeMappings_Call {
	return &AppConfig_TxTypeMappings_Call{Call: _e.mock.On("TxTypeMappings")}
}

func (_c *AppConfig_TxTypeMappings_Call) Run(run func()) *AppConfig_TxTypeMappings_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *AppConfig_TxTypeMappings_Call) Return(_a0 []common.Ct2BpTxMapping, _a1 error) *AppConfig_TxTypeMappings_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AppConfig_TxTypeMappings_Call) RunAndReturn(run func() ([]common.Ct2BpTxMapping, error)) *AppConfig_TxTypeMappings_Call {
	_c.Call.Return(run)
	return _c
}

// NewAppConfig creates a new instance of AppConfig. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppConfig(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppConfig {
	mock := &AppConfig{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
