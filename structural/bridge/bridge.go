package bridge

import "fmt"

// Device is the implementor interface
type Device interface {
	On()
	Off()
}

// TV is a concrete implementor
type TV struct{}

func (t *TV) On() {
	fmt.Println("TV is On")
}

func (t *TV) Off() {
	fmt.Println("TV is Off")
}

// Radio is another concrete implementor
type Radio struct{}

func (r *Radio) On() {
	fmt.Println("Radio is On")
}

func (r *Radio) Off() {
	fmt.Println("Radio is Off")
}

// RemoteControl is the abstraction
type RemoteControl struct {
	Device Device
}

func (r *RemoteControl) TurnOn() {
	r.Device.On()
}

func (r *RemoteControl) TurnOff() {
	r.Device.Off()
}

// AdvancedRemoteControl is a refined abstraction
type AdvancedRemoteControl struct {
	RemoteControl
}

func (a *AdvancedRemoteControl) Mute() {
	fmt.Println("Muting the device")
}
