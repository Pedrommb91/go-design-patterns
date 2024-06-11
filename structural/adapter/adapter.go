package adapter

import "fmt"

// USB is the target interface that the client expects
type USB interface {
	ConnectWithUSBPort()
}

// USBDevice is a concrete implementation of USB
type USBDevice struct{}

func (u *USBDevice) ConnectWithUSBPort() {
	fmt.Println("USB device connected.")
}

// HDMI is the adaptee interface that needs to be adapted
type HDMI interface {
	ConnectWithHDMIPort()
}

// HDMIDevice is a concrete implementation of HDMI
type HDMIDevice struct{}

func (h *HDMIDevice) ConnectWithHDMIPort() {
	fmt.Println("HDMI device connected.")
}

// HDMIToUSBAdapter is the adapter that makes HDMIDevice compatible with USB
type HDMIToUSBAdapter struct {
	HDMIDevice *HDMIDevice
}

func (a *HDMIToUSBAdapter) ConnectWithUSBPort() {
	fmt.Println("Adapter converts USB signal to HDMI.")
	a.HDMIDevice.ConnectWithHDMIPort()
}

// Client code
func ConnectDeviceToComputer(usb USB) {
	usb.ConnectWithUSBPort()
}
