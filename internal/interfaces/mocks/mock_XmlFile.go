// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// XmlFile is an autogenerated mock type for the XmlFile type
type XmlFile struct {
	mock.Mock
}

type XmlFile_Expecter struct {
	mock *mock.Mock
}

func (_m *XmlFile) EXPECT() *XmlFile_Expecter {
	return &XmlFile_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *XmlFile) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// XmlFile_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type XmlFile_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *XmlFile_Expecter) Close() *XmlFile_Close_Call {
	return &XmlFile_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *XmlFile_Close_Call) Run(run func()) *XmlFile_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *XmlFile_Close_Call) Return(_a0 error) *XmlFile_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *XmlFile_Close_Call) RunAndReturn(run func() error) *XmlFile_Close_Call {
	_c.Call.Return(run)
	return _c
}

// SaveAs provides a mock function with given fields: filePath
func (_m *XmlFile) SaveAs(filePath string) error {
	ret := _m.Called(filePath)

	if len(ret) == 0 {
		panic("no return value specified for SaveAs")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// XmlFile_SaveAs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveAs'
type XmlFile_SaveAs_Call struct {
	*mock.Call
}

// SaveAs is a helper method to define mock.On call
//   - filePath string
func (_e *XmlFile_Expecter) SaveAs(filePath interface{}) *XmlFile_SaveAs_Call {
	return &XmlFile_SaveAs_Call{Call: _e.mock.On("SaveAs", filePath)}
}

func (_c *XmlFile_SaveAs_Call) Run(run func(filePath string)) *XmlFile_SaveAs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *XmlFile_SaveAs_Call) Return(_a0 error) *XmlFile_SaveAs_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *XmlFile_SaveAs_Call) RunAndReturn(run func(string) error) *XmlFile_SaveAs_Call {
	_c.Call.Return(run)
	return _c
}

// SetSheetHeader provides a mock function with given fields: sheetNr, headers
func (_m *XmlFile) SetSheetHeader(sheetNr int, headers []string) error {
	ret := _m.Called(sheetNr, headers)

	if len(ret) == 0 {
		panic("no return value specified for SetSheetHeader")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []string) error); ok {
		r0 = rf(sheetNr, headers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// XmlFile_SetSheetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSheetHeader'
type XmlFile_SetSheetHeader_Call struct {
	*mock.Call
}

// SetSheetHeader is a helper method to define mock.On call
//   - sheetNr int
//   - headers []string
func (_e *XmlFile_Expecter) SetSheetHeader(sheetNr interface{}, headers interface{}) *XmlFile_SetSheetHeader_Call {
	return &XmlFile_SetSheetHeader_Call{Call: _e.mock.On("SetSheetHeader", sheetNr, headers)}
}

func (_c *XmlFile_SetSheetHeader_Call) Run(run func(sheetNr int, headers []string)) *XmlFile_SetSheetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].([]string))
	})
	return _c
}

func (_c *XmlFile_SetSheetHeader_Call) Return(_a0 error) *XmlFile_SetSheetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *XmlFile_SetSheetHeader_Call) RunAndReturn(run func(int, []string) error) *XmlFile_SetSheetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// SetSheetRow provides a mock function with given fields: sheetNr, rowNr, data
func (_m *XmlFile) SetSheetRow(sheetNr int, rowNr int, data []interface{}) error {
	ret := _m.Called(sheetNr, rowNr, data)

	if len(ret) == 0 {
		panic("no return value specified for SetSheetRow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, []interface{}) error); ok {
		r0 = rf(sheetNr, rowNr, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// XmlFile_SetSheetRow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetSheetRow'
type XmlFile_SetSheetRow_Call struct {
	*mock.Call
}

// SetSheetRow is a helper method to define mock.On call
//   - sheetNr int
//   - rowNr int
//   - data []interface{}
func (_e *XmlFile_Expecter) SetSheetRow(sheetNr interface{}, rowNr interface{}, data interface{}) *XmlFile_SetSheetRow_Call {
	return &XmlFile_SetSheetRow_Call{Call: _e.mock.On("SetSheetRow", sheetNr, rowNr, data)}
}

func (_c *XmlFile_SetSheetRow_Call) Run(run func(sheetNr int, rowNr int, data []interface{})) *XmlFile_SetSheetRow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].([]interface{}))
	})
	return _c
}

func (_c *XmlFile_SetSheetRow_Call) Return(_a0 error) *XmlFile_SetSheetRow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *XmlFile_SetSheetRow_Call) RunAndReturn(run func(int, int, []interface{}) error) *XmlFile_SetSheetRow_Call {
	_c.Call.Return(run)
	return _c
}

// NewXmlFile creates a new instance of XmlFile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewXmlFile(t interface {
	mock.TestingT
	Cleanup(func())
}) *XmlFile {
	mock := &XmlFile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}