package factory

import "fmt"

// Product is the interface that all concrete products should implement
type Product interface {
	Use() string
}

// ConcreteProductA is a concrete implementation of the Product interface
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using Product A"
}

// ConcreteProductB is a concrete implementation of the Product interface
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using Product B"
}

// ProductType is an enumeration of the different types of products
type ProductType int

const (
	ProductA ProductType = iota
	ProductB
)

// CreateProduct is the factory method that creates products based on the given type
func CreateProduct(productType ProductType) (Product, error) {
	switch productType {
	case ProductA:
		return &ConcreteProductA{}, nil
	case ProductB:
		return &ConcreteProductB{}, nil
	default:
		return nil, fmt.Errorf("unknown product type")
	}
}
