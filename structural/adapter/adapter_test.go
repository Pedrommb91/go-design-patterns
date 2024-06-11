package adapter

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestUSBDevice(t *testing.T) {
	usbDevice := &USBDevice{}
	expectedOutput := "USB device connected.\n"

	if got := captureOutput(usbDevice.ConnectWithUSBPort); got != expectedOutput {
		t.Errorf("USBDevice.ConnectWithUSBPort() = %v, want %v", got, expectedOutput)
	}
}

func TestHDMIDevice(t *testing.T) {
	hdmiDevice := &HDMIDevice{}
	expectedOutput := "HDMI device connected.\n"

	if got := captureOutput(hdmiDevice.ConnectWithHDMIPort); got != expectedOutput {
		t.Errorf("HDMIDevice.ConnectWithHDMIPort() = %v, want %v", got, expectedOutput)
	}
}

func TestHDMIToUSBAdapter(t *testing.T) {
	hdmiDevice := &HDMIDevice{}
	adapter := &HDMIToUSBAdapter{HDMIDevice: hdmiDevice}
	expectedOutput := "Adapter converts USB signal to HDMI.\nHDMI device connected.\n"

	if got := captureOutput(adapter.ConnectWithUSBPort); got != expectedOutput {
		t.Errorf("HDMIToUSBAdapter.ConnectWithUSBPort() = %v, want %v", got, expectedOutput)
	}
}

func TestConnectDeviceToComputer(t *testing.T) {
	usbDevice := &USBDevice{}
	expectedOutput := "USB device connected.\n"

	if got := captureOutput(func() { ConnectDeviceToComputer(usbDevice) }); got != expectedOutput {
		t.Errorf("ConnectDeviceToComputer(usbDevice) = %v, want %v", got, expectedOutput)
	}

	hdmiDevice := &HDMIDevice{}
	adapter := &HDMIToUSBAdapter{HDMIDevice: hdmiDevice}
	expectedOutput = "Adapter converts USB signal to HDMI.\nHDMI device connected.\n"

	if got := captureOutput(func() { ConnectDeviceToComputer(adapter) }); got != expectedOutput {
		t.Errorf("ConnectDeviceToComputer(adapter) = %v, want %v", got, expectedOutput)
	}
}

// captureOutput captures the output of a function that writes to stdout
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
