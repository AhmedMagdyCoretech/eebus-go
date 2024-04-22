package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/enbility/eebus-go/api"
	"github.com/enbility/eebus-go/cmd/usecases"
	"github.com/enbility/eebus-go/service"
	shipapi "github.com/enbility/ship-go/api"
	"github.com/enbility/ship-go/cert"

	// spineapi "github.com/enbility/spine-go/api"

	// "github.com/enbility/eebus-go/util"
	"github.com/enbility/spine-go/model"
)

var remoteSki string

var AvailablePower model.NumberType

type hems struct {
	myService     *service.Service
	myOpevHandler *usecases.OpevHandler
}

func (h *hems) run() {
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
		certificate, err = cert.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-01")
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
		"Demo", "Demo", "HEMS", "123456789",
		model.DeviceTypeTypeEnergyManagementSystem,
		[]model.EntityTypeType{model.EntityTypeTypeEVSE},
		port, certificate, 230, time.Second*4)
	if err != nil {
		log.Fatal(err)
	}
	configuration.SetAlternateIdentifier("Demo-HEMS-123456789")

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

	h.myOpevHandler = usecases.NewOpevHandler(model.UseCaseActorTypeCEM, h.myService.LocalDevice())

	h.myService.Start()
	// defer h.myService.Shutdown()
}

func (h *hems) RemoteSKIConnected(service api.ServiceInterface, ski string) {
	// go func() {
	// 	// Load Control ( Local )
	// 	local_load_control_feature := service.LocalDevice().FeatureByAddress(
	// 		service.LocalDevice().
	// 			EntityForType(model.EntityTypeTypeCEM).
	// 			FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl, model.RoleTypeClient).Address())

	// 	// Load Control ( Remote )
	// 	remote_load_control_feature := service.LocalDevice().RemoteDeviceForSki(ski).FeatureByAddress(service.LocalDevice().RemoteDeviceForSki(ski).
	// 		EntityForType(model.EntityTypeTypeEV).
	// 		FeatureOfTypeAndRole(model.FeatureTypeTypeLoadControl,
	// 			model.RoleTypeServer).Address())

	// 	/*========================================================== Runtime Scenario Communication ==========================================================*/

	// 	// Wait for some time
	// 	time.Sleep(7 * time.Second)

	// 	// Create cmd to be sent
	// 	cmd := model.CmdType{
	// 		LoadControlLimitListData: &model.LoadControlLimitListDataType{
	// 			LoadControlLimitData: []model.LoadControlLimitDataType{
	// 				{
	// 					Value: &model.ScaledNumberType{
	// 						Number: util.Ptr(model.NumberType(3)),
	// 						Scale:  util.Ptr(model.ScaleType(0)),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	}

	// 	// Write the updated cmd and save the MSG Counter to be used in CallBack later
	// 	counter, _ := remote_load_control_feature.Device().Sender().Write(local_load_control_feature.Address(),
	// 		remote_load_control_feature.Address(), cmd)

	// 	// Add CallBack for the previous write operation with the saved MSG Counter
	// 	local_load_control_feature.AddResponseCallback(*counter, func(msg spineapi.ResponseMessage) {
	// 		fmt.Println("GGEZ1")
	// 		if msg.Data == nil {
	// 			return
	// 		}
	// 	})
	// }()
}

func (h *hems) RemoteSKIDisconnected(service api.ServiceInterface, ski string) {}

func (h *hems) VisibleRemoteServicesUpdated(service api.ServiceInterface, entries []shipapi.RemoteService) {
}

func (h *hems) ServiceShipIDUpdate(ski string, shipdID string) {}

func (h *hems) ServicePairingDetailUpdate(ski string, detail *shipapi.ConnectionStateDetail) {
	if ski == remoteSki && detail.State() == shipapi.ConnectionStateRemoteDeniedTrust {
		fmt.Println("The remote service denied trust. Exiting.")
		h.myService.RegisterRemoteSKI(ski, false)
		h.myService.CancelPairingWithSKI(ski)
		h.myService.Shutdown()
		os.Exit(0)
	}
}

func (h *hems) AllowWaitingForTrust(ski string) bool {
	return ski == remoteSki
}

// UCEvseCommisioningConfigurationCemDelegate

// handle device state updates from the remote EVSE device
func (h *hems) HandleEVSEDeviceState(ski string, failure bool, errorCode string) {
	fmt.Println("EVSE Error State:", failure, errorCode)
}

// main app
func usage() {
	fmt.Println("First Run:")
	fmt.Println("  go run /cmd/hems/main.go <serverport>")
	fmt.Println()
	fmt.Println("General Usage:")
	fmt.Println("  go run /cmd/hems/main.go <serverport> <evse-ski> <crtfile> <keyfile>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	h := hems{}
	h.run()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	// User exit
}

// Logging interface

func (h *hems) Trace(args ...interface{}) {
	h.print("TRACE", args...)
}

func (h *hems) Tracef(format string, args ...interface{}) {
	h.printFormat("TRACE", format, args...)
}

func (h *hems) Debug(args ...interface{}) {
	h.print("DEBUG", args...)
}

func (h *hems) Debugf(format string, args ...interface{}) {
	h.printFormat("DEBUG", format, args...)
}

func (h *hems) Info(args ...interface{}) {
	h.print("INFO ", args...)
}

func (h *hems) Infof(format string, args ...interface{}) {
	h.printFormat("INFO ", format, args...)
}

func (h *hems) Error(args ...interface{}) {
	h.print("ERROR", args...)
}

func (h *hems) Errorf(format string, args ...interface{}) {
	h.printFormat("ERROR", format, args...)
}

func (h *hems) currentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (h *hems) print(msgType string, args ...interface{}) {
	value := fmt.Sprintln(args...)
	fmt.Printf("%s %s %s", h.currentTimestamp(), msgType, value)
}

func (h *hems) printFormat(msgType, format string, args ...interface{}) {
	value := fmt.Sprintf(format, args...)
	fmt.Println(h.currentTimestamp(), msgType, value)
}
