package lpc

import (
	"errors"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/features/client"
	ucapi "github.com/enbility/eebus-go/usecases/api"
	internal "github.com/enbility/eebus-go/usecases/internal"
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/util"
)

// Scenario 1

// return the current loadcontrol limit data
//
// parameters:
//   - entity: the entity of the e.g. EVSE
//
// return values:
//   - limit: load limit data
//
// possible errors:
//   - ErrDataNotAvailable if no such limit is (yet) available
//   - and others
func (e *LPC) ConsumptionLimit(entity spineapi.EntityRemoteInterface) (
	limit ucapi.LoadLimit, resultErr error) {
	limit = ucapi.LoadLimit{
		Value:        0.0,
		IsChangeable: false,
		IsActive:     false,
	}

	resultErr = api.ErrNoCompatibleEntity
	if !e.IsCompatibleEntityType(entity) {
		return
	}

	resultErr = api.ErrDataNotAvailable
	loadControl, err := client.NewLoadControl(e.LocalEntity, entity)
	if err != nil || loadControl == nil {
		return
	}

	filter := model.LoadControlLimitDescriptionDataType{
		LimitType:      util.Ptr(model.LoadControlLimitTypeTypeSignDependentAbsValueLimit),
		LimitCategory:  util.Ptr(model.LoadControlCategoryTypeObligation),
		LimitDirection: util.Ptr(model.EnergyDirectionTypeConsume),
		ScopeType:      util.Ptr(model.ScopeTypeTypeActivePowerLimit),
	}
	limitDescriptions, err := loadControl.GetLimitDescriptionsForFilter(filter)
	if err != nil || len(limitDescriptions) != 1 {
		return
	}

	value, err := loadControl.GetLimitDataForId(*limitDescriptions[0].LimitId)
	if err != nil || value.Value == nil {
		return
	}

	limit.Value = value.Value.GetValue()
	limit.IsChangeable = (value.IsLimitChangeable != nil && *value.IsLimitChangeable)
	limit.IsActive = (value.IsLimitActive != nil && *value.IsLimitActive)
	if value.TimePeriod != nil && value.TimePeriod.EndTime != nil {
		if duration, err := value.TimePeriod.GetDuration(); err == nil {
			limit.Duration = duration
		}
	}

	resultErr = nil

	return
}

// send new LoadControlLimits
//
// parameters:
//   - entity: the entity of the e.g. EVSE
//   - limit: load limit data
//   - resultCB: callback function for handling the result response
func (e *LPC) WriteConsumptionLimit(
	entity spineapi.EntityRemoteInterface,
	limit ucapi.LoadLimit,
	resultCB func(result model.ResultDataType),
) (*model.MsgCounterType, error) {
	if !e.IsCompatibleEntityType(entity) {
		return nil, api.ErrNoCompatibleEntity
	}

	filter := model.LoadControlLimitDescriptionDataType{
		LimitType:      util.Ptr(model.LoadControlLimitTypeTypeSignDependentAbsValueLimit),
		LimitCategory:  util.Ptr(model.LoadControlCategoryTypeObligation),
		LimitDirection: util.Ptr(model.EnergyDirectionTypeConsume),
		ScopeType:      util.Ptr(model.ScopeTypeTypeActivePowerLimit),
	}

	return internal.WriteLoadControlLimit(e.LocalEntity, entity, filter, limit, resultCB)
}

// Scenario 2

// return Failsafe limit for the consumed active (real) power of the
// Controllable System. This limit becomes activated in "init" state or "failsafe state".
func (e *LPC) FailsafeConsumptionActivePowerLimit(entity spineapi.EntityRemoteInterface) (float64, error) {
	if !e.IsCompatibleEntityType(entity) {
		return 0, api.ErrNoCompatibleEntity
	}

	keyname := model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit

	deviceConfiguration, err := client.NewDeviceConfiguration(e.LocalEntity, entity)
	if err != nil || deviceConfiguration == nil {
		return 0, api.ErrDataNotAvailable
	}

	filter := model.DeviceConfigurationKeyValueDescriptionDataType{
		KeyName:   &keyname,
		ValueType: util.Ptr(model.DeviceConfigurationKeyValueTypeTypeScaledNumber),
	}
	data, err := deviceConfiguration.GetKeyValueDataForFilter(filter)
	if err != nil || data == nil || data.Value == nil || data.Value.ScaledNumber == nil {
		return 0, api.ErrDataNotAvailable
	}

	return data.Value.ScaledNumber.GetValue(), nil
}

// send new Failsafe Consumption Active Power Limit
//
// parameters:
//   - entity: the entity of the e.g. EVSE
//   - value: the new limit in W
func (e *LPC) WriteFailsafeConsumptionActivePowerLimit(entity spineapi.EntityRemoteInterface, value float64) (*model.MsgCounterType, error) {
	if !e.IsCompatibleEntityType(entity) {
		return nil, api.ErrNoCompatibleEntity
	}

	keyname := model.DeviceConfigurationKeyNameTypeFailsafeConsumptionActivePowerLimit

	deviceConfiguration, err := client.NewDeviceConfiguration(e.LocalEntity, entity)
	if err != nil || deviceConfiguration == nil {
		return nil, api.ErrDataNotAvailable
	}

	filter := model.DeviceConfigurationKeyValueDescriptionDataType{
		KeyName: &keyname,
	}
	data, err := deviceConfiguration.GetKeyValueDescriptionsForFilter(filter)
	if err != nil || data == nil || len(data) != 1 {
		return nil, api.ErrDataNotAvailable
	}

	keyData := []model.DeviceConfigurationKeyValueDataType{
		{
			KeyId: data[0].KeyId,
			Value: &model.DeviceConfigurationKeyValueValueType{
				ScaledNumber: model.NewScaledNumberType(value),
			},
		},
	}

	msgCounter, err := deviceConfiguration.WriteKeyValues(keyData)

	return msgCounter, err
}

// return minimum time the Controllable System remains in "failsafe state" unless conditions
// specified in this Use Case permit leaving the "failsafe state"
func (e *LPC) FailsafeDurationMinimum(entity spineapi.EntityRemoteInterface) (time.Duration, error) {
	if !e.IsCompatibleEntityType(entity) {
		return 0, api.ErrNoCompatibleEntity
	}

	keyname := model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum

	deviceConfiguration, err := client.NewDeviceConfiguration(e.LocalEntity, entity)
	if err != nil || deviceConfiguration == nil {
		return 0, api.ErrDataNotAvailable
	}

	filter := model.DeviceConfigurationKeyValueDescriptionDataType{
		KeyName:   &keyname,
		ValueType: util.Ptr(model.DeviceConfigurationKeyValueTypeTypeDuration),
	}
	data, err := deviceConfiguration.GetKeyValueDataForFilter(filter)
	if err != nil || data == nil || data.Value == nil || data.Value.Duration == nil {
		return 0, api.ErrDataNotAvailable
	}

	return data.Value.Duration.GetTimeDuration()
}

// send new Failsafe Duration Minimum
//
// parameters:
//   - entity: the entity of the e.g. EVSE
//   - duration: the duration, between 2h and 24h
func (e *LPC) WriteFailsafeDurationMinimum(entity spineapi.EntityRemoteInterface, duration time.Duration) (*model.MsgCounterType, error) {
	if !e.IsCompatibleEntityType(entity) {
		return nil, api.ErrNoCompatibleEntity
	}

	if duration < time.Duration(time.Hour*2) || duration > time.Duration(time.Hour*24) {
		return nil, errors.New("duration outside of allowed range")
	}

	keyname := model.DeviceConfigurationKeyNameTypeFailsafeDurationMinimum

	deviceConfiguration, err := client.NewDeviceConfiguration(e.LocalEntity, entity)
	if err != nil || deviceConfiguration == nil {
		return nil, api.ErrDataNotAvailable
	}

	filter := model.DeviceConfigurationKeyValueDescriptionDataType{
		KeyName: &keyname,
	}
	data, err := deviceConfiguration.GetKeyValueDataForFilter(filter)
	if err != nil || data == nil {
		return nil, api.ErrDataNotAvailable
	}

	keyData := []model.DeviceConfigurationKeyValueDataType{
		{
			KeyId: data.KeyId,
			Value: &model.DeviceConfigurationKeyValueValueType{
				Duration: model.NewDurationType(duration),
			},
		},
	}

	msgCounter, err := deviceConfiguration.WriteKeyValues(keyData)

	return msgCounter, err
}

// Scenario 4

// return nominal maximum active (real) power the Controllable System is
// able to consume according to the device label or data sheet.
func (e *LPC) PowerConsumptionNominalMax(entity spineapi.EntityRemoteInterface) (float64, error) {
	if !e.IsCompatibleEntityType(entity) {
		return 0, api.ErrNoCompatibleEntity
	}

	electricalConnection, err := client.NewElectricalConnection(e.LocalEntity, entity)
	if err != nil || electricalConnection == nil {
		return 0, err
	}

	filter := model.ElectricalConnectionCharacteristicDataType{
		CharacteristicContext: util.Ptr(model.ElectricalConnectionCharacteristicContextTypeEntity),
		CharacteristicType:    util.Ptr(model.ElectricalConnectionCharacteristicTypeTypePowerConsumptionNominalMax),
	}
	data, err := electricalConnection.GetCharacteristicsForFilter(filter)
	if err != nil || len(data) == 0 || data[0].Value == nil {
		return 0, err
	}

	return data[0].Value.GetValue(), nil
}
