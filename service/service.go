package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-ble/ble"
	"github.com/go-ble/ble/linux"
)

const (
	characteristicUUID = "1822"
	serviceUUID        = "1821"
)

func must(desc string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", desc, err)
	}
}

func service() {
	// Initialize the Bluetooth adapter.
	d, err := linux.NewDevice()
	must("new device", err)
	defer d.Stop()

	//setup the default device.
	ble.SetDefaultDevice(d)

	// define context
	timeout := 10000 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	//Creating service to have communication using the specified service..
	service := ble.NewService(ble.MustParse(serviceUUID))
	// Creating characteristic to handle data writing...

	// Facing Problems while creating handlers...
	characteristic := service.NewCharacteristic(ble.MustParse(characteristicUUID))
	//

	characteristic.HandleWrite(ble.WriteHandlerFunc(handleRead))
	characteristic.HandleRead(ble.ReadHandlerFunc(handleWrite))

	//Adding the service...
	ble.AddService(service)

	//Defining advertise context...
	advertiseCtx := ble.WithSigHandler(ctx, cancel)
	fmt.Println("Advertising....")

	//Start the service...
	AdvertiseNameandServices(advertiseCtx, service)

}

func AdvertiseNameandServices(ctx context.Context, service *ble.Service) {
	err := ble.AdvertiseNameAndServices(ctx, "TestAttendanceServer", service.UUID)
	if err != nil {
		log.Fatal("Advertisement error occured:", err)
	}

}

func handleWrite(req ble.Request, rsp ble.ResponseWriter) {
	data, _ := rsp.Write([]byte("hello"))

	fmt.Printf("Sent data: %v\n", data)
}

func handleRead(req ble.Request, rsp ble.ResponseWriter) {
	data := req.Data()
	fmt.Printf("Recieved data: %v\n", string(data[:]))

}