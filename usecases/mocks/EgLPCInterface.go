// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	api "github.com/enbility/eebus-go/usecases/api"
	mock "github.com/stretchr/testify/mock"

	model "github.com/enbility/spine-go/model"

	spine_goapi "github.com/enbility/spine-go/api"

	time "time"
)

// EgLPCInterface is an autogenerated mock type for the EgLPCInterface type
type EgLPCInterface struct {
	mock.Mock
}

type EgLPCInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *EgLPCInterface) EXPECT() *EgLPCInterface_Expecter {
	return &EgLPCInterface_Expecter{mock: &_m.Mock}
}

// AddFeatures provides a mock function with given fields:
func (_m *EgLPCInterface) AddFeatures() {
	_m.Called()
}

// EgLPCInterface_AddFeatures_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddFeatures'
type EgLPCInterface_AddFeatures_Call struct {
	*mock.Call
}

// AddFeatures is a helper method to define mock.On call
func (_e *EgLPCInterface_Expecter) AddFeatures() *EgLPCInterface_AddFeatures_Call {
	return &EgLPCInterface_AddFeatures_Call{Call: _e.mock.On("AddFeatures")}
}

func (_c *EgLPCInterface_AddFeatures_Call) Run(run func()) *EgLPCInterface_AddFeatures_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPCInterface_AddFeatures_Call) Return() *EgLPCInterface_AddFeatures_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPCInterface_AddFeatures_Call) RunAndReturn(run func()) *EgLPCInterface_AddFeatures_Call {
	_c.Call.Return(run)
	return _c
}

// AddUseCase provides a mock function with given fields:
func (_m *EgLPCInterface) AddUseCase() {
	_m.Called()
}

// EgLPCInterface_AddUseCase_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUseCase'
type EgLPCInterface_AddUseCase_Call struct {
	*mock.Call
}

// AddUseCase is a helper method to define mock.On call
func (_e *EgLPCInterface_Expecter) AddUseCase() *EgLPCInterface_AddUseCase_Call {
	return &EgLPCInterface_AddUseCase_Call{Call: _e.mock.On("AddUseCase")}
}

func (_c *EgLPCInterface_AddUseCase_Call) Run(run func()) *EgLPCInterface_AddUseCase_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EgLPCInterface_AddUseCase_Call) Return() *EgLPCInterface_AddUseCase_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPCInterface_AddUseCase_Call) RunAndReturn(run func()) *EgLPCInterface_AddUseCase_Call {
	_c.Call.Return(run)
	return _c
}

// ConsumptionLimit provides a mock function with given fields: entity
func (_m *EgLPCInterface) ConsumptionLimit(entity spine_goapi.EntityRemoteInterface) (api.LoadLimit, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for ConsumptionLimit")
	}

	var r0 api.LoadLimit
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (api.LoadLimit, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) api.LoadLimit); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(api.LoadLimit)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_ConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsumptionLimit'
type EgLPCInterface_ConsumptionLimit_Call struct {
	*mock.Call
}

// ConsumptionLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPCInterface_Expecter) ConsumptionLimit(entity interface{}) *EgLPCInterface_ConsumptionLimit_Call {
	return &EgLPCInterface_ConsumptionLimit_Call{Call: _e.mock.On("ConsumptionLimit", entity)}
}

func (_c *EgLPCInterface_ConsumptionLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPCInterface_ConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPCInterface_ConsumptionLimit_Call) Return(limit api.LoadLimit, resultErr error) *EgLPCInterface_ConsumptionLimit_Call {
	_c.Call.Return(limit, resultErr)
	return _c
}

func (_c *EgLPCInterface_ConsumptionLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (api.LoadLimit, error)) *EgLPCInterface_ConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeConsumptionActivePowerLimit provides a mock function with given fields: entity
func (_m *EgLPCInterface) FailsafeConsumptionActivePowerLimit(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for FailsafeConsumptionActivePowerLimit")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeConsumptionActivePowerLimit'
type EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call struct {
	*mock.Call
}

// FailsafeConsumptionActivePowerLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPCInterface_Expecter) FailsafeConsumptionActivePowerLimit(entity interface{}) *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	return &EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call{Call: _e.mock.On("FailsafeConsumptionActivePowerLimit", entity)}
}

func (_c *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call) Return(_a0 float64, _a1 error) *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *EgLPCInterface_FailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// FailsafeDurationMinimum provides a mock function with given fields: entity
func (_m *EgLPCInterface) FailsafeDurationMinimum(entity spine_goapi.EntityRemoteInterface) (time.Duration, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for FailsafeDurationMinimum")
	}

	var r0 time.Duration
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (time.Duration, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) time.Duration); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_FailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FailsafeDurationMinimum'
type EgLPCInterface_FailsafeDurationMinimum_Call struct {
	*mock.Call
}

// FailsafeDurationMinimum is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPCInterface_Expecter) FailsafeDurationMinimum(entity interface{}) *EgLPCInterface_FailsafeDurationMinimum_Call {
	return &EgLPCInterface_FailsafeDurationMinimum_Call{Call: _e.mock.On("FailsafeDurationMinimum", entity)}
}

func (_c *EgLPCInterface_FailsafeDurationMinimum_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPCInterface_FailsafeDurationMinimum_Call) Return(_a0 time.Duration, _a1 error) *EgLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_FailsafeDurationMinimum_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (time.Duration, error)) *EgLPCInterface_FailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// IsCompatibleEntity provides a mock function with given fields: entity
func (_m *EgLPCInterface) IsCompatibleEntity(entity spine_goapi.EntityRemoteInterface) bool {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for IsCompatibleEntity")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EgLPCInterface_IsCompatibleEntity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsCompatibleEntity'
type EgLPCInterface_IsCompatibleEntity_Call struct {
	*mock.Call
}

// IsCompatibleEntity is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPCInterface_Expecter) IsCompatibleEntity(entity interface{}) *EgLPCInterface_IsCompatibleEntity_Call {
	return &EgLPCInterface_IsCompatibleEntity_Call{Call: _e.mock.On("IsCompatibleEntity", entity)}
}

func (_c *EgLPCInterface_IsCompatibleEntity_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPCInterface_IsCompatibleEntity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPCInterface_IsCompatibleEntity_Call) Return(_a0 bool) *EgLPCInterface_IsCompatibleEntity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EgLPCInterface_IsCompatibleEntity_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) bool) *EgLPCInterface_IsCompatibleEntity_Call {
	_c.Call.Return(run)
	return _c
}

// IsUseCaseSupported provides a mock function with given fields: remoteEntity
func (_m *EgLPCInterface) IsUseCaseSupported(remoteEntity spine_goapi.EntityRemoteInterface) (bool, error) {
	ret := _m.Called(remoteEntity)

	if len(ret) == 0 {
		panic("no return value specified for IsUseCaseSupported")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (bool, error)); ok {
		return rf(remoteEntity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) bool); ok {
		r0 = rf(remoteEntity)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(remoteEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_IsUseCaseSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsUseCaseSupported'
type EgLPCInterface_IsUseCaseSupported_Call struct {
	*mock.Call
}

// IsUseCaseSupported is a helper method to define mock.On call
//   - remoteEntity spine_goapi.EntityRemoteInterface
func (_e *EgLPCInterface_Expecter) IsUseCaseSupported(remoteEntity interface{}) *EgLPCInterface_IsUseCaseSupported_Call {
	return &EgLPCInterface_IsUseCaseSupported_Call{Call: _e.mock.On("IsUseCaseSupported", remoteEntity)}
}

func (_c *EgLPCInterface_IsUseCaseSupported_Call) Run(run func(remoteEntity spine_goapi.EntityRemoteInterface)) *EgLPCInterface_IsUseCaseSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPCInterface_IsUseCaseSupported_Call) Return(_a0 bool, _a1 error) *EgLPCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_IsUseCaseSupported_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (bool, error)) *EgLPCInterface_IsUseCaseSupported_Call {
	_c.Call.Return(run)
	return _c
}

// PowerConsumptionNominalMax provides a mock function with given fields: entity
func (_m *EgLPCInterface) PowerConsumptionNominalMax(entity spine_goapi.EntityRemoteInterface) (float64, error) {
	ret := _m.Called(entity)

	if len(ret) == 0 {
		panic("no return value specified for PowerConsumptionNominalMax")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) (float64, error)); ok {
		return rf(entity)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface) float64); ok {
		r0 = rf(entity)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_PowerConsumptionNominalMax_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PowerConsumptionNominalMax'
type EgLPCInterface_PowerConsumptionNominalMax_Call struct {
	*mock.Call
}

// PowerConsumptionNominalMax is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
func (_e *EgLPCInterface_Expecter) PowerConsumptionNominalMax(entity interface{}) *EgLPCInterface_PowerConsumptionNominalMax_Call {
	return &EgLPCInterface_PowerConsumptionNominalMax_Call{Call: _e.mock.On("PowerConsumptionNominalMax", entity)}
}

func (_c *EgLPCInterface_PowerConsumptionNominalMax_Call) Run(run func(entity spine_goapi.EntityRemoteInterface)) *EgLPCInterface_PowerConsumptionNominalMax_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface))
	})
	return _c
}

func (_c *EgLPCInterface_PowerConsumptionNominalMax_Call) Return(_a0 float64, _a1 error) *EgLPCInterface_PowerConsumptionNominalMax_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_PowerConsumptionNominalMax_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface) (float64, error)) *EgLPCInterface_PowerConsumptionNominalMax_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUseCaseAvailability provides a mock function with given fields: available
func (_m *EgLPCInterface) UpdateUseCaseAvailability(available bool) {
	_m.Called(available)
}

// EgLPCInterface_UpdateUseCaseAvailability_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUseCaseAvailability'
type EgLPCInterface_UpdateUseCaseAvailability_Call struct {
	*mock.Call
}

// UpdateUseCaseAvailability is a helper method to define mock.On call
//   - available bool
func (_e *EgLPCInterface_Expecter) UpdateUseCaseAvailability(available interface{}) *EgLPCInterface_UpdateUseCaseAvailability_Call {
	return &EgLPCInterface_UpdateUseCaseAvailability_Call{Call: _e.mock.On("UpdateUseCaseAvailability", available)}
}

func (_c *EgLPCInterface_UpdateUseCaseAvailability_Call) Run(run func(available bool)) *EgLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *EgLPCInterface_UpdateUseCaseAvailability_Call) Return() *EgLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return()
	return _c
}

func (_c *EgLPCInterface_UpdateUseCaseAvailability_Call) RunAndReturn(run func(bool)) *EgLPCInterface_UpdateUseCaseAvailability_Call {
	_c.Call.Return(run)
	return _c
}

// WriteConsumptionLimit provides a mock function with given fields: entity, limit
func (_m *EgLPCInterface) WriteConsumptionLimit(entity spine_goapi.EntityRemoteInterface, limit api.LoadLimit) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, limit)

	if len(ret) == 0 {
		panic("no return value specified for WriteConsumptionLimit")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, api.LoadLimit) (*model.MsgCounterType, error)); ok {
		return rf(entity, limit)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, api.LoadLimit) *model.MsgCounterType); ok {
		r0 = rf(entity, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, api.LoadLimit) error); ok {
		r1 = rf(entity, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_WriteConsumptionLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteConsumptionLimit'
type EgLPCInterface_WriteConsumptionLimit_Call struct {
	*mock.Call
}

// WriteConsumptionLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - limit api.LoadLimit
func (_e *EgLPCInterface_Expecter) WriteConsumptionLimit(entity interface{}, limit interface{}) *EgLPCInterface_WriteConsumptionLimit_Call {
	return &EgLPCInterface_WriteConsumptionLimit_Call{Call: _e.mock.On("WriteConsumptionLimit", entity, limit)}
}

func (_c *EgLPCInterface_WriteConsumptionLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, limit api.LoadLimit)) *EgLPCInterface_WriteConsumptionLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(api.LoadLimit))
	})
	return _c
}

func (_c *EgLPCInterface_WriteConsumptionLimit_Call) Return(_a0 *model.MsgCounterType, _a1 error) *EgLPCInterface_WriteConsumptionLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_WriteConsumptionLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, api.LoadLimit) (*model.MsgCounterType, error)) *EgLPCInterface_WriteConsumptionLimit_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFailsafeConsumptionActivePowerLimit provides a mock function with given fields: entity, value
func (_m *EgLPCInterface) WriteFailsafeConsumptionActivePowerLimit(entity spine_goapi.EntityRemoteInterface, value float64) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, value)

	if len(ret) == 0 {
		panic("no return value specified for WriteFailsafeConsumptionActivePowerLimit")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, float64) (*model.MsgCounterType, error)); ok {
		return rf(entity, value)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, float64) *model.MsgCounterType); ok {
		r0 = rf(entity, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, float64) error); ok {
		r1 = rf(entity, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFailsafeConsumptionActivePowerLimit'
type EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call struct {
	*mock.Call
}

// WriteFailsafeConsumptionActivePowerLimit is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - value float64
func (_e *EgLPCInterface_Expecter) WriteFailsafeConsumptionActivePowerLimit(entity interface{}, value interface{}) *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	return &EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call{Call: _e.mock.On("WriteFailsafeConsumptionActivePowerLimit", entity, value)}
}

func (_c *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, value float64)) *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(float64))
	})
	return _c
}

func (_c *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call) Return(_a0 *model.MsgCounterType, _a1 error) *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, float64) (*model.MsgCounterType, error)) *EgLPCInterface_WriteFailsafeConsumptionActivePowerLimit_Call {
	_c.Call.Return(run)
	return _c
}

// WriteFailsafeDurationMinimum provides a mock function with given fields: entity, duration
func (_m *EgLPCInterface) WriteFailsafeDurationMinimum(entity spine_goapi.EntityRemoteInterface, duration time.Duration) (*model.MsgCounterType, error) {
	ret := _m.Called(entity, duration)

	if len(ret) == 0 {
		panic("no return value specified for WriteFailsafeDurationMinimum")
	}

	var r0 *model.MsgCounterType
	var r1 error
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, time.Duration) (*model.MsgCounterType, error)); ok {
		return rf(entity, duration)
	}
	if rf, ok := ret.Get(0).(func(spine_goapi.EntityRemoteInterface, time.Duration) *model.MsgCounterType); ok {
		r0 = rf(entity, duration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MsgCounterType)
		}
	}

	if rf, ok := ret.Get(1).(func(spine_goapi.EntityRemoteInterface, time.Duration) error); ok {
		r1 = rf(entity, duration)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EgLPCInterface_WriteFailsafeDurationMinimum_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteFailsafeDurationMinimum'
type EgLPCInterface_WriteFailsafeDurationMinimum_Call struct {
	*mock.Call
}

// WriteFailsafeDurationMinimum is a helper method to define mock.On call
//   - entity spine_goapi.EntityRemoteInterface
//   - duration time.Duration
func (_e *EgLPCInterface_Expecter) WriteFailsafeDurationMinimum(entity interface{}, duration interface{}) *EgLPCInterface_WriteFailsafeDurationMinimum_Call {
	return &EgLPCInterface_WriteFailsafeDurationMinimum_Call{Call: _e.mock.On("WriteFailsafeDurationMinimum", entity, duration)}
}

func (_c *EgLPCInterface_WriteFailsafeDurationMinimum_Call) Run(run func(entity spine_goapi.EntityRemoteInterface, duration time.Duration)) *EgLPCInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(spine_goapi.EntityRemoteInterface), args[1].(time.Duration))
	})
	return _c
}

func (_c *EgLPCInterface_WriteFailsafeDurationMinimum_Call) Return(_a0 *model.MsgCounterType, _a1 error) *EgLPCInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *EgLPCInterface_WriteFailsafeDurationMinimum_Call) RunAndReturn(run func(spine_goapi.EntityRemoteInterface, time.Duration) (*model.MsgCounterType, error)) *EgLPCInterface_WriteFailsafeDurationMinimum_Call {
	_c.Call.Return(run)
	return _c
}

// NewEgLPCInterface creates a new instance of EgLPCInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEgLPCInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *EgLPCInterface {
	mock := &EgLPCInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
