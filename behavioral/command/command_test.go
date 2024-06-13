package command

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestTurnOnCommand(t *testing.T) {
	light := &Light{}
	turnOnCmd := &TurnOnCommand{Light: light}

	turnOnCmd.Execute()
	if !light.IsOn {
		t.Error("Expected light to be on, but it's off")
	}
}

func TestTurnOffCommand(t *testing.T) {
	light := &Light{IsOn: true}
	turnOffCmd := &TurnOffCommand{Light: light}

	turnOffCmd.Execute()
	if light.IsOn {
		t.Error("Expected light to be off, but it's on")
	}
}

func TestSwitch(t *testing.T) {
	light := &Light{}
	switcher := NewSwitch()
	switcher.StoreCommand("on", &TurnOnCommand{Light: light})
	switcher.StoreCommand("off", &TurnOffCommand{Light: light})

	switcher.ExecuteCommand("on")
	if !light.IsOn {
		t.Error("Expected light to be on after 'on' command, but it's off")
	}

	switcher.ExecuteCommand("off")
	if light.IsOn {
		t.Error("Expected light to be off after 'off' command, but it's on")
	}

	output := captureOutput(func() {
		switcher.ExecuteCommand("nonexistent")
	})
	expectedOutput := "Command not found: nonexistent\n"
	if output != expectedOutput {
		t.Errorf("Expected output '%s', got '%s'", expectedOutput, output)
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
