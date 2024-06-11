package prototype

import "fmt"

// Shape is the interface that all prototypes should implement
type Shape interface {
	Draw()
	Clone() Shape
}

// Circle is a concrete implementation of the Shape interface
type Circle struct {
	Radius int
	Color  string
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a %s circle with radius %d\n", c.Color, c.Radius)
}

func (c *Circle) Clone() Shape {
	return &Circle{Radius: c.Radius, Color: c.Color + "_clone"}
}

// Rectangle is another concrete implementation of the Shape interface
type Rectangle struct {
	Width  int
	Height int
	Color  string
}

func (r *Rectangle) Draw() {
	fmt.Printf("Drawing a %s rectangle with width %d and height %d\n", r.Color, r.Width, r.Height)
}

func (r *Rectangle) Clone() Shape {
	return &Rectangle{Width: r.Width, Height: r.Height, Color: r.Color + "_clone"}
}
