package prototype

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestCircleClone(t *testing.T) {
	original := &Circle{Radius: 10, Color: "red"}
	clone := original.Clone().(*Circle)

	if original == clone {
		t.Errorf("expected clone to be a different instance")
	}
	if original.Radius != clone.Radius {
		t.Errorf("expected radius to be %d, but got %d", original.Radius, clone.Radius)
	}
	if clone.Color != "red_clone" {
		t.Errorf("expected color to be %s, but got %s", "red_clone", clone.Color)
	}
}

func TestRectangleClone(t *testing.T) {
	original := &Rectangle{Width: 20, Height: 10, Color: "blue"}
	clone := original.Clone().(*Rectangle)

	if original == clone {
		t.Errorf("expected clone to be a different instance")
	}
	if original.Width != clone.Width {
		t.Errorf("expected width to be %d, but got %d", original.Width, clone.Width)
	}
	if original.Height != clone.Height {
		t.Errorf("expected height to be %d, but got %d", original.Height, clone.Height)
	}
	if clone.Color != "blue_clone" {
		t.Errorf("expected color to be %s, but got %s", "blue_clone", clone.Color)
	}
}

func TestCircleDraw(t *testing.T) {
	circle := &Circle{Radius: 10, Color: "red"}
	expectedOutput := "Drawing a red circle with radius 10\n"

	if output := captureOutput(circle.Draw); output != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, output)
	}
}

func TestRectangleDraw(t *testing.T) {
	rectangle := &Rectangle{Width: 20, Height: 10, Color: "blue"}
	expectedOutput := "Drawing a blue rectangle with width 20 and height 10\n"

	if output := captureOutput(rectangle.Draw); output != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, output)
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
