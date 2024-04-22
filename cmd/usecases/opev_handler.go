package usecases

import (
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
	"github.com/enbility/spine-go/util"
)

type OpevHandler struct {
	actorType     model.UseCaseActorType
	localDevice   spineapi.DeviceLocalInterface
	remoteDevices []*spineapi.DeviceRemoteInterface
}

func (h *OpevHandler) HandleEvent(payload spineapi.EventPayload) {
	// Check if the actor is EV
	if h.actorType == model.UseCaseActorTypeEV {
		// Check if the EV is added
		if payload.EventType == spineapi.EventTypeDeviceChange && payload.ChangeType == spineapi.ElementChangeAdd {
			// Accept only one CEM to be added to the EV ( only and only one Energy Guard )
			if len(h.remoteDevices) == 1 {
				return
			}

			cemEntity := payload.Device.EntityForType(model.EntityTypeTypeCEM) // CEM Entity
			// If CEM not found return
			if cemEntity == nil {
				return
			}

			// look for required features

			// 1. LoadControl Feature
			{
				feature := cemEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeClient)
				if feature == nil {
					// Error Essential feature not found
					return
				}
			}
			// 2. ElectricalConnection Feature
			{
				feature := cemEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeElectricalConnection, model.RoleTypeClient)
				if feature == nil {
					// Error Essential feature not found
					return
				}
			}

			// Accept remote device
			h.remoteDevices = append(h.remoteDevices, &payload.Device)
		}
		return
	}

	// Check if the actor is CEM
	if h.actorType == model.UseCaseActorTypeCEM {
		// Check if the EV is added
		if payload.EventType == spineapi.EventTypeDeviceChange && payload.ChangeType == spineapi.ElementChangeAdd {
			var evEntity spineapi.EntityRemoteInterface // EV Entity

			// look for an EVSE->EV entity
			for _, e := range payload.Device.Entities() {
				if e.EntityType() != model.EntityTypeTypeEV {
					continue
				}
				if len(e.Address().Entity) != 2 {
					continue
				}
				if h.localDevice.Entity([]model.AddressEntityType{e.Address().Entity[0]}).EntityType() != model.EntityTypeTypeEVSE {
					continue
				}
				evEntity = e
				break
			}

			// If EV Entity not found return
			if evEntity == nil {
				return
			}

			Scenario1_InitialScenarioCommunication(payload, h, evEntity, true)

		}
		return
	}
}

func NewOpevHandler(aType model.UseCaseActorType, d spineapi.DeviceLocalInterface) *OpevHandler {
	h := OpevHandler{
		actorType:   aType,
		localDevice: d,
	}

	if h.actorType == model.UseCaseActorTypeEV {
		addEV(d.(*spine.DeviceLocal))
	}

	if h.actorType == model.UseCaseActorTypeCEM {
		addCem(d.(*spine.DeviceLocal))
	}

	spine.Events.Subscribe(&h)
	return &h
}

func Scenario1_InitialScenarioCommunication(payload spineapi.EventPayload, h *OpevHandler, evEntity spineapi.EntityRemoteInterface, preFlag bool) {
	var track_counter *model.MsgCounterType // MSG counter tracker

	/*============================================================ Pre Scenario Communication =============================================================*/
	if preFlag {
		// look for features to bind
		var features []spineapi.FeatureRemoteInterface

		// 1. LoadControl Feature
		{
			feature := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeServer)
			if feature == nil {
				return
			}
			features = append(features, feature)
		}

		// bind to features
		for _, feature := range features {
			h.localDevice.EntityForType(model.EntityTypeTypeCEM).FeatureOfTypeAndRole(feature.Type(), model.RoleTypeClient).BindToRemote(feature.Address())
		}

		// look for features to subscribe
		features = []spineapi.FeatureRemoteInterface{}

		// 1. LoadControl Feature
		{
			feature := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeServer)
			if feature == nil {
				return
			}
			features = append(features, feature)
		}
		// 2. ElectricalConnection Feature
		{
			feature := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)
			if feature == nil {
				return
			}
			features = append(features, feature)
		}

		// subscribe to features
		for _, feature := range features {
			track_counter, _ = h.localDevice.EntityForType(model.EntityTypeTypeCEM).FeatureOfTypeAndRole(feature.Type(), model.RoleTypeClient).
				SubscribeToRemote(feature.Address())
		}

		// Accept remote device
		h.remoteDevices = append(h.remoteDevices, &payload.Device)
	}

	/*========================================================== Initial Scenario Communication ==========================================================*/
	//
	// Variables (Entity)
	//
	LocalEntity := h.localDevice.EntityForType(model.EntityTypeTypeCEM) // CEM

	//
	// Variables (Feature)
	//
	local_load_control_feature := LocalEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeClient)
	remote_load_control_feature := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeServer)
	local_electrical_connection_feature := LocalEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeElectricalConnection, model.RoleTypeClient)
	remote_electrical_connection_feature := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)

	//
	// Scenario Sequence
	//
	h.localDevice.NodeManagement().AddResultCallback(func(msg spineapi.ResponseMessage) {
		if msg.MsgCounterReference == *track_counter {
			// LoadControlLimitDescriptionListData (read, reply)
			track_counter, _ = local_load_control_feature.RequestRemoteData(model.FunctionTypeLoadControlLimitDescriptionListData, nil,
				nil, remote_load_control_feature)

			local_load_control_feature.AddResponseCallback(*track_counter, func(msg spineapi.ResponseMessage) {
				// LoadControlLimitListData (read, reply)
				track_counter, _ = local_load_control_feature.RequestRemoteData(model.FunctionTypeLoadControlLimitListData, nil,
					nil, remote_load_control_feature)
			})

			local_load_control_feature.AddResponseCallback(*track_counter+1, func(msg spineapi.ResponseMessage) {
				// ElectricalConnectionParameterDescriptionListDat (read, reply)
				track_counter, _ = local_electrical_connection_feature.RequestRemoteData(model.FunctionTypeElectricalConnectionParameterDescriptionListData, nil,
					nil, remote_electrical_connection_feature)
			})

			local_electrical_connection_feature.AddResponseCallback(*track_counter+2, func(msg spineapi.ResponseMessage) {
				// ElectricalConnectionPermittedValueSetListData (read, reply)
				track_counter, _ = local_electrical_connection_feature.RequestRemoteData(model.FunctionTypeElectricalConnectionPermittedValueSetListData, nil,
					nil, remote_electrical_connection_feature)
			})
		}
	})
}

func addEV(r *spine.DeviceLocal) {
	entityType := model.EntityTypeTypeEV
	entity_evse := r.EntityForType(model.EntityTypeTypeEVSE)

	/* Check if EVSE entity is created */
	if entity_evse == nil {
		// If not create EVSE Entity
		entityAddressId := model.AddressEntityType(len(r.Entities()))              // Entity ID
		entityAddress := []model.AddressEntityType{entityAddressId}                // Entity Address derived from ID
		entity := spine.NewEntityLocal(r, model.EntityTypeTypeEVSE, entityAddress) // EVSE Entity
		r.AddEntity(entity)                                                        // Add EVSE Entity to the device
		entity_evse = entity
	}

	// After making sure that EVSE entity is created, create EV entity next
	entityAddressId := model.AddressEntityType(len(r.Entities()))                                // Entity ID
	entityAddress := []model.AddressEntityType{entity_evse.Address().Entity[0], entityAddressId} // Entity Address derived from ID
	entity := spine.NewEntityLocal(entity_evse.Device(), entityType, entityAddress)              // EV Entity
	r.AddEntity(entity)                                                                          // Add EV Entity to the device

	/* Add * EV Commissioning & Configuration * UseCase */
	entity.AddUseCaseSupport(model.UseCaseActorTypeEV, model.UseCaseNameTypeEVCommissioningAndConfiguration, "0.0.0", "0",
		true, []model.UseCaseScenarioSupportType{1, 6, 8})

	/* Add * Overload Protection by EV Charging Current Curtailment * UseCase */
	entity.AddUseCaseSupport(model.UseCaseActorTypeEV, model.UseCaseNameTypeOverloadProtectionByEVChargingCurrentCurtailment,
		"0.0.0", "0", true, []model.UseCaseScenarioSupportType{1})

	{
		/** electrical connection feature - server **/
		f := spine.NewFeatureLocal(entity.NextFeatureId(), entity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)

		/** Electrical Connection Parameter Description Function **/
		f.AddFunctionType(model.FunctionTypeElectricalConnectionParameterDescriptionListData, true, false)
		ElecricalConnectionlDescription := &model.ElectricalConnectionParameterDescriptionListDataType{
			ElectricalConnectionParameterDescriptionData: []model.ElectricalConnectionParameterDescriptionDataType{
				{
					ElectricalConnectionId: util.Ptr(model.ElectricalConnectionIdType(1)),
					ParameterId:            util.Ptr(model.ElectricalConnectionParameterIdType(2)),
					MeasurementId:          util.Ptr(model.MeasurementIdType(2)),
					AcMeasuredPhases:       util.Ptr(model.ElectricalConnectionPhaseNameTypeAbc),
					ScopeType:              util.Ptr(model.ScopeTypeTypeACPowerTotal),
				},
			},
		}
		f.SetData(model.FunctionTypeElectricalConnectionParameterDescriptionListData, ElecricalConnectionlDescription)

		/** Electrical Connection Permitted Value Set Function **/
		f.AddFunctionType(model.FunctionTypeElectricalConnectionPermittedValueSetListData, true, false)
		// Permitted value set to be added
		permittedValue := model.ScaledNumberSetType{
			Range: []model.ScaledNumberRangeType{
				{
					Min: util.Ptr(model.ScaledNumberType{
						Number: util.Ptr(model.NumberType(1)),
					}),
					Max: util.Ptr(model.ScaledNumberType{
						Number: util.Ptr(model.NumberType(5)),
					}),
				},
			},
			Value: []model.ScaledNumberType{
				{
					Number: util.Ptr(model.NumberType(3)),
					Scale:  util.Ptr(model.ScaleType(0)),
				},
			},
		}
		// Electrical Connection Value
		ElecricalConnectionlPermittedValue := &model.ElectricalConnectionPermittedValueSetListDataType{
			ElectricalConnectionPermittedValueSetData: []model.ElectricalConnectionPermittedValueSetDataType{
				{
					ElectricalConnectionId: util.Ptr(model.ElectricalConnectionIdType(1)),
					ParameterId:            util.Ptr(model.ElectricalConnectionParameterIdType(2)),
					PermittedValueSet:      []model.ScaledNumberSetType{permittedValue},
				},
			},
		}
		f.SetData(model.FunctionTypeElectricalConnectionPermittedValueSetListData, ElecricalConnectionlPermittedValue)

		/** Add the feature to the entity **/
		entity.AddFeature(f)
	}
	{
		/** load control feature - server **/
		f := spine.NewFeatureLocal(entity.NextFeatureId(), entity, model.FeatureTypeTypeLoadControl, model.RoleTypeServer)

		/** load control limit description Function **/
		f.AddFunctionType(model.FunctionTypeLoadControlLimitDescriptionListData, true, false)
		LoadControlDescription := &model.LoadControlLimitDescriptionListDataType{
			LoadControlLimitDescriptionData: []model.LoadControlLimitDescriptionDataType{
				{
					LimitId:       util.Ptr(model.LoadControlLimitIdType(1)),
					LimitType:     util.Ptr(model.LoadControlLimitTypeTypeMaxValueLimit),
					LimitCategory: util.Ptr(model.LoadControlCategoryTypeObligation),
					MeasurementId: util.Ptr(model.MeasurementIdType(2)),
					Unit:          util.Ptr(model.UnitOfMeasurementTypeA),
					ScopeType:     util.Ptr(model.ScopeTypeTypeOverloadProtection),
				},
			},
		}
		f.SetData(model.FunctionTypeLoadControlLimitDescriptionListData, LoadControlDescription)

		/** load control limit data Function **/
		f.AddFunctionType(model.FunctionTypeLoadControlLimitListData, true, true)

		LoadControlData := &model.LoadControlLimitListDataType{
			LoadControlLimitData: []model.LoadControlLimitDataType{
				{
					LimitId:           util.Ptr(model.LoadControlLimitIdType(1)),
					IsLimitChangeable: util.Ptr(true),
					IsLimitActive:     util.Ptr(true),
					Value: &model.ScaledNumberType{
						Number: util.Ptr(model.NumberType(2)),
						Scale:  util.Ptr(model.ScaleType(0)),
					},
				},
			},
		}
		f.SetData(model.FunctionTypeLoadControlLimitListData, LoadControlData)

		/** Add the feature to the entity **/
		entity.AddFeature(f)
	}
}

func addCem(r *spine.DeviceLocal) {
	entityType := model.EntityTypeTypeCEM
	entity := r.EntityForType(entityType)

	/* Check if EVSE entity is created */
	if entity == nil {
		// If not create EVSE Entity
		entityAddressId := model.AddressEntityType(len(r.Entities())) // Entity ID
		entityAddress := []model.AddressEntityType{entityAddressId}   // Entity Address derived from ID
		entity = spine.NewEntityLocal(r, entityType, entityAddress)   // CEM Entity
		r.AddEntity(entity)                                           // Add CEM Entity to the device
	}

	/* Add * EV Commissioning & Configuration * UseCase */
	entity.AddUseCaseSupport(model.UseCaseActorTypeCEM, model.UseCaseNameTypeEVCommissioningAndConfiguration, "0.0.0", "0",
		true, []model.UseCaseScenarioSupportType{1, 6, 8})

	/* Add * Overload Protection by EV Charging Current Curtailment * UseCase */
	entity.AddUseCaseSupport(model.UseCaseActorTypeCEM, model.UseCaseNameTypeOverloadProtectionByEVChargingCurrentCurtailment,
		"0.0.0", "0", true, []model.UseCaseScenarioSupportType{1})

	{
		/** electrical connection feature - client **/
		f := spine.NewFeatureLocal(entity.NextFeatureId(), entity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeClient)

		/** Add the feature to the entity **/
		entity.AddFeature(f)
	}
	{
		/** electrical connection feature - client **/
		f := spine.NewFeatureLocal(entity.NextFeatureId(), entity, model.FeatureTypeTypeLoadControl, model.RoleTypeClient)

		/** Add the feature to the entity **/
		entity.AddFeature(f)
	}
}
