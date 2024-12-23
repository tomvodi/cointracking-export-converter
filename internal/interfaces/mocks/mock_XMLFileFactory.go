// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	interfaces "github.com/tomvodi/cointracking-export-converter/internal/interfaces"
)

// XMLFileFactory is an autogenerated mock type for the XMLFileFactory type
type XMLFileFactory struct {
	mock.Mock
}

type XMLFileFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *XMLFileFactory) EXPECT() *XMLFileFactory_Expecter {
	return &XMLFileFactory_Expecter{mock: &_m.Mock}
}

// NewXMLFile provides a mock function with given fields:
func (_m *XMLFileFactory) NewXMLFile() interfaces.XMLFile {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NewXMLFile")
	}

	var r0 interfaces.XMLFile
	if rf, ok := ret.Get(0).(func() interfaces.XMLFile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.XMLFile)
		}
	}

	return r0
}

// XMLFileFactory_NewXMLFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewXMLFile'
type XMLFileFactory_NewXMLFile_Call struct {
	*mock.Call
}

// NewXMLFile is a helper method to define mock.On call
func (_e *XMLFileFactory_Expecter) NewXMLFile() *XMLFileFactory_NewXMLFile_Call {
	return &XMLFileFactory_NewXMLFile_Call{Call: _e.mock.On("NewXMLFile")}
}

func (_c *XMLFileFactory_NewXMLFile_Call) Run(run func()) *XMLFileFactory_NewXMLFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *XMLFileFactory_NewXMLFile_Call) Return(_a0 interfaces.XMLFile) *XMLFileFactory_NewXMLFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *XMLFileFactory_NewXMLFile_Call) RunAndReturn(run func() interfaces.XMLFile) *XMLFileFactory_NewXMLFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewXMLFileFactory creates a new instance of XMLFileFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewXMLFileFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *XMLFileFactory {
	mock := &XMLFileFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
