package bridge

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestTV(t *testing.T) {
	tv := &TV{}
	remote := &RemoteControl{Device: tv}

	remote.TurnOn()
	if got := captureOutput(remote.TurnOn); got != "TV is On\n" {
		t.Errorf("RemoteControl.TurnOn() = %v, want %v", got, "TV is On\n")
	}

	remote.TurnOff()
	if got := captureOutput(remote.TurnOff); got != "TV is Off\n" {
		t.Errorf("RemoteControl.TurnOff() = %v, want %v", got, "TV is Off\n")
	}
}

func TestRadio(t *testing.T) {
	radio := &Radio{}
	remote := &RemoteControl{Device: radio}

	remote.TurnOn()
	if got := captureOutput(remote.TurnOn); got != "Radio is On\n" {
		t.Errorf("RemoteControl.TurnOn() = %v, want %v", got, "Radio is On\n")
	}

	remote.TurnOff()
	if got := captureOutput(remote.TurnOff); got != "Radio is Off\n" {
		t.Errorf("RemoteControl.TurnOff() = %v, want %v", got, "Radio is Off\n")
	}
}

func TestAdvancedRemoteControl(t *testing.T) {
	radio := &Radio{}
	advancedRemote := &AdvancedRemoteControl{RemoteControl: RemoteControl{Device: radio}}

	advancedRemote.TurnOn()
	if got := captureOutput(advancedRemote.TurnOn); got != "Radio is On\n" {
		t.Errorf("AdvancedRemoteControl.TurnOn() = %v, want %v", got, "Radio is On\n")
	}

	advancedRemote.Mute()
	if got := captureOutput(advancedRemote.Mute); got != "Muting the device\n" {
		t.Errorf("AdvancedRemoteControl.Mute() = %v, want %v", got, "Muting the device\n")
	}

	advancedRemote.TurnOff()
	if got := captureOutput(advancedRemote.TurnOff); got != "Radio is Off\n" {
		t.Errorf("AdvancedRemoteControl.TurnOff() = %v, want %v", got, "Radio is Off\n")
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
