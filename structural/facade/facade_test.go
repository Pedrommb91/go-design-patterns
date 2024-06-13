package facade

import (
	"bytes"
	"io"
	"os"
	"testing"
)

// TestHomeTheaterFacade tests the HomeTheaterFacade's methods
func TestHomeTheaterFacade(t *testing.T) {
	// Create the components of the Home Theater
	amplifier := &Amplifier{}
	dvdPlayer := &DvdPlayer{}
	projector := &Projector{}
	theaterLights := &TheaterLights{}
	screen := &Screen{}

	// Create the Home Theater Facade
	homeTheater := NewHomeTheaterFacade(amplifier, dvdPlayer, projector, theaterLights, screen)

	// Test watching a movie
	output := captureOutput(func() {
		homeTheater.WatchMovie("Inception")
	})
	expectedOutput := "TheaterLights are dimmed to intensity: 10\n" +
		"Screen is down\n" +
		"Projector is turned on\n" +
		"Amplifier is turned on\n" +
		"DvdPlayer is turned on\n" +
		"DvdPlayer is playing movie: Inception\n"
	if output != expectedOutput {
		t.Errorf("Expected output of WatchMovie does not match. Got: %s, Want: %s", output, expectedOutput)
	}

	// Test ending a movie
	output = captureOutput(func() {
		homeTheater.EndMovie()
	})
	expectedOutput = "DvdPlayer has stopped\n" +
		"DvdPlayer has ejected the disk\n" +
		"DvdPlayer is turned off\n" +
		"Amplifier is turned off\n" +
		"Projector is turned off\n" +
		"Screen is up\n" +
		"TheaterLights are turned on\n"
	if output != expectedOutput {
		t.Errorf("Expected output of EndMovie does not match. Got: %s, Want: %s", output, expectedOutput)
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
