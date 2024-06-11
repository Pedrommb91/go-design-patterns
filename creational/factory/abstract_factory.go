package factory

// AbstractProductA is the interface for products in the first family
type AbstractProductA interface {
	UseA() string
}

// AbstractProductB is the interface for products in the second family
type AbstractProductB interface {
	UseB() string
}

// ConcreteProductA1 is a concrete implementation of AbstractProductA
type ConcreteProductA1 struct{}

func (p *ConcreteProductA1) UseA() string {
	return "Using Product A1"
}

// ConcreteProductA2 is a concrete implementation of AbstractProductA
type ConcreteProductA2 struct{}

func (p *ConcreteProductA2) UseA() string {
	return "Using Product A2"
}

// ConcreteProductB1 is a concrete implementation of AbstractProductB
type ConcreteProductB1 struct{}

func (p *ConcreteProductB1) UseB() string {
	return "Using Product B1"
}

// ConcreteProductB2 is a concrete implementation of AbstractProductB
type ConcreteProductB2 struct{}

func (p *ConcreteProductB2) UseB() string {
	return "Using Product B2"
}

// AbstractFactory is the interface for the abstract factory
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

func GetAbstractFactory(name string) AbstractFactory {
	if name == "1" {
		return &ConcreteFactory1{}
	} else if name == "2" {
		return &ConcreteFactory2{}
	}
	return nil
}

// ConcreteFactory1 is a concrete implementation of AbstractFactory
type ConcreteFactory1 struct{}

func (f *ConcreteFactory1) CreateProductA() AbstractProductA {
	return &ConcreteProductA1{}
}

func (f *ConcreteFactory1) CreateProductB() AbstractProductB {
	return &ConcreteProductB1{}
}

// ConcreteFactory2 is a concrete implementation of AbstractFactory
type ConcreteFactory2 struct{}

func (f *ConcreteFactory2) CreateProductA() AbstractProductA {
	return &ConcreteProductA2{}
}

func (f *ConcreteFactory2) CreateProductB() AbstractProductB {
	return &ConcreteProductB2{}
}
