package flyweight

import "fmt"

// TreeType acts as the Flyweight in this example. It stores intrinsic state that is common for many trees in the forest (name, color, texture).
type TreeType struct {
	name    string
	color   string
	texture string
}

// NewTreeType is the factory method to create a new TreeType. It ensures that for each combination of name, color, and texture, only one instance of TreeType is created.
func NewTreeType(name, color, texture string) *TreeType {
	// In a complete implementation, this function would check if a TreeType with the given properties already exists to reuse it.
	return &TreeType{name: name, color: color, texture: texture}
}

// Tree stores extrinsic state that is unique for each tree (position, height, width).
type Tree struct {
	x, y          int
	height, width int
	treeType      *TreeType
}

// NewTree creates a new Tree with the given extrinsic state and a shared TreeType.
func NewTree(x, y, height, width int, treeType *TreeType) *Tree {
	return &Tree{x: x, y: y, height: height, width: width, treeType: treeType}
}

// Draw simulates drawing the tree on the screen.
func (t *Tree) Draw(canvas *Canvas) {
	fmt.Printf("Drawing tree of type '%s' with color '%s' and texture '%s' at position (%d, %d) with height %d and width %d\n",
		t.treeType.name, t.treeType.color, t.treeType.texture, t.x, t.y, t.height, t.width)
}

// Forest is the context in which Trees are used. It might represent a game map or a landscape.
type Forest struct {
	trees []*Tree
}

// NewForest creates a new Forest.
func NewForest() *Forest {
	return &Forest{}
}

// PlantTree adds a tree to the forest.
func (f *Forest) PlantTree(x, y, height, width int, name, color, texture string) {
	treeType := NewTreeType(name, color, texture)
	tree := NewTree(x, y, height, width, treeType)
	f.trees = append(f.trees, tree)
}

// Draw draws all trees in the forest.
func (f *Forest) Draw(canvas *Canvas) {
	for _, tree := range f.trees {
		tree.Draw(canvas)
	}
}

// Canvas represents the drawing area.
type Canvas struct {
	// Canvas would have properties and methods to draw objects on it.
}
