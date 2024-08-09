package visitor

import "fmt"

// ComputerPartVisitor is the interface for any type of visitor that can visit computer parts.
type ComputerPartVisitor interface {
	VisitKeyboard(keyboard *Keyboard)
	VisitMonitor(monitor *Monitor)
	VisitMouse(mouse *Mouse)
	VisitComputer(computer *Computer)
}

// ComputerPart is the interface for elements that can be visited.
type ComputerPart interface {
	Accept(ComputerPartVisitor)
}

// Keyboard is a type of ComputerPart that can be visited.
type Keyboard struct{}

func (k *Keyboard) Accept(visitor ComputerPartVisitor) {
	visitor.VisitKeyboard(k)
}

// Monitor is another type of ComputerPart that can be visited.
type Monitor struct{}

func (m *Monitor) Accept(visitor ComputerPartVisitor) {
	visitor.VisitMonitor(m)
}

// Mouse is another type of ComputerPart that can be visited.
type Mouse struct{}

func (m *Mouse) Accept(visitor ComputerPartVisitor) {
	visitor.VisitMouse(m)
}

// Computer is a composite object that can be visited.
type Computer struct {
	parts []ComputerPart
}

func (c *Computer) Accept(visitor ComputerPartVisitor) {
	for _, part := range c.parts {
		part.Accept(visitor)
	}
	visitor.VisitComputer(c)
}

func NewComputer() *Computer {
	return &Computer{
		parts: []ComputerPart{
			&Keyboard{},
			&Monitor{},
			&Mouse{},
		},
	}
}

// ComputerPartDisplayVisitor is a visitor that displays parts of a computer.
type ComputerPartDisplayVisitor struct{}

func (cpdv *ComputerPartDisplayVisitor) VisitKeyboard(keyboard *Keyboard) {
	fmt.Println("Displaying Keyboard.")
}

func (cpdv *ComputerPartDisplayVisitor) VisitMonitor(monitor *Monitor) {
	fmt.Println("Displaying Monitor.")
}

func (cpdv *ComputerPartDisplayVisitor) VisitMouse(mouse *Mouse) {
	fmt.Println("Displaying Mouse.")
}

func (cpdv *ComputerPartDisplayVisitor) VisitComputer(computer *Computer) {
	fmt.Println("Displaying Computer.")
}
