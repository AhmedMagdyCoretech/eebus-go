package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"

	// "encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/service"
	"github.com/enbility/eebus-go/usecases"
	"github.com/enbility/eebus-go/util"
	shipapi "github.com/enbility/ship-go/api"
	"github.com/enbility/ship-go/cert"
	"github.com/enbility/spine-go/model"
	"github.com/enbility/spine-go/spine"
)

var remoteSki string

type evse struct {
	myService     *service.Service
	myOpevHandler *usecases.OpevHandler
}

func (h *evse) run() {
	var err error
	var certificate tls.Certificate

	if len(os.Args) == 5 {
		remoteSki = os.Args[2]

		certificate, err = tls.LoadX509KeyPair(os.Args[3], os.Args[4])
		if err != nil {
			usage()
			log.Fatal(err)
		}
	} else {
		certificate, err = cert.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-02")
		if err != nil {
			log.Fatal(err)
		}

		pemdata := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certificate.Certificate[0],
		})
		fmt.Println(string(pemdata))

		b, err := x509.MarshalECPrivateKey(certificate.PrivateKey.(*ecdsa.PrivateKey))
		if err != nil {
			log.Fatal(err)
		}
		pemdata = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
		fmt.Println(string(pemdata))
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		log.Fatal(err)
	}

	configuration, err := api.NewConfiguration(
		"Demo", "Demo", "EVSE", "234567890",
		model.DeviceTypeTypeChargingStation,
		[]model.EntityTypeType{model.EntityTypeTypeEVSE},
		port, certificate, 230, time.Second*4)
	if err != nil {
		log.Fatal(err)
	}
	configuration.SetAlternateIdentifier("Demo-EVSE-234567890")

	h.myService = service.NewService(configuration, h)
	h.myService.SetLogging(h)

	if err = h.myService.Setup(); err != nil {
		fmt.Println(err)
		return
	}

	if len(remoteSki) == 0 {
		os.Exit(0)
	}

	h.myService.RegisterRemoteSKI(remoteSki, true)

	// addEV(h.myService.LocalDevice().(*spine.DeviceLocal))
	h.myService.Start()
	// defer h.myService.Shutdown()
}

// EEBUSServiceHandler

func (h *evse) RemoteSKIConnected(service api.ServiceInterface, ski string) {
	go func() {
		time.Sleep(5 * time.Second)
		addEV(h.myService.LocalDevice().(*spine.DeviceLocal))

		time.Sleep(5 * time.Second)

		f := service.LocalDevice().EntityForType(model.EntityTypeTypeEV).FeatureOfTypeAndRole(model.FeatureTypeTypeElectricalConnection, model.RoleTypeServer)

		ElecricalConnectionlDescription := &model.ElectricalConnectionParameterDescriptionListDataType{
			ElectricalConnectionParameterDescriptionData: []model.ElectricalConnectionParameterDescriptionDataType{
				{
					ElectricalConnectionId: util.Ptr(model.ElectricalConnectionIdType(3)),
					ParameterId:            util.Ptr(model.ElectricalConnectionParameterIdType(4)),
				},
			},
		}
		f.SetData(model.FunctionTypeElectricalConnectionParameterDescriptionListData, ElecricalConnectionlDescription)
		println("GGEZ")
	}()
}

func (h *evse) RemoteSKIDisconnected(service api.ServiceInterface, ski string) {}

func (h *evse) VisibleRemoteServicesUpdated(service api.ServiceInterface, entries []shipapi.RemoteService) {
}

func (h *evse) ServiceShipIDUpdate(ski string, shipdID string) {}

func (h *evse) ServicePairingDetailUpdate(ski string, detail *shipapi.ConnectionStateDetail) {
	if ski == remoteSki && detail.State() == shipapi.ConnectionStateRemoteDeniedTrust {
		fmt.Println("The remote service denied trust. Exiting.")
		h.myService.RegisterRemoteSKI(ski, false)
		h.myService.CancelPairingWithSKI(ski)
		h.myService.Shutdown()
		os.Exit(0)
	}
}

func (h *evse) AllowWaitingForTrust(ski string) bool {
	return ski == remoteSki
}

// main app
func usage() {
	fmt.Println("First Run:")
	fmt.Println("  go run /cmd/evse/main.go <serverport>")
	fmt.Println()
	fmt.Println("General Usage:")
	fmt.Println("  go run /cmd/evse/main.go <serverport> <hems-ski> <crtfile> <keyfile>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	h := evse{}
	h.run()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit
}

// Logging interface

func (h *evse) Trace(args ...interface{}) {
	h.print("TRACE", args...)
}

func (h *evse) Tracef(format string, args ...interface{}) {
	h.printFormat("TRACE", format, args...)
}

func (h *evse) Debug(args ...interface{}) {
	h.print("DEBUG", args...)
}

func (h *evse) Debugf(format string, args ...interface{}) {
	h.printFormat("DEBUG", format, args...)
}

func (h *evse) Info(args ...interface{}) {
	h.print("INFO ", args...)
}

func (h *evse) Infof(format string, args ...interface{}) {
	h.printFormat("INFO ", format, args...)
}

func (h *evse) Error(args ...interface{}) {
	h.print("ERROR", args...)
}

func (h *evse) Errorf(format string, args ...interface{}) {
	h.printFormat("ERROR", format, args...)
}

func (h *evse) currentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (h *evse) print(msgType string, args ...interface{}) {
	value := fmt.Sprintln(args...)
	fmt.Printf("%s %s %s", h.currentTimestamp(), msgType, value)
}

func (h *evse) printFormat(msgType, format string, args ...interface{}) {
	value := fmt.Sprintf(format, args...)
	fmt.Println(h.currentTimestamp(), msgType, value)
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
	}

	// After making sure that EVSE entity is created, create EV entity next
	entityAddressId := model.AddressEntityType(len(r.Entities()))                   // Entity ID
	entityAddress := []model.AddressEntityType{entityAddressId}                     // Entity Address derived from ID
	entity := spine.NewEntityLocal(entity_evse.Device(), entityType, entityAddress) // EV Entity
	r.AddEntity(entity)                                                             // Add EV Entity to the device
	entity_evse.Address().Entity = append(entity_evse.Address().Entity, entityAddressId)

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
