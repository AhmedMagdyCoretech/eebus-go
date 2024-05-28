// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	model "github.com/enbility/spine-go/model"
	mock "github.com/stretchr/testify/mock"
)

// IdentificationCommonInterface is an autogenerated mock type for the IdentificationCommonInterface type
type IdentificationCommonInterface struct {
	mock.Mock
}

type IdentificationCommonInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *IdentificationCommonInterface) EXPECT() *IdentificationCommonInterface_Expecter {
	return &IdentificationCommonInterface_Expecter{mock: &_m.Mock}
}

// CheckEventPayloadDataForFilter provides a mock function with given fields: payloadData
func (_m *IdentificationCommonInterface) CheckEventPayloadDataForFilter(payloadData interface{}) bool {
	ret := _m.Called(payloadData)

	if len(ret) == 0 {
		panic("no return value specified for CheckEventPayloadDataForFilter")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}) bool); ok {
		r0 = rf(payloadData)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckEventPayloadDataForFilter'
type IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call struct {
	*mock.Call
}

// CheckEventPayloadDataForFilter is a helper method to define mock.On call
//   - payloadData interface{}
func (_e *IdentificationCommonInterface_Expecter) CheckEventPayloadDataForFilter(payloadData interface{}) *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call {
	return &IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call{Call: _e.mock.On("CheckEventPayloadDataForFilter", payloadData)}
}

func (_c *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call) Run(run func(payloadData interface{})) *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call) Return(_a0 bool) *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call) RunAndReturn(run func(interface{}) bool) *IdentificationCommonInterface_CheckEventPayloadDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// GetDataForFilter provides a mock function with given fields: filter
func (_m *IdentificationCommonInterface) GetDataForFilter(filter model.IdentificationDataType) ([]model.IdentificationDataType, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for GetDataForFilter")
	}

	var r0 []model.IdentificationDataType
	var r1 error
	if rf, ok := ret.Get(0).(func(model.IdentificationDataType) ([]model.IdentificationDataType, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(model.IdentificationDataType) []model.IdentificationDataType); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.IdentificationDataType)
		}
	}

	if rf, ok := ret.Get(1).(func(model.IdentificationDataType) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IdentificationCommonInterface_GetDataForFilter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDataForFilter'
type IdentificationCommonInterface_GetDataForFilter_Call struct {
	*mock.Call
}

// GetDataForFilter is a helper method to define mock.On call
//   - filter model.IdentificationDataType
func (_e *IdentificationCommonInterface_Expecter) GetDataForFilter(filter interface{}) *IdentificationCommonInterface_GetDataForFilter_Call {
	return &IdentificationCommonInterface_GetDataForFilter_Call{Call: _e.mock.On("GetDataForFilter", filter)}
}

func (_c *IdentificationCommonInterface_GetDataForFilter_Call) Run(run func(filter model.IdentificationDataType)) *IdentificationCommonInterface_GetDataForFilter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(model.IdentificationDataType))
	})
	return _c
}

func (_c *IdentificationCommonInterface_GetDataForFilter_Call) Return(_a0 []model.IdentificationDataType, _a1 error) *IdentificationCommonInterface_GetDataForFilter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IdentificationCommonInterface_GetDataForFilter_Call) RunAndReturn(run func(model.IdentificationDataType) ([]model.IdentificationDataType, error)) *IdentificationCommonInterface_GetDataForFilter_Call {
	_c.Call.Return(run)
	return _c
}

// NewIdentificationCommonInterface creates a new instance of IdentificationCommonInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIdentificationCommonInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *IdentificationCommonInterface {
	mock := &IdentificationCommonInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
