package flyweight

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestTree_Draw(t *testing.T) {
	treeType := NewTreeType("Oak", "Brown", "Rough")
	tree := NewTree(1, 2, 100, 200, treeType)
	canvas := &Canvas{}

	output := captureOutput(func() {
		tree.Draw(canvas)
	})

	expectedOutput := "Drawing tree of type 'Oak' with color 'Brown' and texture 'Rough' at position (1, 2) with height 100 and width 200\n"
	if output != expectedOutput {
		t.Errorf("Expected output of tree.Draw does not match. Got: %s, Want: %s", output, expectedOutput)
	}
}

func TestForest_PlantTree(t *testing.T) {
	forest := NewForest()
	canvas := &Canvas{}

	forest.PlantTree(1, 2, 100, 200, "Oak", "Brown", "Rough")
	forest.PlantTree(3, 4, 150, 250, "Pine", "Green", "Smooth")

	output := captureOutput(func() {
		forest.Draw(canvas)
	})

	expectedOutput := "Drawing tree of type 'Oak' with color 'Brown' and texture 'Rough' at position (1, 2) with height 100 and width 200\n" +
		"Drawing tree of type 'Pine' with color 'Green' and texture 'Smooth' at position (3, 4) with height 150 and width 250\n"
	if output != expectedOutput {
		t.Errorf("Expected output of forest.Draw does not match. Got: %s, Want: %s", output, expectedOutput)
	}
}

func TestForest_PlantingAndDrawing(t *testing.T) {
	forest := NewForest()
	canvas := &Canvas{}

	// Planting 5 pine trees in the forest at different locations
	forest.PlantTree(1, 2, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(2, 3, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(3, 4, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(4, 5, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(5, 6, 5, 2, "Pine", "Green", "Pine Texture")

	// Planting 3 oak trees in the forest at different locations
	forest.PlantTree(6, 7, 10, 5, "Oak", "Brown", "Oak Texture")
	forest.PlantTree(7, 8, 10, 5, "Oak", "Brown", "Oak Texture")
	forest.PlantTree(8, 9, 10, 5, "Oak", "Brown", "Oak Texture")

	output := captureOutput(func() {
		forest.Draw(canvas)
	})

	expectedOutput := ""
	for i := 1; i <= 5; i++ {
		expectedOutput += fmt.Sprintf("Drawing tree of type 'Pine' with color 'Green' and texture 'Pine Texture' at position (%d, %d) with height 5 and width 2\n", i, i+1)
	}
	for i := 6; i <= 8; i++ {
		expectedOutput += fmt.Sprintf("Drawing tree of type 'Oak' with color 'Brown' and texture 'Oak Texture' at position (%d, %d) with height 10 and width 5\n", i, i+1)
	}

	if output != expectedOutput {
		t.Errorf("Expected output of forest.Draw does not match. Got: %s, Want: %s", output, expectedOutput)
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
