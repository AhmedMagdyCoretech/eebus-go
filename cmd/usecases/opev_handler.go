package usecases

import (
	spineapi "github.com/enbility/spine-go/api"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

type OpevHandler struct {
	actorType     model.UseCaseActorType
	localDevice   spineapi.DeviceLocalInterface
	remoteDevices []*spineapi.DeviceRemoteInterface
}

func (h *OpevHandler) HandleEvent(payload spineapi.EventPayload) {
	if h.actorType == model.UseCaseActorTypeEV {
		// new device added
		if payload.EventType == spineapi.EventTypeDeviceChange &&
			payload.ChangeType == spineapi.ElementChangeAdd {

			// only allow one CEM
			if len(h.remoteDevices) == 1 {
				return
			}
			cemEntity := payload.Device.EntityForType(model.EntityTypeTypeCEM)
			if cemEntity == nil {
				return
			}

			// look for required features
			var feats []spineapi.FeatureRemoteInterface
			{
				feat := cemEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
				if feat == nil {
					return
				}
				feats = append(feats, feat)
			}

			// subscribe to features
			for _, feat := range feats {
				h.localDevice.EntityForType(model.EntityTypeTypeEV).
					FeatureOfTypeAndRole(feat.Type(), model.RoleTypeClient).
					SubscribeToRemote(feat.Address())
			}

			// accept remote device
			h.remoteDevices = append(h.remoteDevices, &payload.Device)
		}
		return
	}

	if h.actorType == model.UseCaseActorTypeCEM {
		// new device added
		if payload.EventType == spineapi.EventTypeDeviceChange &&
			payload.ChangeType == spineapi.ElementChangeAdd {
			// look for an EVSE->EV entity
			var evEntity spineapi.EntityRemoteInterface
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
			if evEntity == nil {
				return
			}

			// look for features
			var feats []spineapi.FeatureRemoteInterface
			{
				feat := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeServer)
				if feat == nil {
					return
				}
				feats = append(feats, feat)
			}
			// bind to features
			for _, feat := range feats {
				h.localDevice.EntityForType(model.EntityTypeTypeCEM).
					FeatureOfTypeAndRole(feat.Type(), model.RoleTypeClient).
					BindToRemote(feat.Address())
			}

			// look for features
			feats = []spineapi.FeatureRemoteInterface{}
			{
				feat := evEntity.FeatureOfTypeAndRole(model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)
				if feat == nil {
					return
				}
				feats = append(feats, feat)
			}

			// subscribe to features
			for _, feat := range feats {
				h.localDevice.EntityForType(model.EntityTypeTypeCEM).
					FeatureOfTypeAndRole(feat.Type(), model.RoleTypeClient).
					SubscribeToRemote(feat.Address())
			}

			// accept remote device
			h.remoteDevices = append(h.remoteDevices, &payload.Device)
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
		evseAddress := []model.AddressEntityType{2}
		evAddress := append(evseAddress, 9)
		h.localDevice.AddEntity(spine.NewEntityLocal(
			h.localDevice,
			model.EntityTypeTypeEVSE,
			evseAddress))
		evEntity := spine.NewEntityLocal(
			h.localDevice,
			model.EntityTypeTypeEV,
			evAddress,
		)
		{
			f := spine.NewFeatureLocal(0, evEntity, model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeClient)
			f.AddFunctionType(model.FunctionTypeDeviceDiagnosisHeartbeatData, true, false)
			f.AddFunctionType(model.FunctionTypeDeviceDiagnosisStateData, true, false)
			evEntity.AddFeature(f)
		}
		{
			f := spine.NewFeatureLocal(1, evEntity, model.FeatureTypeTypeLoadControl, model.RoleTypeServer)
			evEntity.AddFeature(f)
		}
		{
			f := spine.NewFeatureLocal(2, evEntity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)
			evEntity.AddFeature(f)
		}

		evEntity.AddUseCaseSupport(model.UseCaseActorTypeEV, model.UseCaseNameTypeOverloadProtectionByEVChargingCurrentCurtailment, "0.0.0", "0", true, []model.UseCaseScenarioSupportType{1, 2, 3})

		h.localDevice.AddEntity(evEntity)
	}

	if h.actorType == model.UseCaseActorTypeCEM {
		cemAddress := []model.AddressEntityType{2}
		cemEntity := spine.NewEntityLocal(
			h.localDevice,
			model.EntityTypeTypeCEM,
			cemAddress,
		)

		{
			f := spine.NewFeatureLocal(0, cemEntity, model.FeatureTypeTypeLoadControl, model.RoleTypeClient)
			cemEntity.AddFeature(f)
		}
		{
			f := spine.NewFeatureLocal(1, cemEntity, model.FeatureTypeTypeElectricalConnection, model.RoleTypeClient)
			cemEntity.AddFeature(f)
		}
		{
			f := spine.NewFeatureLocal(2, cemEntity, model.FeatureTypeTypeDeviceDiagnosis, model.RoleTypeServer)
			f.AddFunctionType(model.FunctionTypeDeviceDiagnosisHeartbeatData, true, false)
			f.AddFunctionType(model.FunctionTypeDeviceDiagnosisStateData, true, false)
			h.localDevice.HeartbeatManager().SetLocalFeature(cemEntity, f)
			cemEntity.AddFeature(f)
		}

		cemEntity.AddUseCaseSupport(model.UseCaseActorTypeCEM, model.UseCaseNameTypeOverloadProtectionByEVChargingCurrentCurtailment, "0.0.0", "0", true, []model.UseCaseScenarioSupportType{1, 2, 3})

		h.localDevice.AddEntity(cemEntity)
	}

	spine.Events.Subscribe(&h)
	return &h
}
