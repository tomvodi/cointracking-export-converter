// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	common "github.com/tomvodi/cointracking-export-converter/internal/common"

	mock "github.com/stretchr/testify/mock"
)

// TransactionsFileWriter is an autogenerated mock type for the TransactionsFileWriter type
type TransactionsFileWriter struct {
	mock.Mock
}

type TransactionsFileWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *TransactionsFileWriter) EXPECT() *TransactionsFileWriter_Expecter {
	return &TransactionsFileWriter_Expecter{mock: &_m.Mock}
}

// WriteTransactionsToFile provides a mock function with given fields: filePath, transactions
func (_m *TransactionsFileWriter) WriteTransactionsToFile(filePath string, transactions []*common.CointrackingTx) error {
	ret := _m.Called(filePath, transactions)

	if len(ret) == 0 {
		panic("no return value specified for WriteTransactionsToFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []*common.CointrackingTx) error); ok {
		r0 = rf(filePath, transactions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransactionsFileWriter_WriteTransactionsToFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteTransactionsToFile'
type TransactionsFileWriter_WriteTransactionsToFile_Call struct {
	*mock.Call
}

// WriteTransactionsToFile is a helper method to define mock.On call
//   - filePath string
//   - transactions []*common.CointrackingTx
func (_e *TransactionsFileWriter_Expecter) WriteTransactionsToFile(filePath interface{}, transactions interface{}) *TransactionsFileWriter_WriteTransactionsToFile_Call {
	return &TransactionsFileWriter_WriteTransactionsToFile_Call{Call: _e.mock.On("WriteTransactionsToFile", filePath, transactions)}
}

func (_c *TransactionsFileWriter_WriteTransactionsToFile_Call) Run(run func(filePath string, transactions []*common.CointrackingTx)) *TransactionsFileWriter_WriteTransactionsToFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]*common.CointrackingTx))
	})
	return _c
}

func (_c *TransactionsFileWriter_WriteTransactionsToFile_Call) Return(_a0 error) *TransactionsFileWriter_WriteTransactionsToFile_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TransactionsFileWriter_WriteTransactionsToFile_Call) RunAndReturn(run func(string, []*common.CointrackingTx) error) *TransactionsFileWriter_WriteTransactionsToFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransactionsFileWriter creates a new instance of TransactionsFileWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionsFileWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionsFileWriter {
	mock := &TransactionsFileWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
