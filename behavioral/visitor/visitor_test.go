package visitor

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestVisitorPattern(t *testing.T) {
	computer := NewComputer()
	displayVisitor := &ComputerPartDisplayVisitor{}

	expected := "Displaying Keyboard.\nDisplaying Monitor.\nDisplaying Mouse.\nDisplaying Computer.\n"
	output := captureOutput(computer.Accept, displayVisitor)
	if output != expected {
		t.Errorf("expected %q but got %q", expected, output)
	}
}

// captureOutput captures the output of a function that writes to stdout
func captureOutput(f func(visitor ComputerPartVisitor), visitor ComputerPartVisitor) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f(visitor)

	w.Close()
	os.Stdout = old // restoring the real stdout

	var buf bytes.Buffer
	io.Copy(&buf, r) // copy the output to the buffer
	r.Close()

	return buf.String()
}
